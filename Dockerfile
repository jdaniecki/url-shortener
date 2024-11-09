# Use the official Golang image to create a build artifact.
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /workdir

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install mage
RUN go install github.com/magefile/mage@v1.15.0

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with CGO disabled
RUN mage -v binary:build

# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /workdir/build/url-shortener /url-shortener

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/url-shortener"]