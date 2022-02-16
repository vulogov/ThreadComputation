all: compile
pre:
	go get github.com/gammazero/deque
	go get github.com/antlr/antlr4/runtime/Go/antlr
	go get github.com/sirupsen/logrus
	go get github.com/cstockton/go-conv
	go get github.com/lrita/cmap
	go get github.com/google/uuid
	go get -u -t gonum.org/v1/gonum/...
	go get -u github.com/c2fo/vfs/v6/...
	go get -u github.com/levigross/grequests
build:
	rm -rf ./parser
	antlr -Dlanguage=Go -o parser ThreadComputation.g4
c:
	go build  -v ./...
test:
	go test -v
rebuild: pre build dynmod c test
compile: c test
