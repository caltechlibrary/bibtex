//
// webapp.go a Web Application of bibtex package
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
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex"

	"github.com/gopherjs/gopherjs/js"
)

// This is where we hang our BibTeX object
type BibTeX struct {
}

func (b *BibTeX) Parse(buf, include, exclude string) string {
	var out []string
	elements, err := bibtex.Parse([]byte(buf))
	if err != nil {
		println(err)
	}
	if include == "" {
		include = bibtex.DefaultInclude
	}
	for _, element := range elements {
		if strings.Contains(include, element.Type) {
			if len(exclude) == 0 || strings.Contains(exclude, element.Type) == false {
				out = append(out, element.String())
			}
		}
	}
	return strings.Join(out, "\n")
}

func (b *BibTeX) Join(srcA, srcB string) string {
	var out []string
	listA, err := bibtex.Parse([]byte(srcA))
	if err != nil {
		println(err)
	}
	listB, err := bibtex.Parse([]byte(srcB))
	if err != nil {
		println(err)
	}
	listC := bibtex.Join(listA, listB)
	for _, element := range listC {
		out = append(out, element.String())
	}
	return strings.Join(out, "\n")
}

func (b *BibTeX) Diff(srcA, srcB string) string {
	var out []string
	listA, err := bibtex.Parse([]byte(srcA))
	if err != nil {
		println(err)
	}
	listB, err := bibtex.Parse([]byte(srcB))
	if err != nil {
		println(err)
	}
	listC := bibtex.Diff(listA, listB)
	for _, element := range listC {
		out = append(out, element.String())
	}
	return strings.Join(out, "\n")
}

func (b *BibTeX) Intersect(srcA, srcB string) string {
	var out []string
	listA, err := bibtex.Parse([]byte(srcA))
	if err != nil {
		println(err)
	}
	listB, err := bibtex.Parse([]byte(srcB))
	if err != nil {
		println(err)
	}
	listC := bibtex.Intersect(listA, listB)
	for _, element := range listC {
		out = append(out, element.String())
	}
	return strings.Join(out, "\n")
}

func (b *BibTeX) Exclusive(srcA, srcB string) string {
	var out []string
	listA, err := bibtex.Parse([]byte(srcA))
	if err != nil {
		println(err)
	}
	listB, err := bibtex.Parse([]byte(srcB))
	if err != nil {
		println(err)
	}
	listC := bibtex.Exclusive(listA, listB)
	for _, element := range listC {
		out = append(out, element.String())
	}
	return strings.Join(out, "\n")
}

func New() *js.Object {
	return js.MakeWrapper(&BibTeX{})
}

func main() {
	js.Global.Set("bibtex", map[string]interface{}{
		"New": New,
	})
}
