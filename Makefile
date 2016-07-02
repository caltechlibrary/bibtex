#
# Simple Makefile
#
build:
	go build -o bin/bibfilter cmds/bibfilter/bibfilter.go

install:
	env GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go

test:
	go test

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

release:
	./mk-release.sh

