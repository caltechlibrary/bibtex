
# bibtex

A Golang BibTeX package and collection of related command line utilities.

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

Union of two bibtex files

```
    bibmerge -union mybib1.bib mybib2.bib
```

Intersection of two bibtex

```
    bibmerge -intersect mybib1.bib mybib2.bib
```

Difference of two bibtex

```
    bibmerge -diff mybib1.bib mybib2.bib
```

