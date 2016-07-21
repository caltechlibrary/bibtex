//
// bibtex/scrape.go is a plain text scraping package related to creating BibTeX output
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
package scrape

import (
	"bytes"
	"regexp"
)

// Entry takes a a regexp finding start of record and a buffer []byte and returns entry
// as []byte, remaining buffer []byte
func Entry(re *regexp.Regexp, buf []byte) ([]byte, []byte) {
	start := re.FindIndex(buf)
	if start == nil {
		return []byte{}, []byte{}
	}
	i := start[1]
	if i < len(buf) {
		end := re.FindIndex(buf[i:])
		if end != nil {
			j := end[0]
			return buf[start[0] : i+j], buf[i+j:]
		}
	}
	return buf[start[0]:], []byte{}
}

func NextLine(buf []byte) ([]byte, []byte) {
	i := bytes.Index(buf, []byte("\n"))
	if i < 0 {
		return buf, []byte{}
	}
	return buf[0:i], buf[i:]
}

func Doi(buf []byte) []byte {
	re := regexp.MustCompile(`([dD][oO][iI] |[dD][oO][iI]:|)( +|)[0-9]+\.[0-9]+/([0-9]|[a-z]|[A-Z]|\.)+`)
	return bytes.TrimSpace(re.Find(buf))
}

func ISSN(buf []byte) []byte {
	re := regexp.MustCompile(`([iI][sS][sS][nN]|[eE][iI][sS][sS][nN]( +|))[0-9][0-9][0-9][0-9]-[0-9][0-9][0-9][0-9a-zA-Z]+`)
	return bytes.TrimSpace(re.Find(buf))
}

func PageRange(buf []byte) []byte {
	re := regexp.MustCompile(`[pP][pP]. [0-9]+(--|-)[0-9]+`)
	return re.Find(buf)
}

func PubYear(buf []byte) []byte {
	re := regexp.MustCompile(`\(([0-9][0-9][0-9][0-9]|[a-zA-Z]+ [0-9][0-9][0-9][0-9])\)`)
	return re.Find(buf)
}
