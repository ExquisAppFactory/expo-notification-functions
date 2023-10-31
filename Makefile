.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o main

zip:
	zip main.zip main

deploy: clean build

