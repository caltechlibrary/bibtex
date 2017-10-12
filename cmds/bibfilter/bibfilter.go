//
// bibfilter reads a bibfile and writes it out. It can exclude entry types
// or list only included entry types.
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
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex"
	"github.com/caltechlibrary/cli"
)

var (
	// Standard Options
	showHelp     bool
	showVersion  bool
	showLicense  bool
	showExamples bool

	// Application Options
	include = bibtex.DefaultInclude
	exclude = ""

	usage = `USAGE: %s [OPTION] BIBFILE`

	description = `

SYSNOPSIS

%s filters BibTeX files by entry type.

`

	examples = `

EXAMPLES
	
	%s -include author,title my-works.bib

Renders a BibTeX file containing only author and title from my-works.bib

`
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")

	// Application Options
	flag.StringVar(&include, "include", include, "a comma separated list of tags to include")
	flag.StringVar(&exclude, "exclude", exclude, "a comma separated list of tags to exclude")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, strings.ToUpper(appName), bibtex.Version)
	cfg.LicenseText = fmt.Sprintf(bibtex.LicenseText, appName, bibtex.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionText = "OPTIONS\n\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}
	if showExamples == true {
		if len(args) > 0 {
			fmt.Println(cfg.Example(args...))
		} else {
			fmt.Println(cfg.ExampleText)
		}
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	var (
		err      error
		elements []*bibtex.Element
		buf      []byte
	)

	in := os.Stdin
	out := os.Stdout

	if len(args) > 0 {
		fname := args[0]
		args = args[1:]
		buf, err = ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, %s\n", fname, err)
			os.Exit(1)
		}
	} else {
		data := make([]byte, 100)
		for {
			_, err = in.Read(data)
			if err != nil {
				break
			}
			buf = append(buf[:], data[:]...)
		}
	}

	if len(args) > 0 {
		fname := args[0]
		args = args[1:]
		out, err := os.Create(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s, %s\n", fname, err)
		}
		defer out.Close()
	}

	elements, err = bibtex.Parse(buf)

	for _, element := range elements {
		if strings.Contains(include, element.Type) {
			if len(exclude) == 0 || strings.Contains(exclude, element.Type) == false {
				fmt.Fprintf(out, "%s\n", element)
			}
		}
	}
}
