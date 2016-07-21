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
	"testing"
)

func TestDoi(t *testing.T) {
	sample := []byte(`Blair, K., & Hoy, C. (2006). Paying attention to adult learning online: The pedagogy and politics of community. Computers and Composition, 23(1), 32-48. doi:10.1016/j.compcom.2005.12.006`)

	expected := []byte(`doi:10.1016/j.compcom.2005.12.006`)
	doi := Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}

	sample = []byte(`this has no DOI in it.`)
	expected = []byte(``)
	doi = Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}

	sample = []byte(`vlah bahl poerwpoi 'unicorn 10.1000/xyz000' opiewr lad`)
	expected = []byte(`10.1000/xyz000`)
	doi = Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}

	sample = []byte(` opiewn aslkds doi:10.1000/xyz000 opewirwer qw`)
	expected = []byte(`doi:10.1000/xyz000`)
	doi = Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}

	sample = []byte(`Bai, H. (2009). Facilitating students' critical thinking in online discussion: An instructor's experience. Journal of Interactive Online Learning, 8(2), 156-164. Retrieved from http://www.ncolr.org/jiol/`)
	expected = []byte(``)
	doi = Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}

	sample = []byte(`Manny, F. A. (1909). A study in adult education. The School Review, 17(3), 174-177. Retrieved from http://www.jstor.org/`)
	expected = []byte(``)
	doi = Doi(sample)
	if bytes.Compare(doi, expected) != 0 {
		t.Errorf("Expected %q, found %q", expected, doi)
	}
}

func TestISSN(t *testing.T) {
	var (
		sample   []byte
		expected []byte
		issn     []byte
	)

	OK := func(expected, found []byte) {
		if bytes.Compare(found, expected) != 0 {
			t.Errorf("Expected %q, found %q", expected, found)
		}
	}

	sample = []byte(`urn:ISSN:1534-0481`)
	expected = []byte(`ISSN:1534-0481`)
	issn = ISSN(sample)
	OK(expected, issn)
}

/*
func TestScrape(t *testing.T) {
	var (
		entry []byte
		buf   []byte
		err   error
	)

	sampleBuf := func(b []byte, l int) []byte {
		if l < len(b) {
			return b[0:l]
		}
		return b
	}

	// Source was http://www.wag.caltech.edu/publications/papers/ on 2016-07-20 at 12:06 PDT
	fname := path.Join("testdata", "goddard-sample1.txt")
	goddardSample1, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Scan buffer, find entries and convert to BibTeX records
	re := regexp.MustCompile(`[0-9]+\.\n`)
	buf = goddardSample1[:]
	i := 1
	for k := 0; k < len(goddardSample1); k++ {
		entry, buf = Entry(re, buf)
		fmt.Printf("%d: [%s] [%s]\n", i, sampleBuf(entry, 24), sampleBuf(buf, 24))
		if len(buf) == 0 || buf == nil {
			break
		}
		i++
	}
}
*/
