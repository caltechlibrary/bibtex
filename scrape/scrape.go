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
	"fmt"
	"regexp"

	// Caltech Library Packages
	"github.com/caltechlibrary/bibtex"
)

var (
	// LineSeparator a regexp describing common line separation
	LineSeparator = regexp.MustCompile(`(\n|\r\n)`)
	// DefaultEntrySeparator a regexp describing common entry separator
	DefaultEntrySeparator = regexp.MustCompile(`(\n\n|\r\n\r\n)`)
)

// DOI extract the next DOI found in []byte or return an empty []byte
func DOI(buf []byte) []byte {
	re := regexp.MustCompile(`([dD][oO][iI] |[dD][oO][iI]:|)( +|)[0-9]+\.[0-9]+/([0-9]|[a-z]|[A-Z]|\.)+`)
	return bytes.TrimSpace(re.Find(buf))
}

// ISSN extract the next ISSN found in []byte or return an empty []byte
func ISSN(buf []byte) []byte {
	re := regexp.MustCompile(`([iI][sS][sS][nN]|[eE][iI][sS][sS][nN])( +|:|)([0-9][0-9][0-9][0-9])-([0-9][0-9][0-9][0-9a-zA-Z])+`)
	return bytes.TrimSpace(re.Find(buf))
}

// PageRange extrasct the next page range found in []byte or return an empty []byte
func PageRange(buf []byte) []byte {
	re := regexp.MustCompile(`([pP][pP]|[pP][pP]\.|[pP][gG]|[pP][gG]\.|[pP][gG][sS]|[pP][gG][sS]\.|[pP][aA][gG][eE]|[pP][aA][gG][eE][sS])( +)[0-9]+(--|-| - | -- )[0-9]+`)
	return re.Find(buf)
}

// Year extract the next year found in []byte or return an empty []byte
func Year(buf []byte) []byte {
	re := regexp.MustCompile(`\(([0-9][0-9][0-9][0-9]|[a-zA-Z]+ [0-9][0-9][0-9][0-9])\)`)
	return re.Find(buf)
}

// NextEntry takes a buffer of []byte a Regular expression for splitting the plain text entries
// and returns the next entry and a remainder buffer both of type []byte
func NextEntry(buf []byte, re *regexp.Regexp) ([]byte, []byte) {
	loc := re.FindIndex(buf)
	if loc == nil {
		return buf, nil
	}
	return buf[0:loc[0]], buf[loc[1]:]
}

// Scrape takes a buffer of []byte and detects tags based on new lines encountered and
// specific elements like page ranges, years and month/year phrases,
// returns a new bibtex.Element
func Scrape(entry []byte) *bibtex.Element {
	elem := new(bibtex.Element)
	elem.Type = "pseudo"
	elem.Tags = make(map[string]string)
	// Find year
	year := Year(entry)
	if len(year) > 0 {
		elem.Tags["year"] = fmt.Sprintf("%q", year)
		entry = bytes.Replace(entry, year, []byte("\n"), 1)
	}
	issn := ISSN(entry)
	if len(issn) > 0 {
		elem.Tags["issn"] = fmt.Sprintf("%q", issn)
		entry = bytes.Replace(entry, issn, []byte("\n"), 1)
	}
	doi := DOI(entry)
	if len(doi) > 0 {
		elem.Tags["doi"] = fmt.Sprintf("%q", doi)
		entry = bytes.Replace(entry, doi, []byte("\n"), 1)
	}

	for i, val := range bytes.Split(entry, []byte("\n")) {
		val = bytes.TrimSpace(val)
		if len(val) > 0 {
			elem.Tags[fmt.Sprintf("unknown%d", i)] = fmt.Sprintf("%q", val)
		}
	}
	return elem
}
