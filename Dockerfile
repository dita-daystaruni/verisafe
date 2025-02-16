# Base Image
FROM golang:latest

# Copy all the app contents
COPY . /app/

# Default working directory
WORKDIR /app

# Create a directory for the uploads within /app
RUN mkdir -p /app/uploads

# Install the go modules
RUN go mod tidy

# Install Goose for database migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Run app
CMD ["go", "run", "main.go"]

EXPOSE 3000
