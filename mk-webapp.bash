#!/bin/bash
CWD=$(pwd)
cd webapp
mkpage "nav=nav.md" index.tmpl >index.html
mkpage "nav=nav.md" bibfilter.tmpl >bibfilter.html
mkpage "nav=nav.md" bibmerge.tmpl >bibmerge.html
mkpage "nav=nav.md" bibscrape.tmpl >bibscrape.html
gopherjs build
cd "$CWD"
