# Define Go flags
BINARY_NAME=gomovies
GOFLAGS = -ldflags="-s -w"

.PHONY: clean all

# use @ before recipe (command) to not print it
# we have 4 targets : test, build, clean, all

test:
	go test ./...

build:
	@echo "running binary build .."
	go build $(GOFLAGS) -o bin/${BINARY_NAME} main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

all: build