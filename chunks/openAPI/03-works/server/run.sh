#!/bin/bash

# Install Go dependencies
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger

# get swagger
go get -u github.com/swaggo/swag/cmd/swag

# Generate Swagger documentation
swag init

# Create a new Go module
go mod init github.com/your-username/your-project

# Run the server
echo "Starting the server..."
go run main.go

# Run the client
echo "Running the client..."
go run client.go
