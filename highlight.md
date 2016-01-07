# Markdown with inline syntax highlighting

There are great libraries for markdown rendering in golang. I have had a chance
to use  some of them namely
[blackfriday](https://github.com/russross/blackfriday) and
[mark](https://github.com/a8m/mark). 

Today I would like to share with you my latest adventure. Now you can render
markdown files with syntax highlighting that is inline.

## What the hell is inline styles?

If you want to add styles to html files there are two ways.

* Link with external files in the header

* Inline the styles with the `style` attribute.

So whatever the trick the syntax highlighter is doing it has those two options
to bring color to your rendered html documents.

Majority uses the [code-pretttify](https://github.com/google/code-prettify)
classes to the document and requires you to
link to the [code-pretttify
stylesheet](https://github.com/google/code-prettify#setup) to have the document
highlighted.

## Why inline styles for your code snippets?

* Your page is already clouded with many links to external stylesheets and
you dont feel like it is worthy to add another dependency 
to get the that short code snippet highlighted.

* For the sake of knowledge( why not?).


After saying that, I would like to introduce to you
[frontman](https://github.com/gernest/frontman) my attempt to 
collect useful  stuffs for static web( In golang of course).For the past few
weeks I have been taking apart some of the golang projects and messing with them
by adding fun features. I have successfully put together a markdown renderer
which support inline syntax highlighting.

Install the package as follows

```bash
  go get github.com/gernest/frontman
 ``` 

# Example

You can render this document into html with zero dependency and
beautiful syntax highlighting, [the example code is
included in this respository](code/highlight.go).

You can also see the results in the themes section below after applying
different themes that are shipped with the
[frontman](https://github.com/gernest/frontman)   package.

The full example code

```go
package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"

	"github.com/gernest/frontman"
)

func main() {
	var (
		fileName   string
		themeName  string
		outputFile string
		inline     bool
		highlight  bool
	)

	flag.StringVar(&fileName, "f", "highlight.md", "Input markdown document")
	flag.StringVar(&themeName, "t", "prettify", "Theme name")
	flag.StringVar(&outputFile, "o", "index.html", "File to save outuput")
	flag.BoolVar(&inline, "i", true, "enable inline style")
	flag.BoolVar(&highlight, "h", true, "enable syntax highlight")
	flag.Parse()

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	out := &bytes.Buffer{}
	err = frontman.Markdown(out, bytes.NewReader(data), highlight,
  inline,themeName)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(outputFile, out.Bytes(), 0600)
}
```

then you can do the following to render this document

```shell
$ git clone git@github.com:gernest/talk.git

$ cd talk

$ go run code/highlight.go  -f highlight.md

$ sensible-browser index.html
```

In case you want to play with the example code. You can see all the available
oprions via

```shell
$ go run code/highlight.go -help
```

# Themes
This package comes shipped with a couple of theme that you can choose from( not
they are taken from the [code-pretttify
themes](https://rawgit.com/google/code-prettify/master/styles/index.html)


__Themes Demo( of this document rendered with frontman)__

* [Deault](http://gernest.github.io/frontman/highlight/prettify.html)
* [Desert](http://gernest.github.io/frontman/highlight/desert.html)
* [doxy](http://gernest.github.io/frontman/highlight/doxy.html)
* [obsidian](http://gernest.github.io/frontman/highlight/obsidian.html)
* [sunburst](http://gernest.github.io/frontman/highlight/sunburst.html)

# Summary
Markdown is a very good standard and golang has fantastic tools for it.

I plan to add more interesting utilities into  
[frontman](https://github.com/gernest/frontman) , and you are very
welcome if you have something in mind to help people build a better static web
experience.

__PS__: This repository is my blog and   my twitter account
[@gernesti](https://twitter.com/gernesti) is my
RSS feed.

