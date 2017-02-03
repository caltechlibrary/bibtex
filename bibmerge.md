
# USAGE

    bibmerge [OPTIONS] BIB_FILE1 BIB_FILE2

## SYNOPSIS

bibmerge will merge combine two BibTeX files via one of the following
operations -diff, -exclusive, -intersect or -join.

## OPTIONS

```
	-diff	take the difference (asymmetric) between two bib files
	-exclusive	generate a symmetric difference between two bib files
	-h	display help information
	-intersect	generate a bib listing from the intersection of two bib files
	-join	join two bib files
	-l	display license
	-v	display version information
```

## EXAMPLES

```
    bibmerge -join my-old-articles.bib my-recent-articles.bib
```

Combine to BibTeX files into one using join.
