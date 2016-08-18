
# Experimenting with a bibtexParseJS drop in replacement

There are some issues in the [ORCID bibtexParserJS](https://github.com/ORCID/bibtexParserJS) library raised in the 
the ORCID support forums

    http://support.orcid.org/forums/175591-orcid-ideas-forum/suggestions/6822051-bibtex-import-is-flaky

I it possible that github.com/caltechlibrary/bibtex may solve some of those issues and this directory is an experiment
in create a drop in replacement for the ORCID bibtexParserJS object via cross compile from Go.

## Building

+ bibtexParse.go will map the bibtex Go function and expose them as a bibtex object
+ wrapper.js will provide any additional wrapping to make things align well
+ these will be concatenated to render the bibtexParseJS.js replacement

## Testing

+ Running under NodeJS we should be able to pass all the same tests as currently listed in ORCID's repository
+ Some of these tests should be migrated into the core testing under Go to ensure correctness and compatibility

## Reference

+ [ORCID bibtexParseJS](https://github.com/ORCID/bibtexParseJS)
+ [bibtex-parser](https://github.com/mikolalysenko/bibtex-parser)
+ [bibtex-js](http://home.in.tum.de/~muehe/bibtex-js/src/bibtex_js.js)
    + [demo](http://home.in.tum.de/~muehe/bibtex-js/demo/bibtex.html)

