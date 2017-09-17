# vi: ft=make
.PHONY: proto test get ci docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I $$GOPATH/src/github.com/baopham/gominderproto summarizer.proto --lile-server_out=. --go_out=plugins=grpc:$$GOPATH/src
	

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

docker:
	docker build . -t baopham/summarizer:latest
