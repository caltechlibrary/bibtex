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
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/tok"
)

const (
	// Version of BibTeX package
	Version = "v0.0.1"

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
	for _, ky := range element.Keys {
		out = append(out, fmt.Sprintf("    %s,", ky))
	}
	for ky, val := range element.Tags {
		out = append(out, fmt.Sprintf("    %s = %s,", ky, val))
	}

	out = append(out, fmt.Sprintf("}"))
	return strings.Join(out, "\n")
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
			val = append(val, []byte("{")[0])
			val = append(val[:], between[:]...)
			val = append(val, []byte("}")[0])
		case token.Type == tok.DoubleQuote:
			buf = tok.Backup(token, buf)
			between, buf, err = tok.Between([]byte("\""), []byte("\""), []byte(""), buf)
			if err != nil {
				return element, err
			}
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
				key = nil
			} else {
				// append to element keys
				keys = append(keys, string(val))
			}
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
