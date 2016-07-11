
# bibtex

A Golang BibTeX package and *bibfilter* tool. *bibfilter* can parse and filter BibTeX entries.

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

