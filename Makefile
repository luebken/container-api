export GOPATH=$(PWD)

default:
	go get github.com/luebken/container-api

run:
	go run src/github.com/luebken/container-api/main.go luebken/currentweather-nodejs:latest