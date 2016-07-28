#
# Simple Makefile
#
build:
	go build -o bin/bibfilter cmds/bibfilter/bibfilter.go
	go build -o bin/bibmerge cmds/bibmerge/bibmerge.go
	go build -o bin/bibscrape cmds/bibscrape/bibscrape.go
	./mk-webapp.bash
	./mk-website.bash

install:
	env GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go
	env GOBIN=$(HOME)/bin go install cmds/bibmerge/bibmerge.go
	env GOBIN=$(HOME)/bin go install cmds/bibscrape/bibscrape.go

test:
	go test

save:
	git commit -am "Quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f webapp/webapp.js ]; then rm -f webapp/webapp.js; fi
	if [ -f webapp/webapp.js.map ]; then rm -f webapp/webapp.js.map; fi
	if [ -f bibtex-binary-release.zip ]; then rm -f bibtex-binary-release.zip; fi

publish:
	./publish.bash

release:
	./mk-webapp.bash
	./mk-website.bash
	./mk-release.bash

