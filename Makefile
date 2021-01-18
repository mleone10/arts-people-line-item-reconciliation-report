build: clean test
	go build -o aprecon

clean:
	rm -f aprecon

test:
	go test ./...
