
# USAGE

    bibfilter [OPTION] BIBFILE

## SYSNOPSIS

bibfilter filters BibTeX files by entry type.

## OPTIONS

```
	-exclude	a comma separated list of tags to exclude
	-h	display help information
	-include	a comma separated list of tags to include
	-l	display license
	-v	display version information
```

## EXAMPLES
	
```
	bibfilter -include author,title my-works.bib
```

Renders a BibTeX file containing only author and title from my-works.bib

