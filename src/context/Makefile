SRC=./src/a.go
GO=go
GOPATH_=$(shell go env GOPATH)
GOFLAGS=-x
GOFLAGS+=-ldflags '-w -s' 
BIN=./bin/a.out
MKDIR=mkdir
RM=rm
UPX=upx
all:compress
$(BIN):export GOPATH=$(GOPATH_)
$(BIN):$(SRC)
	@-$(MKDIR) -p $(@D)
	$(GO) build $(GOFLAGS) -o $(BIN) $(SRC)
compress:$(BIN)
	@-$(UPX) $(BIN)
run:$(BIN)
	@-$(BIN)
test:$(BIN)
	@-echo testing request interrupt...
	@-curl http://localhost:8090/hello&
	@-sleep 1&&killall -s9 curl
.phony:clean
clean:
	$(RM) $(BIN)
