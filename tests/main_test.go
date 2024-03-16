//go:build integration

package tests

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	client  *mongo.Client
	pool    *dockertest.Pool
	appPort string
)

func TestMain(m *testing.M) {

	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Println("Trying to connect to Docker")
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	log.Println("Trying to create docker network")
	network, err := pool.CreateNetwork("backend")
	if err != nil {
		log.Fatalf("Could not create network to docker: %s \n", err)
	}

	var mongoCnt, apiCnt *dockertest.Resource
	mongoCnt, err = startMongoDB(pool, "latest", network)
	if err != nil {
		cleanUp(1, network, mongoCnt, apiCnt)
	}

	apiCnt, err = startAPI(pool, network)
	if err != nil {
		cleanUp(1, network, mongoCnt, apiCnt)
	}

	println("Starting tests")
	code := m.Run()
	println("Stopping tests")

	defer cleanUp(code, network, mongoCnt, apiCnt)
}

func startMongoDB(pool *dockertest.Pool, mongoVersion string, network *dockertest.Network) (*dockertest.Resource, error) {
	r, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       "mongo",
		Repository: "mongo",
		Tag:        mongoVersion,
		//Mounts:     []string{getProjectRootPath() + "/mongo_scripts/init_db.js:/docker-entrypoint-initdb.d/init_db.js"},
		Networks: []*dockertest.Network{network},
	})
	if err != nil {
		fmt.Printf("Could not start Mongodb: %v \n", err)
		return r, err
	}
	mongoPort := r.GetPort("27017/tcp")

	fmt.Printf("mongo-%s - connecting to : %s \n", mongoVersion, fmt.Sprintf("mongodb://localhost:%s", mongoPort))
	if err := pool.Retry(func() error {
		var err error

		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s", mongoPort))
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}

		err = client.Ping(context.TODO(), nil)
		if err == nil {
			fmt.Println("successfully connected to Mongodb.")
		}
		return err

	}); err != nil {
		fmt.Printf("Could not connect to mongodb container: %v \n", err)
		return r, err
	}

	return r, nil
}

func startAPI(pool *dockertest.Pool, network *dockertest.Network) (*dockertest.Resource, error) {
	mongoDbPort := "27017"
	apiContainerName := "url-shortener-integration-test"

	envs := []string{
		fmt.Sprintf("MONGO_DB_URI=mongodb://mongodb:%s/url-shortener", mongoDbPort),
		"SERVER_ADDR=0.0.0.0",
		"SERVER_PORT=8080",
	}

	r, err := pool.BuildAndRunWithBuildOptions(
		&dockertest.BuildOptions{
			ContextDir: "../",
			Dockerfile: "Dockerfile",
		},
		&dockertest.RunOptions{
			Name:     apiContainerName,
			Env:      envs,
			Networks: []*dockertest.Network{network},
		})
	if err != nil {
		fmt.Printf("Could not start %s: %v \n", apiContainerName, err)
		return r, err
	}

	waiter, err := pool.Client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container:    apiContainerName,
		OutputStream: log.Writer(),
		ErrorStream:  log.Writer(),
		RawTerminal:  true,
		Logs:         true,
		Stream:       true,
		Stdout:       true,
		Stderr:       true,
	})
	if err != nil {
		fmt.Println("unable to get LOGS: ", err)
	}
	defer waiter.Close()

	appPort = r.GetPort("8080/tcp")
	if err := pool.Retry(func() error {

		resp, err := http.Get("http://localhost:" + appPort + "/ping")
		if err != nil {
			fmt.Printf("trying to connect to %s on localhost:%s, got : %v \n", apiContainerName, appPort, err)
			return err
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("trying to connect to %s on localhost:%s, got : %v , status: %v \n", apiContainerName, appPort, err, resp.StatusCode)
			return err
		}

		fmt.Println("status: ", resp.StatusCode)
		rs, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("RESPONSE: %s \n", rs)
		return nil
	}); err != nil {
		fmt.Printf("Could not connect to %s container: %v \n", apiContainerName, err)
		return r, err
	}

	return r, nil
}

func cleanUpResources(resources []*dockertest.Resource) {
	fmt.Println("removing resources.")
	for _, resource := range resources {
		if resource != nil {
			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s\n", err)
			}
		}
	}
}

func cleanUp(code int, network *dockertest.Network, resources ...*dockertest.Resource) {
	cleanUpResources(resources)
	network.Close()
	os.Exit(code)
}
