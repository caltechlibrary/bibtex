#!/bin/bash
echo "Generating website with shorthand"
echo "Rendering index.html"
shorthand -e "{{pageContent}} :[<: index.md" page.shorthand > index.html
echo "Rendering installation.html"
shorthand -e "{{pageContent}} :[<: INSTALL.md" page.shorthand > installation.html
echo "Rendering readme.html"
shorthand -e "{{pageContent}} :[<: README.md" page.shorthand > readme.html
echo "Rendering license.html"
shorthand -e "{{pageContent}} :[<: LICENSE" page.shorthand > license.html
