all: build

run: build
	./rb ./examples/hello-world-simple.rnb

build:
	go build src/rb.go

clean:
	rm ./rb
