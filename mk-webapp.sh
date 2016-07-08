#!/bin/bash
D=$(pwd)
cd webapp
shorthand index.shorthand > index.html
gopherjs build
cd "$D"
