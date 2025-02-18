# Base Image
FROM golang:latest AS main

# Copy all the app contents
COPY . /app/

# Default working directory
WORKDIR /app

# Create a directory for the uploads within /app
RUN mkdir -p /app/uploads

# Install the go modules
RUN go mod tidy

# Run app
CMD ["go", "run", "main.go"]

EXPOSE 3000
