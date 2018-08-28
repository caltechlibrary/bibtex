#
# Simple Makefile
#
PROJECT = bibtex

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bibtex.go cmd/bibfilter/bibfilter.go cmd/bibmerge/bibmerge.go cmd/bibscrape/bibscrape.go
	go build -o bin/bibfilter cmd/bibfilter/bibfilter.go
	go build -o bin/bibmerge cmd/bibmerge/bibmerge.go
	go build -o bin/bibscrape cmd/bibscrape/bibscrape.go

install:
	env GOBIN=$(GOPATH)/bin go install cmd/bibfilter/bibfilter.go
	env GOBIN=$(GOPATH)/bin go install cmd/bibmerge/bibmerge.go
	env GOBIN=$(GOPATH)/bin go install cmd/bibscrape/bibscrape.go

test:
	go test

status:
	git status

save:
	if [ "$(pwd)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick save"; fi
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
	./package-versions.bash > dist/package-versions.txt
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

dist/linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibfilter cmd/bibfilter/bibfilter.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibmerge cmd/bibmerge/bibmerge.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/bibscrape cmd/bibscrape/bibscrape.go

dist/windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibfilter.exe cmd/bibfilter/bibfilter.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibmerge.exe cmd/bibmerge/bibmerge.go
	env GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/bibscrape.exe cmd/bibscrape/bibscrape.go

dist/macosx-amd64:
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibfilter cmd/bibfilter/bibfilter.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibmerge cmd/bibmerge/bibmerge.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/bibscrape cmd/bibscrape/bibscrape.go


dist/raspbian-arm7:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibfilter cmd/bibfilter/bibfilter.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibmerge cmd/bibmerge/bibmerge.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/bibscrape cmd/bibscrape/bibscrape.go

