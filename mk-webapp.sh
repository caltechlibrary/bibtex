#!/bin/bash
CWD=$(pwd)
cd webapp
shorthand index.shorthand > index.html
shorthand bibfilter.shorthand > bibfilter.html
shorthand bibmerge.shorthand > bibmerge.html
shorthand bibscrape.shorthand > bibscrape.html
gopherjs build
cd "$CWD"
