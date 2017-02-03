#
# Simple Makefile
#
PROJECT = bibtex

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bibtex.go cmds/bibfilter/bibfilter.go cmds/bibmerge/bibmerge.go cmds/bibscrape/bibscrape.go
	env CGO_ENABLED=0 go build -o bin/bibfilter cmds/bibfilter/bibfilter.go
	env CGO_ENABLED=0 go build -o bin/bibmerge cmds/bibmerge/bibmerge.go
	env CGO_ENABLED=0 go build -o bin/bibscrape cmds/bibscrape/bibscrape.go

install:
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibmerge/bibmerge.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibscrape/bibscrape.go

test:
	go test

status:
	git status

save:
	git commit -am "Quick save"
	git push origin $(BRANCH)

clean:
	if [ -f index.html ]; then /bin/rm *.html; fi
	if [ -f webapp/index.html ]; then /bin/rm webapp/*.html; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f webapp/webapp.js ]; then rm -f webapp/webapp.js; fi
	if [ -f webapp/webapp.js.map ]; then rm -f webapp/webapp.js.map; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

website:
	./mk-website.bash
	./mk-webapp.bash

publish:
	./mk-website.bash
	./mk-webapp.bash
	./publish.bash

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v bibfilter.md dist/
	cp -v bibmerge.md dist/
	cp -v bibscrape.md dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

dist/linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibfilter cmds/bibfilter/bibfilter.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibmerge cmds/bibmerge/bibmerge.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibscrape cmds/bibscrape/bibscrape.go

dist/windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibfilter.exe cmds/bibfilter/bibfilter.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibmerge.exe cmds/bibmerge/bibmerge.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibscrape.exe cmds/bibscrape/bibscrape.go

dist/macosx-amd64:
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibfilter cmds/bibfilter/bibfilter.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibmerge cmds/bibmerge/bibmerge.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibscrape cmds/bibscrape/bibscrape.go


dist/raspbian-arm7:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibfilter cmds/bibfilter/bibfilter.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibmerge cmds/bibmerge/bibmerge.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibscrape cmds/bibscrape/bibscrape.go

