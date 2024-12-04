# Use the official Golang image to create a build artifact.
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /workdir

# Install mage build tool
RUN go install github.com/magefile/mage@v1.15.0

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the go binary
RUN mage -v binary:build

# Start a new stage from scratch
FROM scratch

# Run as non-root user
USER 1000:1000

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /workdir/build/url-shortener /url-shortener

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/url-shortener"]
CMD ["-port", "8080"]