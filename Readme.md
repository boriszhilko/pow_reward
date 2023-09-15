# Proof of Work Server

A simple TCP server that issues a Proof of Work (PoW) challenge to clients. Once the client successfully solves the challenge, the server rewards them with a word of wisdom.

## Features
- TCP server protected with a PoW challenge-response mechanism.
- Returns a quote from a collection of wisdoms once the challenge is met.
- Client utility to solve PoW and receive a quote.

## Prerequisites
- Docker
- Go (version 1.20)
- `make` tool installed (commonly available on UNIX systems)

## Quick Start
For a quick setup and run of the Wisdom Server and Client, you can use Docker Compose. Make sure you have Docker, Docker Compose, and `make` installed. Navigate to the project root directory and execute: `make all`. This command will set up and run both the server and client using Docker Compose.

## Detailed Build and Run
If you wish to build and run the Wisdom Server and Client individually, follow these steps:

### Wisdom Server
1. **Building the Server Docker Image:** Navigate to the project root directory and execute: `make build-server`.
2. **Running the Server Docker Container:** Once the image is built, start the server with: `make run-server`.

### Wisdom Client
1. **Building the Client Docker Image:** Navigate to the project root directory and execute: `make build-client`.
2. **Running the Client Docker Container:** After building, run the client with: `make run-client`.

## Clean-up
To remove the Docker containers and images for both the server and client, execute: `make clean`.

## Help
To view a list of all available commands provided by the Makefile, execute: `make help`.
