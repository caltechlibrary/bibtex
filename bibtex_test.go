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
	"fmt"
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

// TestEquality tests Equal() and NotEqual()
func TestEquality(t *testing.T) {
	bib1 := new(Element)
	bib1.Tags = make(map[string]string)
	bib2 := new(Element)
	bib2.Tags = make(map[string]string)

	isTrue := func(expr bool, msg string, fail bool) {
		if expr == false {
			t.Error(msg)
			if fail == true {
				t.FailNow()
			}
		}
	}

	isTrue(Equal(bib1, bib2), "Empty bib1, bib2 should return IsEqual() true", true)

	bib1.Type = "string"
	isTrue(NotEqual(bib1, bib2), "bib1 has a type set, should not equal bib2", true)
	bib2.Type = "string"
	isTrue(Equal(bib1, bib2), "Should be equal again", true)

	bib1.Tags["author"] = "R. S. Doiel"
	isTrue(NotEqual(bib1, bib2), "bib1 has an author field now", true)
	bib2.Tags["author"] = "R. S. Doiel"
	isTrue(Equal(bib1, bib2), "bib2 should have an author field now", true)
}

// TestContains see if an Element is contained in an array of Elements
func TestContains(t *testing.T) {
	noError := func(err error, failNow bool) bool {
		if err != nil {
			t.Error(err)
			if failNow == true {
				t.FailNow()
			}
			return false
		}
		return true
	}
	isTrue := func(expr bool, msg string, fail bool) {
		if expr == false {
			t.Error(msg)
			if fail == true {
				t.FailNow()
			}
		}
	}

	fname := path.Join("testdata", "sample1.bib")
	src, err := ioutil.ReadFile(fname)
	noError(err, true)

	bibList, err := Parse(src)
	noError(err, true)
	elem := Clone(bibList[3])
	isTrue(Contains(bibList, elem), fmt.Sprintf("Should find bibList[3] with Contains, %s, %s", bibList, elem), true)
	elem.Type = "misc"
	isTrue(Contains(bibList, elem) == false, fmt.Sprintf("Should not find  a modified element with Contains, %s, %s", bibList, elem), true)
}

// TestJoin test joining to Element arrays without duplication
func TestJoin(t *testing.T) {
	noError := func(err error, failNow bool) bool {
		if err != nil {
			t.Error(err)
			if failNow == true {
				t.FailNow()
			}
			return false
		}
		return true
	}
	fname1 := path.Join("testdata", "sample1.bib")
	fname2 := path.Join("testdata", "sample2.bib")
	src1, err := ioutil.ReadFile(fname1)
	noError(err, true)
	src2, err := ioutil.ReadFile(fname2)
	noError(err, true)

	bibList1, err := Parse(src1)
	noError(err, true)
	bibList2, err := Parse(src2)
	noError(err, true)

	// Join an overlapping list
	bibList3 := Join(bibList1, bibList2)
	if len(bibList2) != len(bibList3) {
		t.Errorf("Expected bibList3 to be length of bibList2 - %s, %s", bibList2, bibList3)
		t.FailNow()
	}

	// Test where two bib records have different formatting but same content.
	src1 = []byte(`@article{12034,author={Robert}}`)
	src2 = []byte(`@article{12034,author="Robert"}`)
	bibList1, err = Parse(src1)
	bibList2, err = Parse(src2)
	bibList3 = Join(bibList1, bibList2)
	if len(bibList3) != 1 {
		t.Errorf("1. Expected a single bib record - %s, %s, %s", bibList1, bibList2, bibList3)
	}

	src1 = []byte(`@article{12034,
author={Robert}
}`)
	src2 = []byte(`@article{12034,
	author="Robert"
}`)
	bibList1, err = Parse(src1)
	bibList2, err = Parse(src2)
	bibList3 = Join(bibList1, bibList2)
	if len(bibList3) != 1 {
		t.Errorf("2. Expected a single bib record - %s, %s, %s", bibList1, bibList2, bibList3)
	}

	src1 = []byte(`@article{12034,author="Robert"}`)
	src2 = []byte(`
@article{12034,
author="Robert"
}
`)
	bibList1, err = Parse(src1)
	bibList2, err = Parse(src2)
	if strings.Compare(bibList1[0].Type, bibList2[0].Type) != 0 {
		t.Errorf("Types don't match %s != %s", bibList1[0].Type, bibList2[0].Type)
	}
	if strings.Compare(strings.Join(bibList1[0].Keys, ""), strings.Join(bibList2[0].Keys, "")) != 0 {
		t.Errorf("Keys don't match %s != %s", bibList1[0].Keys, bibList2[0].Keys)
	}
	if len(bibList1[0].Tags) != len(bibList2[0].Tags) {
		t.Errorf("Tags don't match\nsrc1:\n%s\n|-> %+v\nsrc2:\n%s\n|-> %+v\n", src1, bibList1[0].Tags, src2, bibList2[0].Tags)
	}
	bibList3 = Join(bibList1, bibList2)
	if len(bibList3) != 1 {
		t.Errorf("3. Expected a single bib record - %s, %s, %s", bibList1, bibList2, bibList3)
	}
}

// TestDiff test taking the difference between Element lists
func TestDiff(t *testing.T) {
	noError := func(err error, failNow bool) bool {
		if err != nil {
			t.Error(err)
			if failNow == true {
				t.FailNow()
			}
			return false
		}
		return true
	}
	isTrue := func(expr bool, msg string, fail bool) {
		if expr == false {
			t.Error(msg)
			if fail == true {
				t.FailNow()
			}
		}
	}
	fname1 := path.Join("testdata", "sample1.bib")
	fname2 := path.Join("testdata", "sample2.bib")
	src1, err := ioutil.ReadFile(fname1)
	noError(err, true)
	src2, err := ioutil.ReadFile(fname2)
	noError(err, true)

	elemList1, err := Parse(src1)
	noError(err, true)
	elemList2, err := Parse(src2)
	noError(err, true)

	elemList3 := Diff(elemList1, elemList2)
	isTrue(len(elemList3) == 0, fmt.Sprintf("elemList3 should be zero - %s, %s, %s", elemList1, elemList2, elemList3), true)
	elemList3 = Diff(elemList2, elemList1)
	isTrue(len(elemList3) == 1, fmt.Sprintf("len(elemList3) should be 1 - %s, %s, %s", elemList1, elemList2, elemList3), true)
}

// TestExclusive for symetrical differences of Element Lists
func TestExclusive(t *testing.T) {
	noError := func(err error, failNow bool) bool {
		if err != nil {
			t.Error(err)
			if failNow == true {
				t.FailNow()
			}
			return false
		}
		return true
	}
	isTrue := func(expr bool, msg string, fail bool) {
		if expr == false {
			t.Error(msg)
			if fail == true {
				t.FailNow()
			}
		}
	}
	fname1 := path.Join("testdata", "sample1.bib")
	fname2 := path.Join("testdata", "sample2.bib")
	src1, err := ioutil.ReadFile(fname1)
	noError(err, true)
	src2, err := ioutil.ReadFile(fname2)
	noError(err, true)

	elemList1, err := Parse(src1)
	noError(err, true)
	elemList2, err := Parse(src2)
	noError(err, true)

	elemList3 := Exclusive(elemList1, elemList2)
	isTrue(len(elemList3) == 1, fmt.Sprintf("Exclusive set (A xor B) should be len 1 - %s, %s, %s", elemList1, elemList2, elemList3), true)
	elemList3 = Exclusive(elemList2, elemList1)
	isTrue(len(elemList3) == 1, fmt.Sprintf("Exclusive set (B xor A) should be len 1 - %s, %s, %s", elemList1, elemList2, elemList3), true)

	fname1 = path.Join("testdata", "sample3a.bib")
	fname2 = path.Join("testdata", "sample3b.bib")
	src1, err = ioutil.ReadFile(fname1)
	noError(err, true)
	src2, err = ioutil.ReadFile(fname2)
	noError(err, true)

	elemList1, err = Parse(src1)
	noError(err, true)
	elemList2, err = Parse(src2)
	noError(err, true)

	/*
	   Should not get a missing @book like:

	   @comment{
	   	    id0,
	   		"This is some sort of comment",
	   		"Yet another comment line",
	   }
	*/
	elemList3 = Exclusive(elemList1, elemList2)
	isTrue(len(elemList3) == 2, fmt.Sprintf("Exclusive (A xor B) should be len 2 -\n%s\n\n%s\n\n%s\n", elemList1, elemList2, elemList3), true)
}
