# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o greetings-app

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./greetings-app"]
