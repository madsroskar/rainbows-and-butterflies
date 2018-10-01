all: build

run: 
	go run ./... ./examples/hello-world-simple.rnb

build:
	go build src/rb.go

clean:
	rm ./rb
