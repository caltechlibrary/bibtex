/*
 * bibmerge.js binds the form elements to access the bibtex object implemented with GopherJS
 */
(function (window, document) {
    'use strict';
    var inputTextArea = document.getElementById("input-source"),
        outputTextArea = document.getElementById("output-bibtex"),
        entryType = document.getElementById("entry-type"),
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
            separator = "\n\n",
            entry_type = "pseudo",
            src = "",
            entries = [],
            entry = "";

        // Update the default values if needed
        if (entrySeparator.value != undefined && entrySeparator.value !== "") {
            separator = entrySeparator.value;
        }
        if (entryType.value != undefined && entryType.value !== "") {
            entry_type = entryType.value;
        }

        src = inputTextArea.value;
        entries = splitEntries(src, separator)
        entries.forEach(function(item, i) {
            if (item.trim() !== "") {
                out.push(scrape.Scrape(item, entry_type, "pseudo_id_" + i));
            }
        });

        outputTextArea.value = out.join("\n\n") 
        if (cmdExample) {
            cmd.push("bibscrape");
            cmd.push("-t " + entry_type);
            cmd.push("-e '" + separator + "'");
            cmd.push("-k");
            cmd.push("example.txt");
            cmd.push("> example.bib");
            cmdExample.value = cmd.join(" ");
        }
    });
}(window, document));
