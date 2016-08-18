#!/bin/bash
CWD=$(pwd)
cd webapp
echo "Building gopherjs parts"
gopherjs build
echo "Generating index.html"
mkpage "nav=nav.md" index.tmpl > index.html
echo "Generating bibfilter.html"
mkpage "nav=nav.md" bibfilter.tmpl > bibfilter.html
echo "Generating bibmerge.html"
mkpage "nav=nav.md" bibmerge.tmpl > bibmerge.html
echo "Generating bibscrape.html"
mkpage "nav=nav.md" bibscrape.tmpl > bibscrape.html
cd "$CWD"
