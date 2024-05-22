# Go chat

## Description

Simple Golang text chat web server

## Features

* Room conversation
* User-to-user conversation
* User groups conversation
* Private chats with messages disappearing after a certain time period
* Avatar uploading

## Technology Stack

* Golang
* Gin
* MongoDB
* Websocket messaging
* Docker

## Local launch

### Prerequisites

* [docker compose](https://docs.docker.com/compose/)
* [go](https://go.dev/)

1. Clone the repository

```bash
git clone git@github.com:Hellwest/go-chat.git
cd go-chat
```

2. Start the database and its web interface

* `docker compose up`

The web interface is now available at [localhost:8081](http://localhost:8081)

3. Run the go-chat app

* `go run main.go`
