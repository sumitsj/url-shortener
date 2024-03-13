FROM golang:1.22

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]