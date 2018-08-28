[![Go Report Card](http://goreportcard.com/badge/caltechlibrary/tok)](http://goreportcard.com/report/caltechlibrary/tok)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)


# tok

A niave tokenizer library

## Public Interface

+ Backup - given a token and buffer return a new buffer with the token's value as prefix
    + parameters
        + Token
        + buffer (byte array)
    returns
        + buffer (byte array)
+ Between - returns the value between an opening and closing delimiter values, 
    + parameters
        + open value (byte array)
        + close value (byte array)
        + escape vaue (byte array)
        + buffer (byte array)
    + returns
        + between content (byte array)
        + buffer (byte array)
        + error value if closing value not found before end of buffer
+ Peek - returns the next token without consuming the buffer being scanned
    + parameters
        + buffer (byte array)
    + returns
        + Token
+ Skip - scans through a buffer until a token is found, returns skipped content, token and remaining buffer
    + parameters
        + Token
        + buffer (byte array)
    + returns
        + skipped content (byte array)
        + Token
        + buffer (byte array)
+ Skip2 - like Skip but allows a Tokenizer to be passed in rather than using the default Tok().
    + parameters
        + Token
        + buffer (byte array)
        + Tokenizer function
    + returns
        + skipped content (byte array)
        + Token
        + buffer (byte array)
+ Token - a simple structure 
    + properties
        + Type is a string holding the label of the token type
        + Value is a byte array holding the value of the token
+ Tokenizer - is a type of function that can be applied by Tok2, may be recursive
    + parameters
        + byte array
        + a Tokenizer function
    + returns
        + Token
        + byte array of remaining buffer
+ Tok - is a simple, non-look ahead tokenizer
    + parameter
        + a byte array representing the buffer to evaluate
    + returns
        + a Token of Type *Letter*, *Numeral*, *Punctuation* and *Space*
        + the remaining buffer byte array
+ Tok2 - is a function the take
    + parameters
        + a byte array representing the buffer to evaluate
        + A Tokenizer function
    + returns
        + a Token of Type defined by the Tokenizer function
        + the remaining buffer byte array
+ Words - Is an example Tokenizer function
    + returns tokens of type *Numeral*, *Punctuation*, *Space* and *Word*

