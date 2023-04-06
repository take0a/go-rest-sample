.PHONY: clean build

clean:
	rm -rf dist

build: clean
	go build -o ./dist/server ./server

start: build
	source ./.env.dev && ./dist/server
