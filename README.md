# Word of Wisdom TCP Server

This is a TCP server that provides a quote from a collection of quotes from the "Word of Wisdom" book after Proof of Work (PoW) verification.

## Proof of Work

The server is protected from DDOS attacks using Proof of Work. This is achieved using a challenge-response protocol where the client has to solve a PoW challenge before the server sends a quote. The PoW algorithm used in this implementation is finding a hash that is less than a target value. This algorithm is chosen because it is widely using in blockchain-based systems and is considered to be a computationally expensive algorithm.

## Server

The server runs as a Docker container. It listens on a specified TCP port and expects a PoW challenge from a client. After the client has successfully solved the challenge, the server sends a quote from a collection of quotes.

## Client

The client runs as a Docker container. It connects to the server on a specified TCP port and solves the PoW challenge sent by the server. After the challenge is solved, the client displays the quote received from the server.

## Usage

To build the server and client Docker images, run:

```
make build
```

To start the server and client containers, run:

```
make run
```

To stop the server and client containers, run:

```
make stop
```

To stop and remove the containers and images, run:

```
make clean
```

The server listens on port `8090` by default, and client connects to `server:8090`, but this can be changed in the `.env` file.

## Requirements

To run this server and client, you will need to have Docker installed on your system. The implementation uses Go version 1.20.3, but this is not required to be installed on your system as Docker images are used to build and run the server and client.