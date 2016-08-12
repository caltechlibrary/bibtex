
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

Symmetric versus asymmetric (not sure this really makes sense, maybe it is not even needed)

1. Element of A not in list B (asymmetric)
2. (Element of A not in B) union (Element of B not in A) (symmetric)

Some how I think my boolean logic is faulty...


