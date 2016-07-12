/*
 * bibmerge.js binds the form elements to access the bibtex object implemented with GopherJS
 */
(function (window, document) {
    'use strict';
    var inputTextAreaA = document.getElementById("inputA-bibtex"),
        inputTextAreaB = document.getElementById("inputB-bibtex"),
        outputTextArea = document.getElementById("output-bibtex"),
        operation = document.getElementById("merge-type-bibtex"),
        submitButton = document.getElementById("merge-bibtex"),
        cmdExample = document.getElementById("cmd-example-bibtex");
        
    submitButton.addEventListener("click", function (evt) {
        var merge = bibtex.New(),
            cmd = [],
            out = "",
            opOK = true;

        switch (operation.value) {
            case "join":
                outputTextArea.value = merge.Join(inputTextAreaA.value, inputTextAreaB.value);
                opOK = true;
                break;
            case "diff":
                outputTextArea.value = merge.Diff(inputTextAreaA.value, inputTextAreaB.value);
                opOK = true;
                break;
            case "intersect":
                outputTextArea.value = merge.Intersect(inputTextAreaA.value, inputTextAreaB.value);
                opOK = true;
                break;
            case "exclusive":
                outputTextArea.value = merge.Exclusive(inputTextAreaA.value, inputTextAreaB.value);
                opOK = true;
                break;
            default:
                outputTextArea.value = "unknown operation requested";
                opOK = false;
        }
        if (cmdExample && opOK === true) {
            cmd.push("bibmerge");
            cmd.push("-" + operation.value);
            cmd.push("example1.bib");
            cmd.push("example2.bib");
            cmdExample.value = cmd.join(" ");
        }
    });
}(window, document));
