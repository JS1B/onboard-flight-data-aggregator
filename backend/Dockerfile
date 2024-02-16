# Start with a base image that includes the Go tools and libraries
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests and download modules
# This layer is cached and reused if these files do not change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's source code, excluding files in .dockerignore
COPY . .

# Build the application with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux go build -o ofda-server .

# Start a new stage from a smaller base image for a smaller final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /ofda

# Copy the binary from the previous stage
COPY --from=builder /app/ofda-server ./ofda-server

# Set executable permissions for the binary
RUN chmod +x ofda-server

# Copy the .env file to the /ofda directory
COPY .env ./

# Expose the port the app runs on
EXPOSE 8080

# Install sqlite
RUN apk add --no-cache sqlite

# Copy the database initialization script
COPY db/init.sql ./db/init.sql

# Create a directory for the database
RUN mkdir -p /var/lib/ofda

# Initialize the database
RUN sqlite3 /var/lib/ofda/ofda.db < db/init.sql

# Make a shortcut to the database
RUN ln -s /var/lib/ofda/ofda.db ofda.db

# Command to run the binary
CMD ["./ofda-server"]
