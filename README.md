go.space
===

go.space is a simple example and template for new go projects.

It allows multiple custom packages within the `src` directory, some space for
docs and third party packages that can be installed to the `_vendor` directory in case
you want to employ the strategy to put external packages under source control to
better control their state and share them consistently with your team mates.

External packages used in the project can be freezed to the packages file and
easily installed with `make vendor_get`.

The whole app can be build with a simple `make`, that outputs the binaries to
`bin`.


#### Some thoughts about go packages and how to structure go apps in general

While there are different opinions and many ways to structure go code and apps,
respectively, go.space is first and foremost one way and possible answer
achieving this task and intended for small and kind of internal go projects.

At the end of the day, structuring your go project heavily depends on the nature
of your "custom package and/or packages" used within the project.

If it's intended to be of general use, consider employing the so-called ["github
code layout"](https://code.google.com/p/go-wiki/wiki/GithubCodeLayout) making
your library a separate go get-table project.

As a starting point, be sure to read and understand the ["How to write Go
code"](http://golang.org/doc/code.html) document.


#### Contribute
Feel free to fork, contribute, share ideas and thoughts, or request a pull.

