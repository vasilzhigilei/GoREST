# syntax=docker/dockerfile:1
FROM golang:1.17

# Download Go modules
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN go build -o /gorest
 
EXPOSE 8080
 
CMD ["/gorest"]