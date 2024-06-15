Shrivel
=======

Shrivel is both a minification (but not obfuscation) tool and Go library. It
minifies files by removing whitespace in a context-free manner. This means
that there is no language specific parsing, shrivel will simply remove any
unicode whitespace character as it comes across it, unless the space is
significant (such as whitespace within strings).

Currently, shrivel supports SQL.

In order to install shrivel, it is easy to build from source like so:

    git clone https://github.com/jordanocokoljic/shrivel
    cd shrivel
    go build ./cmd/shrivel

Then copy the compiled executable somewhere it can be found on PATH.
