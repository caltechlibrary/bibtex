
# TODO

## Bugs

## Next Steps

+ [ ] Port GopherJS codebase to WASM, update for Go >= 1.11.5

## Someday, Maybe

+ Given an Array of JSON records a map (e.g. "title=.title", "journal=.journalName" "authors=.authors[:].familyName,.authors[:].givenName") render a BibTeX file to stdout
+ Get to the bottom of what I thinking regarding list diffs - asymmetric (A - B)  versus symmetric (A - B U B - A, am I thinking XOR versus OR???)
+ bib2bib take a bibtex file as input and render a new bib file with strings substitution, concatenation and comments stripped
+ bibjson2bib so we can easily import BibJSON into JabRef
+ text2bib which can import plain text formatted citation and output a BibTeX file
+ restapi2bib would be a tool that can talk to public APIs like CrossRef, DataCite and ORCID and render a BibTeX from the JSON API response
+ for web version bibfilter
    + generate the equavalent command line for current filter settings

## Completed

+ [x] migrate from caltechlibrary/tok to Golang native scanner package
