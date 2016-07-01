//
// Package mkbib is a quick and dirty plain text parser for generating
// a Bibtex citation
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// This package
	"github.com/caltechlibrary/bibtex"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	flag.BoolVar(&showLicense, "l", false, "display license information")
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTIONS] [ORIG_FILENAME] [BIBTEX_FILENAME]

 OPTIONS

`, appname)

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s  (defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
		})

		fmt.Printf("\n\n Version: %s\n", bibtex.Version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf(" Version: %s\n", bibtex.Version)
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Printf(`
 %s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`, appname)
		os.Exit(0)
	}

	var err error
	in := os.Stdin
	out := os.Stdout
	if len(args) > 0 {
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open %s, %s\n", args[0], err)
			os.Exit(1)
		}
	}
	if len(args) > 1 {
		out, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open %s, %s\n", args[0], err)
			os.Exit(1)
		}
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	src, err := bibtex.Parse(src)
	fmt.Fprintf(out, "%s", src)
}
