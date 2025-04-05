# Use the official Go 1.18 image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /backend

# Copy the Go module files to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire backend directory
COPY . .

# If Go files are in a subdirectory like 'src', specify it
WORKDIR /backend/cmd/app 
# Build the Go application
RUN go build -o ../../socialnetwork  

WORKDIR /backend
# Expose port 8080 for the backend server
EXPOSE 8080

# Run the application
CMD ["./socialnetwork"]