BINARY_NAME=gomovies

# use @ before recipe (command) to not print it
# we have 2 targets : build, clean
build:
	@echo "running binary build .."
	go build -ldflags="-s -w" -o bin/${BINARY_NAME} main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

all: build