//
// Package bibtex is a quick and dirty BibTeX parser for working with
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
package bibtex

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	// Caltech Library packages
	"github.com/caltechlibrary/tok"
)

// TestBib tests the Bib tokenizer
func TestBib(t *testing.T) {
	fname1 := path.Join("testdata", "sample0.txt")
	fname2 := path.Join("testdata", "expected0.txt")

	src1, err := ioutil.ReadFile(fname1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	src2, err := ioutil.ReadFile(fname2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := strings.Split(strings.TrimSpace(string(src2)), "\n")
	var (
		token *tok.Token
		i     int
	)
	for i, expectedType := range expected {
		token, src1 = tok.Tok2(src1, Bib)
		if strings.Compare(token.Type, strings.TrimSpace(expectedType)) != 0 {
			t.Errorf("%d: %s != %s", i, token, expectedType)
		}
	}
	if len(src1) != 0 {
		t.Errorf("Expected to have len(src1) == 1, %d [%s]", i, src1)
	}
}

// TestParse tests the parsing function
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
	expectedTypes := []string{"comment", "string", "misc", "article", "article"}
	if len(elements) != len(expectedTypes) {
		t.Errorf("Expected 5 elements: %s\n", elements)
		t.FailNow()
	}
	for i, element := range elements {
		if i > len(expectedTypes) {
			t.Errorf("expectedTypes array shorter than required: %d\n", i)
			t.FailNow()
		}
		if element.Type != expectedTypes[i] {
			t.Errorf("expected %s, found %s", expectedTypes[i], element.Type)
		}
		if element.Type == "comment" {
			if len(element.Keys) != 3 {
				t.Errorf("Expected 3 keys in comment entry: %s", element)
			}
			if strings.HasPrefix(element.Keys[0], "\"") == true {
				t.Errorf("Expected first element to be a key: [%s]", element.Keys[0])
			}
			if strings.HasPrefix(element.Keys[1], "\"") == false {
				t.Errorf("Expected second elements to be quoted strings: [%s] [%s]", element.Keys[1], element)
			}
			if strings.HasPrefix(element.Keys[2], "\"") == false {
				t.Errorf("Expected third element to be quoted strings: [%s] [%s]", element.Keys[2], element)
			}
		} else if len(element.Tags) == 0 {
			t.Errorf("Expected tags in element: [%s]", element)
		}

	}
}
