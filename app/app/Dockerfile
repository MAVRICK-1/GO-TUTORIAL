# Stage 1: Build the Go application
FROM golang:1.19.3-buster AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download && go mod verify

# Copy the source code into the container
COPY main.go ./

# Build the Go application
RUN go build -o /app

# Stage 2: Create a minimal image using distroless
FROM gcr.io/distroless/base-debian11

# Copy the built application from the previous stage
COPY --from=build /go/bin/app /app

# Set the entry point to the binary
ENTRYPOINT ["/app"]
