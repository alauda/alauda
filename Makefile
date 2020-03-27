GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
SOURCE=cmd/alauda/main.go
TARGET=/usr/local/bin/alauda

all: build test

build: 
	$(GOBUILD) -o $(TARGET) -v $(SOURCE)

test: 
	$(GOTEST) -v ./...

clean:
	rm $(TARGET)