#!/bin/bash
CWD=$(pwd)
cd webapp
shorthand index.shorthand > index.html
shorthand bibfilter.shorthand > bibfilter.html
shorthand bibmerge.shorthand > bibmerge.html
gopherjs build
cd "$CWD"
