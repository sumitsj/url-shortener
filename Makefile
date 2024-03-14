build:
	go build -o main .

run:
	go run ./main.go

docker-build:
	docker-compose build

docker-run:
	docker-compose up

run-tests:
	go clean -testcache
	go test `go list ./...`
