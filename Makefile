# Makefile for Bug Archives

TAG = "\\033[32\;1mMakefile\\033[0m"

build:
	go build -o buglog buglog.go

new_bug:
	./buglog
