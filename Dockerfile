# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22.3 AS build

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download -x

RUN go mod tidy

# Copy the .env file
COPY .env ./

COPY services.json ./

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o /bin/server ./cmd

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /bin/server /bin/

# Copy the .env file from the build stage
COPY --from=build /app/.env ./

# Copy the services.json file from the build stage
COPY --from=build /app/services.json ./

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Expose the port that the application listens on
EXPOSE 82

# Run the server when the container starts
CMD ["/bin/server"]