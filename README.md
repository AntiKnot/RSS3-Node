# RSS3 Node

## Description

## Setup

1. Install ``protoc`` and ``protoc-gen-go``

```bash
# MacOS
brew install protobuf
# Linux
apt install -y protobuf-compiler
protoc --version # Ensure compiler version is 3+

go install google.golang.org/protobuf/cmd/protoc-gen-go
```

## Run

To test locally:

```bash
make 
# In terminal 1
./build/rss3node -p 2233
# In another terminal
./build/rss3node -p 2234
```

## Project structure

```bash

├── commands  # Handling user commands
├── communication # Discovering and connecting with other peers
├── db  # Distributed key-value db based on pubsub
├── validator  # Validator used for updating the value in db
├── rss3  # Interactions with RSS3File
├── storage # IPFS storage layer
├── types # Wrapper for some common types
├── node  # Running a node with multiple services
├── config
├── main.go
├── flags.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
