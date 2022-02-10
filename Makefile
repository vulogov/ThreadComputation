all: check
pre:
	go get github.com/gammazero/deque
	go get github.com/antlr/antlr4/runtime/Go/antlr
	go get github.com/sirupsen/logrus
	go get github.com/cstockton/go-conv
build:
	rm -rf ./parser
	antlr -Dlanguage=Go -o parser ThreadComputation.g4
	go build -v ./...
test:
	go test -v
check: pre build test
