test:
	go test -v -cover ./test/...

run:
	go run .

build:
	go build -o ethiocal .

.PHONY: test run build