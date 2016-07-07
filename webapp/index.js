/*
 * index.js binds the form elements to access the bibfilter implemented with GopherJS
 */
(function (window, document) {
    'use strict';
    var inputTextArea = document.getElementById("input-bibtex"),
        outputTextArea = document.getElementById("output-bibtex"),
        includeInput = document.getElementById("include-bibtex"),
        excludeInput = document.getElementById("exclude-bibtex"),
        submitButton = document.getElementById("filter-bibtex");
        
    submitButton.addEventListener("click", function (evt) {
        var filter = bibfilter.New();
        outputTextArea.value = filter.Parse(inputTextArea.value, includeInput.value, excludeInput.value);
    });
    console.log(submitButton);
}(window, document));
