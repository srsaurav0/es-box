# Use the official Golang base image with Go 1.18+ for compatibility with 'slices' and 'cmp'
FROM golang:1.23.2-alpine

# Install dependencies, including Git and bee tool
RUN apk add --no-cache git bash && \
    go install github.com/beego/bee@latest

# Set the Go module name (no go.mod file)
ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV MODULE_NAME=es-box

# Allow Git to recognize the /app directory as a safe directory
RUN git config --global --add safe.directory /app

# Set the working directory
WORKDIR /app

# Copy the Go code into the container
COPY . /app

# Copy the app.conf file into the container
COPY conf/app.conf /app/conf/app.conf

# Initialize Go module and tidy dependencies
RUN go mod init es-box || echo "go.mod already initialized" && go mod tidy

# Install dependencies (this should now work with the go.mod file)
RUN go mod tidy && go mod download

# Expose the Go app port (adjust this based on your app's settings)
EXPOSE 8080

# Set the entrypoint to run the 'bee run' command when the container starts
CMD ["bee", "run"]
