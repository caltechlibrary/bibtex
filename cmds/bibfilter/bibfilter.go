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
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	include = "comment,string,article,book,booklet,inbook,incollection,inproceedings,conference,manual,masterthesis,misc,phdthesis,proceedings,techreport,unpublished"
	exclude = ""
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	flag.BoolVar(&showLicense, "l", false, "display license")

	flag.StringVar(&include, "include", include, "a comma separated list of tags to include")
	flag.StringVar(&exclude, "exclude", exclude, "a comma separated list of tags to exclude")
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTION] BIBFILE

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

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
 
 * Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
 
 * Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.
 
 * Neither the name of mkbib nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.
 
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 
 `, appname)
		os.Exit(0)
	}

	var (
		err      error
		elements []*bibtex.Element
		buf      []byte
	)

	in := os.Stdin
	out := os.Stdout

	args := flag.Args()
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
