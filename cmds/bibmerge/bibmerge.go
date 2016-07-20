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
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	mergeJoin      bool
	mergeDiff      bool
	mergeIntersect bool
	mergeExclusive bool
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	flag.BoolVar(&showLicense, "l", false, "display license")

	flag.BoolVar(&mergeJoin, "join", false, "join two bib files")
	flag.BoolVar(&mergeDiff, "diff", false, "take the difference (asymmetric) between two bib files")
	flag.BoolVar(&mergeIntersect, "intersect", false, "generate a bib listing from the intersection of two bib files")
	flag.BoolVar(&mergeExclusive, "exclusive", false, "generate a symmetric difference between two bib files")
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTION] BIBFILE1 BIBFILE2

 OPTIONS:

`, appname)

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s %s\n", f.Name, f.Usage)
		})

		fmt.Printf("\n\n Version %s\n", bibtex.Version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf(" Version %s\n", bibtex.Version)
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Printf(`
 %s
 
 copyright (c) 2016, Caltech
 All rights not granted herein are expressly reserved by Caltech.

 Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
 
 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 
 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 
 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`, appname)
		os.Exit(0)
	}

	var (
		err   error
		listA []*bibtex.Element
		listB []*bibtex.Element
		listC []*bibtex.Element
	)

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Must include two BibTeX filenames, try %s -h for details", appname)
		os.Exit(1)
	}
	src, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s", args[0], err)
		os.Exit(1)
	}
	listA, err = bibtex.Parse(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse %s, %s", args[0], err)
	}
	src, err = ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read %s, %s", args[1], err)
		os.Exit(1)
	}
	listB, err = bibtex.Parse(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse %s, %s", args[1], err)
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
		fmt.Fprintf(os.Stderr, "Missing type of merge operation, try %s -h for details", appname)
		os.Exit(1)
	}
	for _, elem := range listC {
		fmt.Fprintf(os.Stdout, "%s\n", elem)
	}
}
