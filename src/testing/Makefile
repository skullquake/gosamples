SRC=./src/a_test.go#_test.go extension of relevance to go driver
GO=go
GOPATH_=$(shell go env GOPATH)
GOFLAGS=
BIN=./bin/a.out
MKDIR=mkdir
$(BIN):export GOPATH=$(GOPATH_)
$(BIN):$(SRC)
	$(GO) test -v $(SRC)
.phony:clean
