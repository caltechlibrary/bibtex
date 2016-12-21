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
	./mk-webapp.bash

install:
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibmerge/bibmerge.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/bibscrape/bibscrape.go

test:
	go test

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
	./mk-webapp.bash
	./mk-website.bash
	./publish.bash

release:
	./mk-webapp.bash
	./mk-website.bash
	./mk-release.bash

