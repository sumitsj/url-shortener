# url-shortener

> Product Engineer Assignment || InfraCloud Technologies

> This is URL shortener service that accepts a URL as an argument over a REST API and
return a shortened URL as a result.

Prerequisites
-------------

- Git
- Docker
- Docker Compose
- [Go 1.22.1](https://go.dev/doc/install) and above
- Mongo DB (No need to install separately if you run this application using docker compose)
- [Make](https://formulae.brew.sh/formula/make)
- [Mockery](https://vektra.github.io/mockery/latest/installation/#installation)

> macOS 14.3 is used while developing this project. To avoid any compatibility issues, run it using docker.

Getting Started
---------------

#### Clone the repository

```bash
git clone https://github.com/sumitsj/url-shortener.git

cd url-shortener
```

#### Using Docker [Recommended]

```bash
# Build docker image
make docker-build

# Run application and dependencies using docker
make docker-run
```

#### Using Local Environment

```bash
# Create Env file
cp ./env.local .env

# Download modules
go mod download

# Build the project
make build

# Run the project
make run
```

The application starts at port 8080:

- `GET /ping` Health check endpoint, returns 'pong' message

---

- `POST /short` Creates shortened url for given url. Try to open shortened url in browser.

---

#### Running Unit Tests

```bash
# Run unit tests
make run-tests

# Run unit tests with coverage report in html format
make run-tests-with-coverage
```
---
#### Running Integration Tests

```bash
make run-integration-tests
```

#### Troubleshooting
If you get an error `Could not connect to Docker: Get "http://unix.sock/_ping": dial unix /var/run/docker.sock: connect: no such file or directory`, try setting `DOCKER_HOST` environment variable.

In case of Colima, you can do this by running below command.
```bash
export DOCKER_HOST="unix://$HOME/.colima/docker.sock"
```

---