# Stage 1: Build the Go application
FROM golang:1.22.4 AS build

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Set environment variables for the build
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct

# Build the Go application
RUN go build -v -o app .

# Stage 2: Create a minimal image with the built binary
FROM scratch

# Copy the built Go binary from the build stage
COPY --from=build /go/src/app/app /go/bin/app

ENV ENVIRONMENT=DOCKER

# Set the entrypoint to the built binary
ENTRYPOINT ["/go/bin/app"]

