export GOPATH=$(PWD)

default:
	go get github.com/luebken/container-api
	go run src/github.com/luebken/main.go

run:
	go run src/github.com/luebken/container-api/main.go luebken/currentweather-nodejs