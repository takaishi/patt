BUILD=tmp/bin

default: build

build:
	go build -o $(BUILD)/patt
