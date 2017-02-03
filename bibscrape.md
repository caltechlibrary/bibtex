
 USAGE bibscrape [OPTIONS] FILENAME

 Parse the plain text file for BibTeX entry making a best guess
 to generate pseudo bib entries that can import into JabRef for
 cleanup.

 OPTIONS

   -e Set the default entry separator (defaults to \n\n)
   -h display help
   -k add a missing key
   -l display license
   -t Set the entry type  (defaults to pseudo)
   -v display version
 

 EXAMPLE
 	
	Use an 4 digit ID number and period to indicate the start of my bib
	records.

	    bibscrape -entry-separator "[0-9][0-9]0-9][0-9]\.\n" mytest.bib

	
 Version v0.0.8

