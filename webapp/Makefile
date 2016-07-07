#
# Build our webapp using GopherJS
#
build: index.html index.js webapp.go bibfilter/bibfilter.go
	gopherjs build

clean:
	if [ -f webapp.js ]; then /bin/rm webapp.js; fi
	if [ -f webapp.js.map ]; then /bin/rm webapp.js.map; fi

