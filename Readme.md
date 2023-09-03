# PoW Wisdom Server

A simple TCP server that issues a Proof of Work (PoW) challenge to clients. Once the client successfully solves the
challenge, the server rewards them with a word of wisdom.

## Features

- TCP server protected with a PoW challenge-response mechanism.
- Returns a quote from a collection of wisdoms once the challenge is met.
- Client utility to solve PoW and receive a quote.

## Prerequisites

- Docker
- Go (version 1.20)

## Quick Start

For a quick setup and run of the Wisdom Server and Client, you can use Docker Compose.

1. **Run with Docker Compose:**

   Make sure you have Docker and Docker Compose installed. Navigate to the project root directory and execute:

    ```bash
    docker-compose up
    ```

   This command will set up and run both the server and client using the configurations provided in the `docker-compose.yml` file.

---

## Detailed Build and Run

If you wish to build and run the Wisdom Server and Client individually, follow these detailed steps:

### Wisdom Server

1. **Building the Server Docker Image:**
  - Navigate to the project root directory and execute:
    ```bash
    docker build -t wisdom-server -f cmd/server/Dockerfile .
    ```
  - This will build the Wisdom Server Docker image.

2. **Running the Server Docker Container:**
  - Once the image is built, run the server with:
    ```bash
    docker run --name wisdom-server -p 8080:8080 wisdom-server
    ```
  - The server will start and listen on ports `8080` for TCP connections and `8081` for HTTP connections.

### Wisdom Client

1. **Building the Client Docker Image:**
  - Navigate to the project root directory and execute:
    ```bash
    docker build -t wisdom-client -f cmd/client/Dockerfile .
    ```
  - This will build the Wisdom Client Docker image.

2. **Running the Client Docker Container:**
  - After building, run the client with:
    ```bash
    docker run -e SERVER_URL=wisdom-server:8080 --link wisdom-server wisdom-client
    ```
  - The client will connect to the Wisdom Server, solve the PoW challenge, and receive a word of wisdom.
