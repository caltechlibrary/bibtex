# USAGE

    bibscrape [OPTIONS] FILENAME

## SYSNOPSIS

bibscrape parses a plain text file for BibTeX entry making a best guess to generate pseudo bib entries that can import into JabRef for clea

## OPTIONS

```
	-e	Set the default entry separator (defaults to \n\n)
	-h	display help
	-k	add a missing key
	-l	display license
	-t	Set the entry type  (defaults to pseudo)
	-v	display version
```

## EXAMPLES

```
    bibscrape -entry-separator "[0-9][0-9]0-9][0-9]\.\n" mytest.bib
```
