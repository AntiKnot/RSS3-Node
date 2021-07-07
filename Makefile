PROG = rss3node
OUTPUT = ./build

all:
	go get
	mkdir -p $(OUTPUT)
	go build -o $(OUTPUT)/$(PROG)
