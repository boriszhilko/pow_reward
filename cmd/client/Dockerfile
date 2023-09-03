# Use the specific Go version.
FROM golang:1.20-alpine

WORKDIR /app

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Build the client
COPY . .
WORKDIR /app/cmd/client
RUN go build -o client

CMD ["./client"]
