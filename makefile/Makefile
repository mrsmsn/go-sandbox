VERSION  := $(shell git describe --tags 2>/dev/null)
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null)

build:
	go build -ldflags '-X main.Version=${VERSION} -X main.Revision=${REVISION}' main.go
