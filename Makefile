default: test

.PHONY: test
test: deps
	go test -v -covermode=count -coverprofile="./count.out" ./...

.PHONY: deps
deps:
	go get -u -v ./...

docs:
	cat doc/head.md > README.md
	godoc2md github.com/go-lo/go-lo >> README.md
