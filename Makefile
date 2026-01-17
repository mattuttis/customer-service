.PHONY: build run test clean docker-build docker-run lint tidy

# Variables
APP_NAME := customer-service
MAIN_PATH := ./cmd/api
DOCKER_IMAGE := customer-service:latest

# Build the application
build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	go run $(MAIN_PATH)

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Tidy dependencies
tidy:
	go mod tidy

# Lint (requires glangci-lint)
lint:
	golangci-lint run

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Format code
fmt:
	g fmt ./...

