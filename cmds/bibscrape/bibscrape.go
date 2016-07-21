// bibscrape - scrape a plain text file and render a pseudo BibTeX record that will import into JabRef.
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
	"regexp"

	// Caltech Library Packages
	"github.com/caltechlibrary/bibtex"
	"github.com/caltechlibrary/bibtex/scrape"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	entrySeparator = "(\n|\r\n)"
)

func usage(fp *os.File, appname string) {
	fmt.Fprintf(fp, `
 USAGE %s [OPTIONS] FILENAME

 Parse the plain text file for BibTeX entry making a best guess
 to generate pseudo bib entries that can import into JabRef for
 cleanup.

 OPTIONS

`, appname)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fp, "   -%s %s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fp, ` 

 EXAMPLE
 	
	Use an 4 digit ID number and period to indicate the start of my bib
	records.

	    %s -entry-separator "[0-9][0-9]0-9][0-9]\.\n" mytest.bib

	
 Version %s

`, appname, bibtex.Version)
}

func license(fp *os.File, appname string) {
	fmt.Fprintf(fp, ` %s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`, appname)
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")

	flag.StringVar(&entrySeparator, "e", entrySeparator, `Set the default entry separator (\n\n)`)
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(os.Stdout, appname)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf(" Version: %s", bibtex.Version)
		os.Exit(0)
	}

	if showLicense == true {
		license(os.Stdout, appname)
		os.Exit(0)
	}

	reEntrySeparator := regexp.MustCompile(entrySeparator)

	scrapeFile := func(fname string, re *regexp.Regexp) {
		var (
			buf   []byte
			entry []byte
			elem  *bibtex.Element
			err   error
		)
		buf, err = ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		for {
			entry, buf = scrape.NextEntry(buf, re)
			if len(entry) > 0 {
				elem = scrape.Scrape(entry)
				fmt.Fprintf(os.Stdout, "%s\n\n", elem)
			}
			entry = nil
			elem = nil
			if buf == nil || len(buf) == 0 {
				break
			}
		}
	}

	args := flag.Args()
	for _, fname := range args {
		scrapeFile(fname, reEntrySeparator)
	}
}
