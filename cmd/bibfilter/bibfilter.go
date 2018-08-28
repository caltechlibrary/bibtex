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
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex"
	"github.com/caltechlibrary/cli"
)

var (
	// Standard Options
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	newLine          bool
	generateMarkdown bool
	generateManPage  bool
	inputFName       string
	outputFName      string
	quiet            bool

	// Application Options
	include = bibtex.DefaultInclude
	exclude = ""

	synopsis = `filter a bibTeX file for specific fields`

	description = `
_bibfilter_ filters BibTeX files by entry type.
`

	examples = `
` + "```" + `
	bibfilter -include article,book my-works.bib
` + "```" + `
Renders a BibTeX file containing only article and book from my-works.bib
`
)

func main() {
	// Configuration and command line interation
	app := cli.NewCli(bibtex.Version)
	app.AddParams("[OPTIONS]", "BIBFILE")

	// Add Help Docs
	app.SectionNo = 1 // Manual page section number to document
	app.AddHelp("synopsis", []byte(synopsis))
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display examples")
	app.StringVar(&inputFName, "i,input", "", "input file name")
	app.StringVar(&outputFName, "o,output", "", "output file name")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// Application Options
	app.StringVar(&include, "include", include, "a comma separated list of entry type(s) to include")
	app.StringVar(&exclude, "exclude", exclude, "a comma separated list of entry type(s) to exclude")

	// We're ready to process args
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Handle options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintf(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	var (
		elements []*bibtex.Element
		buf      []byte
	)

	in := app.In
	out := app.Out

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
		out, err = os.Create(fname)
		if err != nil {
			fmt.Fprintf(app.Eout, "%s, %s\n", fname, err)
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
	if newLine {
		fmt.Fprintf(out, "\n")
	}
}
