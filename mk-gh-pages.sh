#!/bin/bash

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html from $content and $nav"
    shorthand \
        -e "{{navContent}} :[<: $nav" \
        -e "{{pageContent}} :[<: $content" \
        page.shorthand > $html
}
echo "Generating website with shorthand"
mkPage nav.md index.md index.html
mkPage nav.md README.md readme.html
mkPage nav.md INSTALL.md installation.html
mkPage nav.md LICENSE license.html
