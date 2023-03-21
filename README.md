## CARD-API

This is a simple REST API for managing decks of cards. The API is built using the Go programming language, the Gorilla Mux package for routing, and listens on port 8000.

## Installation
Install Go, if not already installed. Instructions can be found at https://golang.org/doc/install.
Install the Gorilla Mux package:
```bash
$ go get -u github.com/gorilla/mux
```

## Running the app
To run the server, execute the following command in the project directory:
```bash
$ make run
```
Alternatively , if you do not have make installed , simply run the following command: 
```bash
$ go run main.go handler.go deck.go card.go
```
The server will start on port 8000.

To build the binary without running it, execute the following command in the project directory:
```bash
$ make build
```

To run unit tests:
```bash
$ make test
```

To clean the build directory, execute the following command in the project directory:
```bash
$ make clean
```

## API Endpoints
- `POST /deck/new`: Create a new deck of cards.
- `GET /deck/{id}`: Get a deck of cards by its unique identifier.
- `POST /deck/{id}/draw`: Draw one or more cards from the specified deck

## Stay in touch

- Author - [Bryan Toromo]

