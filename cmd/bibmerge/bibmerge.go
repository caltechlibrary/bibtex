//
// bibset reads in two bibfiles and writes out new one based on operation selected (e.g. join, diff, intersect, exclusive).
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

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex"
	"github.com/caltechlibrary/cli"
)

var (
	showHelp         bool
	showVersion      bool
	showLicense      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	newLine          bool
	quiet            bool
	generateMarkdown bool
	generateManPage  bool

	mergeJoin      bool
	mergeDiff      bool
	mergeIntersect bool
	mergeExclusive bool

	synopsis = `merge BibTeX files`

	description = `
_bibmerge_ will merge combine two BibTeX files via one of the following
operations -diff, -exclusive, -intersect or -join.
`

	examples = `
` + "```" + `
    bibmerge -join my-old-articles.bib my-recent-articles.bib
` + "```" + `

Combine to BibTeX files into one using join.
`
)

func main() {
	app := cli.NewCli(bibtex.Version)

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
	app.BoolVar(&mergeJoin, "join", false, "join two bib files")
	app.BoolVar(&mergeDiff, "diff", false, "take the difference (asymmetric) between two bib files")
	app.BoolVar(&mergeIntersect, "intersect", false, "generate a bib listing from the intersection of two bib files")
	app.BoolVar(&mergeExclusive, "exclusive", false, "generate a symmetric difference between two bib files")

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
		listA []*bibtex.Element
		listB []*bibtex.Element
		listC []*bibtex.Element
	)

	if len(args) != 2 {
		fmt.Fprintf(app.Eout, "Must include two BibTeX filenames, try bibmerge -h for details\n")
		os.Exit(1)
	}
	src, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(app.Eout, "Can't read %s, %s", args[0], err)
		os.Exit(1)
	}
	listA, err = bibtex.Parse(src)
	if err != nil {
		fmt.Fprintf(app.Eout, "Can't parse %s, %s", args[0], err)
	}
	src, err = ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Fprintf(app.Eout, "Can't read %s, %s", args[1], err)
		os.Exit(1)
	}
	listB, err = bibtex.Parse(src)
	if err != nil {
		fmt.Fprintf(app.Eout, "Can't parse %s, %s", args[1], err)
	}
	switch {
	case mergeJoin:
		listC = bibtex.Join(listA, listB)
	case mergeDiff:
		listC = bibtex.Diff(listA, listB)
	case mergeIntersect:
		listC = bibtex.Intersect(listA, listB)
	case mergeExclusive:
		listC = bibtex.Exclusive(listA, listB)
	default:
		fmt.Fprintf(app.Eout, "Missing type of merge operation, try bibmerge -h for details\n")
		os.Exit(1)
	}
	for _, elem := range listC {
		fmt.Fprintf(app.Out, "%s\n", elem)
	}
	if newLine {
		fmt.Fprintf(app.Out, "\n")
	}
}
