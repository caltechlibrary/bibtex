#!/bin/bash
D=$(pwd)
cd webapp
gopherjs build
cd "$D"
