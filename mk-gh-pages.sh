#!/bin/bash
echo "Generating website with shorthand"
shorthand -e "{{pageContent}} :[<: index.md" page.shorthand > index.html
shorthand -e "{{pageContent}} :[<: INSTALL.md" page.shorthand > installation.html
shorthand -e "{{pageContent}} :[<: README.md" page.shorthand > readme.html
shorthand -e "{{pageContent}} :[<: LICENSE" page.shorthand > license.html
