.PHONY: dev build publish

all: build

build: 
	go build -o Mystat main.go

publish: 
	go build -o Mystat main.go
	strip Mystat
	upx Mystat
