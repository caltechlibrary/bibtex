//
// Package bibtex is a quick and dirty plain text parser for generating
// a Bibtex citation
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of mkbib nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package bibtex

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"
)

func TestParse(t *testing.T) {
	fname := path.Join("testdata", "sample1.bib")
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	elements, err := Parse(src)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	if len(elements) != 4 {
		t.Errorf("Expected 4 elements: %s\n", elements)
		t.FailNow()
	}
	expectedTypes := []string{"comment", "misc", "article", "article"}
	for i, element := range elements {
		fmt.Printf("DEBUG %d element: %s\n", i, element)
		if len(expectedTypes) > i {
			t.Errorf("expectedTypes array shorter than required: %d\n", i)
			t.FailNow()
		}
		if element.Type != expectedTypes[i] {
			t.Errorf("expected %s, found %s", expectedTypes[i], element.Type)
		}
	}
}