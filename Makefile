.DEFAULT_GOAL := build

build:
	go build -o ./bin/write-to-file-cli main.go
