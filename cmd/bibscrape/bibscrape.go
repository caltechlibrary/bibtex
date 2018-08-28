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
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	// Caltech Library Packages
	"github.com/caltechlibrary/bibtex"
	"github.com/caltechlibrary/bibtex/scrape"
	"github.com/caltechlibrary/cli"
)

var (
	// Standard Options
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	quiet            bool
	newLine          bool
	generateMarkdown bool
	generateManPage  bool

	entrySeparator = "(\n|\r\n)"
	useType        = `pseudo`
	addKeys        bool

	synopsis = `
parse plain text making a best guess to generate pseudo bib entries
`

	description = `
_bibscrape_ parses a plain text file for BibTeX entry making a best guess to generate pseudo bib entries 
that can import into JabRef for cleaning
`

	examples = `
` + "```" + `
    bibscrape -entry-separator "[0-9][0-9]0-9][0-9]\.\n" mytest.bib
` + "```" + `
`
)

func main() {
	pseudo_id := 0

	app := cli.NewCli(bibtex.Version)
	app.AddParams("[OPTIONS]", "FILENAME")

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
	app.BoolVar(&addKeys, "k", false, "add a missing key")
	app.StringVar(&entrySeparator, "e,entry-separator", entrySeparator, `Set the default entry separator (defaults to \n\n)`)
	app.StringVar(&useType, "t", useType, `Set the entry type  (defaults to pseudo)`)

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
			fmt.Fprintf(app.Eout, "%s", err)
			os.Exit(1)
		}
		for {
			entry, buf = scrape.NextEntry(buf, re)
			if len(entry) > 0 {
				elem = scrape.Scrape(entry)
				if useType != "" {
					elem.Type = useType
				}
				if addKeys == true && len(elem.Keys) == 0 {
					elem.Keys = append(elem.Keys, fmt.Sprintf("pseudo_id_%d", pseudo_id))
					pseudo_id++
				}
				fmt.Fprintf(app.Out, "%s\n\n", elem)
			}
			entry = nil
			elem = nil
			if buf == nil || len(buf) == 0 {
				break
			}
		}
	}

	for _, fname := range args {
		scrapeFile(fname, reEntrySeparator)
	}
	if newLine {
		fmt.Fprintf(app.Out, "\n")
	}
}
