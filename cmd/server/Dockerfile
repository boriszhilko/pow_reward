# Use the specific Go version.
FROM golang:1.20-alpine

WORKDIR /app

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Build the server
COPY . .
WORKDIR /app/cmd/server
RUN go build -o server

EXPOSE 8080

CMD ["./server"]
