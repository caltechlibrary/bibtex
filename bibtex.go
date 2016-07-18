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
	"bytes"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/tok"
)

const (
	// Version of BibTeX package
	Version = "v0.0.4"

	// DefaultInclude list
	DefaultInclude = "comment,string,article,book,booklet,inbook,incollection,inproceedings,conference,manual,masterthesis,misc,phdthesis,proceedings,techreport,unpublished"

	// A template for printing an element
	ElementTmplSrc = `
@{{- .Type -}}{
    {{-range .Keys}}
	{{ . -}},
	{{end}}
	{{-range $key, $val := .Tags}}
		{{- $key -}} = {{- $val -}},
	{{end}}
}
`
)

// Generic Element
type Element struct {
	XMLName xml.Name          `json:"-"`
	Type    string            `xml:"type" json:"type"`
	Keys    []string          `xml:"keys" json:"keys"`
	Tags    map[string]string `xml:"tags" json:"tags"`
}
type Elements []*Element

type TagTypes struct {
	Required []string
	Optional []string
}

// Entry types
var (
	elementTypes = &map[string]*TagTypes{
		"article": &TagTypes{
			Required: []string{"author", "title", "journal", "year", "volume"},
			Optional: []string{"number", "pages", "month", "note"},
		},
		"book": &TagTypes{
			Required: []string{"author", "editor", "title", "publisher", "year"},
			Optional: []string{"volume", "number", "series", "address", "edition", "month", "note"},
		},
		"booklet": &TagTypes{
			Required: []string{"Title"},
			Optional: []string{"author", "howpublished", "address", "month", "year", "note"},
		},
		"inbook": &TagTypes{
			Required: []string{"author", "editor", "title", "chapter", "pages", "publisher", "year"},
			Optional: []string{"volume", "number", "series", "type", "address", "edition", "month", "note"},
		},
		"incollection": &TagTypes{
			Required: []string{"author", "title", "booktitle", "publisher", "year"},
			Optional: []string{"editor", "volume", "number", "series", "type", "chapter", "pages", "address", "edition", "month", "note"},
		},
		"inproceedings": &TagTypes{
			Required: []string{"author", "title", "booktitle", "year"},
			Optional: []string{"editor", "volume", "number", "series", "pages", "address", "month", "organization", "publisher", "note"},
		},
		"conference": &TagTypes{
			Required: []string{"author", "title", "booktitle", "year"},
			Optional: []string{"editor", "volume", "number", "series", "pages", "address", "month", "organization", "publisher", "note"},
		},
		"manual": &TagTypes{
			Required: []string{"title"},
			Optional: []string{"author", "organization", "address", "edition", "month", "year", "note"},
		},
		"masterthesis": &TagTypes{
			Required: []string{"author", "title", "school", "year"},
			Optional: []string{"type", "address", "month", "note"},
		},
		"misc": &TagTypes{
			Required: []string{},
			Optional: []string{"author", "title", "howpublished", "month", "year", "note"},
		},
		"phdthesis": &TagTypes{
			Required: []string{"author", "title", "school", "year"},
			Optional: []string{"type", "address", "month", "note"},
		},
		"proceedings": &TagTypes{
			Required: []string{"title", "year"},
			Optional: []string{"editor", "volume", "series", "address", "month", "publisher", "organization", "note"},
		},
		"techreport": &TagTypes{
			Required: []string{"author", "title", "institution", "year"},
			Optional: []string{"type", "number", "address", "month", "note"},
		},
		"unpublished": &TagTypes{
			Required: []string{"author", "title", "note"},
			Optional: []string{"month", "year"},
		},
	}
)

// Render a single BibTeX element
func (element *Element) String() string {
	var out []string

	out = append(out, fmt.Sprintf("@%s{", element.Type))
	if len(element.Keys) > 0 {
		for _, ky := range element.Keys {
			if len(ky) > 0 {
				out = append(out, fmt.Sprintf("%s,\n", ky))
			}
		}
	}

	if len(element.Tags) > 0 {
		for ky, val := range element.Tags {
			out = append(out, fmt.Sprintf("    %s = %s,\n", ky, val))
		}
	}

	out = append(out, fmt.Sprintf("}\n"))
	return strings.Join(out, "")
}

//
// Parser related structures
//

// Bib is a niave BibTeX Tokenizer function
// Note: there is an English bias in the AlphaNumeric check
func Bib(token *tok.Token, buf []byte) (*tok.Token, []byte) {
	switch {
	case token.Type == tok.AtSign || token.Type == "BibElement":
		// Get the next Token
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type != tok.OpenCurlyBracket {
			token.Type = "BibElement"
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	case token.Type == tok.Space:
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type == tok.Space {
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	case token.Type == tok.Letter || token.Type == tok.Numeral || token.Type == "AlphaNumeric":
		// Convert Letters and Numerals to AlphaNumeric Type.
		token.Type = "AlphaNumeric"
		// Get the next Token
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type == tok.Letter || newTok.Type == tok.Numeral {
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	default:
		// Revaluate token for more specific token types.
		token = tok.TokenFromMap(token, map[string][]byte{
			tok.OpenCurlyBracket:  tok.OpenCurlyBrackets,
			tok.CloseCurlyBracket: tok.CloseCurlyBrackets,
			tok.AtSign:            tok.AtSignMark,
			tok.EqualSign:         tok.EqualMark,
			tok.DoubleQuote:       tok.DoubleQuoteMark,
			tok.SingleQuote:       tok.SingleQuoteMark,
			"Comma":               []byte(","),
		})
	}

	return token, buf
}

func mkElement(elementType string, buf []byte) (*Element, error) {
	var (
		key     []byte
		val     []byte
		between []byte
		token   *tok.Token
		err     error
		keys    []string
		tags    map[string]string
	)

	element := new(Element)
	element.Type = elementType
	tags = make(map[string]string)

	for {
		if len(buf) == 0 {
			if len(key) > 0 {
				// We have a trailing key/value pair to save.
				tags[string(key)] = string(val)
			} else if len(val) > 0 {
				// We have a trailing key to save.
				keys = append(keys, string(val))
			}
			break
		}
		_, token, buf = tok.Skip2(tok.Space, buf, Bib)
		switch {
		case token.Type == tok.OpenCurlyBracket:
			buf = tok.Backup(token, buf)
			between, buf, err = tok.Between([]byte("{"), []byte("}"), []byte(""), buf)
			if err != nil {
				return element, err
			}
			// Non-destructively copy the quote into val
			val = append(val, []byte("{")[0])
			val = append(val[:], between[:]...)
			val = append(val, []byte("}")[0])
		case token.Type == tok.DoubleQuote:
			buf = tok.Backup(token, buf)
			between, buf, err = tok.Between([]byte("\""), []byte("\""), []byte(""), buf)
			if err != nil {
				return element, err
			}
			// Non-destructively copy the quote into val
			val = append(val, []byte("\"")[0])
			val = append(val[:], between[:]...)
			val = append(val, []byte("\"")[0])
		case token.Type == tok.EqualSign:
			key = val
			val = nil
		case token.Type == "Comma" || len(buf) == 0:
			if len(key) > 0 {
				//make a map entry
				tags[string(key)] = string(val)
			} else {
				// append to element keys
				keys = append(keys, string(val))
			}
			key = nil
			val = nil
		case token.Type == tok.Punctuation && bytes.Equal(token.Value, []byte("#")):
			val = append(val[:], []byte(" # ")[:]...)
		default:
			val = append(val[:], token.Value[:]...)
		}
	}
	element.Keys = keys
	element.Tags = tags
	return element, nil
}

// Parse a BibTeX file into appropriate structures
func Parse(buf []byte) ([]*Element, error) {
	var (
		lineNo      int
		token       *tok.Token
		elements    []*Element
		err         error
		skipped     []byte
		entrySource []byte
		LF          = []byte("\n")
	)
	lineNo = 1
	for {
		if len(buf) == 0 {
			break
		}
		skipped, token, buf = tok.Skip2(tok.Space, buf, Bib)
		lineNo = lineNo + bytes.Count(skipped, LF)
		if token.Type == tok.AtSign {
			// We may have a entry key
			token, buf = tok.Tok2(buf, Bib)
			if token.Type == "AlphaNumeric" {
				elementType := token.Value[:]
				skipped, token, buf = tok.Skip2(tok.Space, buf, Bib)
				lineNo = lineNo + bytes.Count(skipped, LF)
				if token.Type == tok.OpenCurlyBracket {
					// Ok it looks like we have a Bib entry now.
					buf = tok.Backup(token, buf)
					entrySource, buf, err = tok.Between([]byte("{"), []byte("}"), []byte(""), buf)
					if err != nil {
						return elements, fmt.Errorf("Problem parsing entry at %d", lineNo)
					}
					// OK, we have an entry, let's process it.
					element, err := mkElement(string(elementType), entrySource)
					if err != nil {
						return elements, fmt.Errorf("Error parsing element at %d, %s", lineNo, err)
					}
					lineNo = lineNo + bytes.Count(entrySource, LF)
					// OK, we have an element, let's append to our array...
					elements = append(elements, element)
				}
			}
		}
	}
	if len(elements) == 0 {
		err = fmt.Errorf("no elements found")
	}
	return elements, nil
}

// ByKey struct is for sorting Element Keys
type ByKey []string

// Len of ByKey array
func (a ByKey) Len() int {
	return len(a)
}

// Swap of ByKey array elements
func (a ByKey) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less return the lesser of ByKey array elements
func (a ByKey) Less(i, j int) bool {
	return strings.Compare(a[i], a[j]) < 0
}

func compareTagValues(val1, val2 string) bool {
	if strings.Compare(val1, val2) == 0 {
		return true
	}
	if len(val1) > 2 && len(val2) > 2 {
		// Drop the quoting char and compare the string.
		i1 := len(val1) - 1
		i2 := len(val2) - 1
		if strings.Compare(val1[1:i1], val2[1:i2]) == 0 {
			return true
		}
	}
	return false
}

// Equal compares two Element structures and sees if the contents agree
func Equal(elem1, elem2 *Element) bool {
	if strings.Compare(elem1.Type, elem2.Type) != 0 {
		return false
	}
	// We have differing number of keys or Tags then we're not equal
	if len(elem1.Keys) != len(elem2.Keys) || len(elem1.Tags) != len(elem2.Tags) {
		return false
	}

	// Sort and compare the keys
	keys1 := elem1.Keys[0:]
	keys2 := elem2.Keys[0:]
	sort.Sort(ByKey(keys1))
	sort.Sort(ByKey(keys2))

	// We have to find
	for i, ky := range keys1 {
		if strings.Compare(keys2[i], ky) != 0 {
			return false
		}
	}

	for ky, val1 := range elem1.Tags {
		if val2, ok := elem2.Tags[ky]; ok != true {
			return false
		} else if compareTagValues(val1, val2) == false {
			return false
		}
	}

	return true
}

// NotEqual compares two element structures and see if the contents disagree
func NotEqual(elem1, elem2 *Element) bool {
	return Equal(elem1, elem2) == false
}

// Clone creates a new Element based on an existing element
func Clone(elem *Element) *Element {
	newElem := new(Element)
	newElem.XMLName = elem.XMLName
	newElem.Type = elem.Type
	newElem.Tags = make(map[string]string)
	for _, ky := range elem.Keys {
		newElem.Keys = append(newElem.Keys, ky)
	}
	for ky, val := range elem.Tags {
		newElem.Tags[ky] = val
	}
	return newElem
}

// Contains checks an array of Elements for a specific element
func Contains(elemList []*Element, target *Element) bool {
	for _, elem := range elemList {
		if Equal(elem, target) == true {
			return true
		}
	}
	return false
}

// Join create a new Element array by combining to Element arrays without creating duplicate entries
func Join(elemList1, elemList2 []*Element) []*Element {
	var result []*Element
	result = elemList1[0:]
	for _, elem := range elemList2 {
		if Contains(result, elem) == false {
			result = append(result, elem)
		}
	}
	return result
}

// Diff creates a new Element Array of all the elements in elemList1 an not in elemList2
func Diff(elemList1, elemList2 []*Element) []*Element {
	var result []*Element
	for _, elem := range elemList1 {
		if Contains(elemList2, elem) == false {
			result = append(result, elem)
		}
	}
	return result
}

// Intersect create a new Element Array of elements in both elemList1 and elemList2
func Intersect(elemList1, elemList2 []*Element) []*Element {
	var result []*Element
	for _, elem := range elemList1 {
		if Contains(elemList2, elem) == true {
			result = append(result, elem)
		}
	}
	return result
}

// Exclusive create a new Element Array with elements that only exist in elemList1 or elemList2
func Exclusive(elemList1, elemList2 []*Element) []*Element {
	A := Diff(elemList1, elemList2)
	B := Diff(elemList2, elemList1)
	return Join(A, B)
}
