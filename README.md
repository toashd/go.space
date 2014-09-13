go.space
===

go.space is a simple example and template for new go projects.

It allows multiple custom packages within the src directory, some space for docs and
third party packages that can be installed to the _vendor directory if you want follow the
strategy to put these under source control to better control their state and
stability and share the whole thing consistent with your team mates.

External packages used in the project can be freezed to the packages file and
easily installed with `make vendor_get`.

The whole app can be build with a simple `make`, that outputs the binaries to bin.



# Some thoughts about go packages and how to structure go apps in general

While there are different opinions and many ways to structure go code and apps, respectively, go.space is just one way and approach achieving this task and mainly intended for small and kind of internal go projects.

At the end of the day, structuring a go project heavily depends on the nature of "custom package and/or packages" used within the project.

If it's intended to be of general use, consider employing the so-called ["github code layout"](https://code.google.com/p/go-wiki/wiki/GithubCodeLayout) making your library a separate go get-table project.

As a starting point, be sure to read and understand the ["How to write Go code"](http://golang.org/doc/code.html) document.



# Contribute
Feel free to fork, contribute, share ideas and thoughts, or request a pull.

