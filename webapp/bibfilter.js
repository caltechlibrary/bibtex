/*
 * bibfilter.js binds the form elements to access the bibtex object implemented with GopherJS
 */
(function (window, document) {
    'use strict';
    var inputTextArea = document.getElementById("input-bibtex"),
        outputTextArea = document.getElementById("output-bibtex"),
        includeInput = document.getElementById("include-bibtex"),
        excludeInput = document.getElementById("exclude-bibtex"),
        submitButton = document.getElementById("filter-bibtex"),
        cmdExample = document.getElementById("cmd-example-bibtex");
        
    submitButton.addEventListener("click", function (evt) {
        var filter = bibtex.New(),
            cmd = [];
        outputTextArea.value = filter.Parse(inputTextArea.value, includeInput.value, excludeInput.value);
        if (cmdExample) {
            cmd.push("bibfilter")
            if (includeInput.value.trim() !== "") {
                cmd.push(['-include="', includeInput.value.trim(), '"'].join(""))
            }
            if (excludeInput.value.trim() !== "") {
                cmd.push(['-exclude="', excludeInput.value.trim(), '"'].join(""))
            }
            cmd.push("example.bib")
            cmdExample.value = cmd.join(" ")
        }
    });
}(window, document));
