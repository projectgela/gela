.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/projectgela/gela/cmd/swarm
	go build -o bin/gela     github.com/projectgela/gela/cmd/gela
	go build -o bin/bootnode github.com/projectgela/gela/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/projectgela/gela/swarm/...
