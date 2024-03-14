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
cp ./env.docker .env

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

#### Running Tests

```bash
make run-tests
```