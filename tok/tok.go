//
// Package tok is a niave tokenizer
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
package tok

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
)

const (
	// Version of  tok package
	Version = `v0.0.2`

	//Base token types, these are single letter tokens, no look ahead needed

	// Letter is an alphabetical letter (e.g. A-Z, a-z in English)
	Letter = "Letter"
	// Numeral is a single digit
	Numeral = "Numeral"
	// Punctuation is any non-number, non alphametical character, non-space (e.g. periods, colons, bang, hash mark)
	Punctuation = "Punctuation"
	// Space characters representing white space (e.g. space, tab, new line, carriage return)
	Space = "Space"

	// These are some common specialized token types provided for convientent.

	// Words a sequence of characters delimited by spaces
	Word = "Word"
	// OpenCurly bracket, e.g. "{"
	OpenCurlyBracket = "OpenCurlyBracket"
	// CloseCurly bracket, e.g. "}"
	CloseCurlyBracket = "CloseCurlyBracket"
	// CurlyBracket, e.g. "{}"
	CurlyBracket = "CurlyBracket"
	// OpenSquareBracket, e.g. "["
	OpenSquareBracket = "OpenSquareBracket"
	// CloseSquareBracket, e.g. "]"
	CloseSquareBracket = "CloseSquareBracket"
	// SquareBracket, e.g. "[]"
	SquareBracket = "SquareBracket"
	// OpenAngleBracket, e.g. "<"
	OpenAngleBracket = "OpenAngleBracket"
	// CloseAngleBracket, e.g. ">"
	CloseAngleBracket = "CloseAngleBracket"
	// AngleBracket, e.g. "<>"
	AngleBracket = "AngleBracket"
	// AtSign, e.g. "@"
	AtSign = "AtSign"
	// EqualSign, e.g. "="
	EqualSign = "EqualSign"
	// DoubleQuote, e.g. "\""
	DoubleQuote = "DoubleQuote"
	// SingleQuote, e.g., "'"
	SingleQuote = "SingleQuote"

	// EOF is an end of file token type. It is separate form Space only because of it being a common stop condition
	EOF = "EOF"
)

var (
	// Numerals is a map of numbers as strings
	Numerals = []byte("0123456789")

	// Spaces is a map space symbols as strings
	Spaces = []byte(" \t\r\n")

	// PunctuationMarks map as strings
	PunctuationMarks = []byte("~!@#$%^&*()_+`-=:{}|[]\\:;\"'<>?,./")

	// These map to the specialized tokens
	AtSignMark = []byte("@")
	// EqualMark, e.g. =
	EqualMark = []byte("=")
	// DoubleQuoteMark, e.g. "\""
	DoubleQuoteMark = []byte("\"")
	// SingleQuoteMark, e.g. "'"
	SingleQuoteMark = []byte("'")

	// OpenCurlyBrackets token
	OpenCurlyBrackets = []byte("{")
	// CloseCurlyBrackets token
	CloseCurlyBrackets = []byte("}")
	// CurlyBrackets tokens
	CurlyBrackets = []byte("{}")

	// OpenSquareBrackets token
	OpenSquareBrackets = []byte("[")
	// CloseSquareBrackets token
	CloseSquareBrackets = []byte("]")
	// SquareBrackets tokens
	SquareBrackets = []byte("[]")

	// OpenAngleBrackets token
	OpenAngleBrackets = []byte("<")
	// CloseAngleBrackets token
	CloseAngleBrackets = []byte(">")
	// AngleBrackets tokens
	AngleBrackets = []byte("<>")
)

// Token structure for emitting simply tokens and value from Tok() and Tok2()
type Token struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:"type" json:"type"`
	Value   []byte   `xml:"value" json:"value"`
}

// TokenMap is a map of simple token names and associated array of possible bytes
type TokenMap map[string][]byte

// Tokenizer is a function that takes a current token, looks ahead in []byte and returns a revised token and remaining []byte
type Tokenizer func(*Token, []byte) (*Token, []byte)

// String returns a human readable Token struct
func (t *Token) String() string {
	return fmt.Sprintf("{%q: %q}", t.Type, t.Value)
}

// IsSpace checks to see if []byte is a space or not
func IsSpace(b []byte) bool {
	for _, val := range b {
		if bytes.IndexByte(Spaces, val) == -1 {
			return false
		}
	}
	return true
}

// IsPunctuation checks to see if []byte is some punctuation or not
func IsPunctuation(b []byte) bool {
	for _, val := range b {
		if bytes.IndexByte(PunctuationMarks, val) == -1 {
			return false
		}
	}
	return true
}

// IsNumeral checks to see if []byte is a number or not
func IsNumeral(b []byte) bool {
	if bytes.Count(b, []byte(".")) > 1 {
		return false
	}
	for _, val := range b {
		if bytes.IndexByte([]byte("0123456789."), val) == -1 {
			return false
		}
	}
	return true
}

// Tok is a naive tokenizer that looks only at the next character by shifting it off the []byte and returning a token found with remaining []byte
func Tok(buf []byte) (*Token, []byte) {
	var (
		s []byte
	)
	if len(buf) == 0 {
		return &Token{
			Type:  EOF,
			Value: []byte(""),
		}, nil
	}
	s, buf = buf[0:1], buf[1:]
	switch {
	case IsPunctuation(s) == true:
		return &Token{
			Type:  Punctuation,
			Value: s,
		}, buf
	case IsSpace(s) == true:
		return &Token{
			Type:  Space,
			Value: s,
		}, buf
	case IsNumeral(s) == true:
		return &Token{
			Type:  Numeral,
			Value: s,
		}, buf
	default:
		return &Token{
			Type:  Letter,
			Value: s,
		}, buf
	}
}

// TokenFromMap, revaluates token type against a map of type names and byte arrays
// returns modified Token
func TokenFromMap(t *Token, m map[string][]byte) *Token {
	for k, v := range m {
		if bytes.Contains(v, t.Value) {
			return &Token{
				Type:  k,
				Value: t.Value,
			}
		}
	}
	return &Token{
		Type:  t.Type,
		Value: t.Value,
	}
}

// Tok2 provides an easy to implement look ahead tokenizer by defining a look ahead function
func Tok2(buf []byte, fn Tokenizer) (*Token, []byte) {
	tok, rest := Tok(buf)
	return fn(tok, rest)
}

// Skip provides a means to advance to the next non-target Token.
func Skip(tokenType string, buf []byte) ([]byte, *Token, []byte) {
	var (
		skipped []byte
		token   *Token
	)
	// Handle an empty buffer gracefully
	if len(buf) == 0 {
		token.Type = EOF
		token.Value = []byte("")
		return skipped, token, buf
	}
	for {
		token, buf = Tok(buf)
		if token.Type != tokenType {
			break
		}
		skipped = append(skipped[:], token.Value[:]...)
		if len(buf) == 0 {
			break
		}
	}

	return skipped, token, buf
}

func Skip2(tokenType string, buf []byte, fn Tokenizer) ([]byte, *Token, []byte) {
	var (
		skipped []byte
		token   *Token
	)
	// Handle an empty buffer gracefully
	if len(buf) == 0 {
		token.Type = EOF
		token.Value = []byte("")
		return skipped, token, buf
	}
	for {
		token, buf = Tok2(buf, fn)
		if token.Type != tokenType {
			break
		}
		skipped = append(skipped[:], token.Value[:]...)
		if len(buf) == 0 {
			break
		}
	}

	return skipped, token, buf
}

// Peek generates a token without consuming the buffer
func Peek(buf []byte) *Token {
	var (
		s []byte
	)
	if len(buf) == 0 {
		return &Token{
			Type:  EOF,
			Value: []byte(""),
		}
	}
	s = buf[0:1]
	switch {
	case IsPunctuation(s) == true:
		return &Token{
			Type:  Punctuation,
			Value: s,
		}
	case IsSpace(s) == true:
		return &Token{
			Type:  Space,
			Value: s,
		}
	case IsNumeral(s) == true:
		return &Token{
			Type:  Numeral,
			Value: s,
		}
	default:
		return &Token{
			Type:  Letter,
			Value: s,
		}
	}
}

// Between returns the buf between two delimiters (e.g. curly braces)
func Between(openValue []byte, closeValue []byte, escapeValue []byte, buf []byte) ([]byte, []byte, error) {
	var (
		between        []byte
		hasEscapeValue bool
		token          *Token
		isQuote        bool
	)

	isQuote = bytes.Equal(openValue, closeValue)
	if len(escapeValue) > 0 {
		hasEscapeValue = true
	}

	quoteCount := 0
	// Advance to start of Between token types
	if len(buf) == 0 {
		return between, buf, fmt.Errorf("missing opening %s", openValue)
	}
	for {
		if len(buf) == 0 {
			return between, buf, fmt.Errorf("missing closing %s", closeValue)
		}
		token, buf = Tok(buf)
		switch {
		case hasEscapeValue == true && bytes.Equal(escapeValue, token.Value):
			between = append(between[:], token.Value[:]...)
			token, buf = Tok(buf)
			between = append(between[:], token.Value[:]...)
		case isQuote == true && bytes.Equal(openValue, token.Value):
			if quoteCount == 0 {
				quoteCount++
			} else {
				quoteCount--
			}
			if quoteCount > 1 {
				between = append(between[:], token.Value[:]...)
			}
			if quoteCount == 0 {
				return between, buf, nil
			}
		case bytes.Equal(openValue, token.Value):
			quoteCount++
			if quoteCount > 1 {
				between = append(between[:], token.Value[:]...)
			}
		case bytes.Equal(closeValue, token.Value):
			quoteCount--
			if quoteCount == 0 {
				return between, buf, nil
			}
			between = append(between[:], token.Value[:]...)
		default:
			if quoteCount > 0 {
				between = append(between[:], token.Value[:]...)
			}
		}
	}

	return between, buf, nil
}

// Backup pushes a Token back onto the front of a Buffer
func Backup(token *Token, buf []byte) []byte {
	return append(token.Value[:], buf[:]...)
}

// Words is an example of implementing a Tokenizer function
func Words(tok *Token, buf []byte) (*Token, []byte) {
	if tok.Type == Letter || tok.Type == Word {
		// Get the next Token
		newTok, newBuf := Tok(buf)
		if newTok.Type == Letter {
			tok.Type = Word
			tok.Value = append(tok.Value, newTok.Value[0])
			tok, buf = Words(tok, newBuf)
		}
	}
	return tok, buf
}

// Next takes a buffer ([]byte) and a regular expression (string) and returns
// two []byte, first is the sub []byte until the expression is found or end of buf
// and the second is the remaining []byte array
func Next(buf []byte, re *regexp.Regexp) ([]byte, []byte) {
	loc := re.FindIndex(buf)
	if loc == nil {
		return buf, nil
	}
	return buf[0:loc[0]], buf[loc[1]:]
}

// NextLine takes a buffer ([]byte) and returns the next line as
// a []byte and the remainder as a []byte.
func NextLine(buf []byte) ([]byte, []byte) {
	return Next(buf, regexp.MustCompile(`(\n|\r\n)`))
}
