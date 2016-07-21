/*
 * bibmerge.js binds the form elements to access the bibtex object implemented with GopherJS
 */
(function (window, document) {
    'use strict';
    var inputTextArea = document.getElementById("input-source"),
        outputTextArea = document.getElementById("output-bibtex"),
        entrySeparator = document.getElementById("entry-separator"),
        submitButton = document.getElementById("scrape"),
        cmdExample = document.getElementById("cmd-example-bibtex");
        
    function splitEntries(src, sep) {
        var re,
            i = 0;

        re = new RegExp(sep, "g");
        return src.split(re);
    }

    
    submitButton.addEventListener("click", function (evt) {
        var scrape = bibtex.New(),
            cmd = [],
            out = [],
            separator = "",
            src = "",
            entries = [],
            entry = "";

        entry = "    ";
        separator = entrySeparator.value;
        src = "    " + inputTextArea.value;
        entries = splitEntries(src, separator)
        entries.forEach(function(item) {
            if (item.trim() != "") {
                console.log("DEBUG item", item);
                out.push(scrape.Scrape(item));
            }
        });

        outputTextArea.value = out.join("\n\n") 
        if (cmdExample) {
            cmd.push("bibscrape");
            cmd.push("-e '" + separator + "'");
            cmd.push("example.txt");
            cmd.push("> example.bib");
            cmdExample.value = cmd.join(" ");
        }
    });
}(window, document));
