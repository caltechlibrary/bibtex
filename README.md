
# bibtex

A Golang BibTeX package and collection of related command line utilities.

[bibtex](https://github.com/caltechlibrary/bibtex) is a golang package for working with BibTeX files. 
It includes *bibfilter* and *bibmerge* which are a command line utilities for working with BibTeX files
(e.g. removing comments before importing into JabRef, merge bib lists).  *bibtex* also can be compiled to 
JavaScript via GopherJS. A web version of *bibfiter* and *bibmerge* commands is available in the 
[webapp](webapp/) directory. The command line utilities and webapp use
the same *bibtex* golang package for implementation.


## bibfilter

Output _my.bib_ file without comment entries

```
  bibfilter -exclude=comment my.bib
```

Output only articles from _my.bib_

```
    bibfilter -include=article my.bib
```

Output only articles and conference proceedings from _my.bib_

```
    bibfilter -include=article,inproceedings my.bib
```

## bibmerge

Output a new bibtex file based on the contents of two other bibtex files.

Join of two bibtex files

```
    bibmerge -join mybib1.bib mybib2.bib
```

Difference (asymmetric, set A - B may not equal set B - A) of two bibtex files

```
    bibmerge -diff mybib1.bib mybib2.bib
```

Intersection of two bibtex files

```
    bibmerge -intersect mybib1.bib mybib2.bib
```

Excluse difference (symmetric difference, inverse of intersection) of two bibtex files

```
    bibmerge -exclusive mybib1.bib mybib2.bib
```

Symmetric versus asymmetric (not sure this really makes sense in practice)

1. (asymmetric) A - B
2. (symmetric) (A - B) Union (B - A)


