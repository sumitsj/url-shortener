build:
	go build -o main .

run:
	go run ./main.go

docker-build:
	docker-compose build

docker-run:
	docker-compose up
