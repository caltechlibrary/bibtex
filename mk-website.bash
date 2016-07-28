#!/bin/bash

function softwareCheck() {
    for CMD in shorthand; do
        APP=$(which $CMD)
        if [ "$APP" = "" ]; then
            echo "Skipping, missing $CMD"
            exit 1;
        fi 
    done
}

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html from $content and $nav"
    shorthand \
        -e "{{navContent}} :import-markdown: $nav" \
        -e "{{pageContent}} :import-markdown: $content" \
        page.shorthand > $html
}

softwareCheck
echo "Generating website with shorthand"
mkPage nav.md index.md index.html
mkPage nav.md README.md readme.html
mkPage nav.md INSTALL.md installation.html
mkPage nav.md LICENSE license.html
