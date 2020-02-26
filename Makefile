default: test

.PHONY: test
test:
	go test -v -covermode=count -coverprofile="./count.out" ./...

README.md:
	cat doc/head.md > README.md
	godoc2md github.com/go-lo/go-lo >> README.md

go-lo.pb.go:
	protoc -I protos/ protos/go-lo.proto --go_out=plugins=grpc:.
