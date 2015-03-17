# go-github-markdown [![Build Status](https://travis-ci.org/frozzare/go-github-markdown.svg?branch=master)](https://travis-ci.org/frozzare/go-github-markdown)

 Fetch rendered markdown using the GitHub Markdown API

 View the [docs](http://godoc.org/github.com/frozzare/go-github-markdown).

## Installation

```
$ go get github.com/frozzare/go-github-markdown
```

## Example

```go
package main

import (
    "os"

	"github.com/frozzare/go-github-markdown"
)

func main() {
    // Fetch generated GitHub Markdown to string
	gm.Fetch("# Hello")
	//=> <h1><a id="user-content-hello" class="anchor" href="#hello" aria-hidden="true"><span class="octicon octicon-link"></span></a>Hello</h1>

    // Write using io.Writer interface
    // Example, write to file.
    f, _ := os.Create("/tmp/gm_example")
    defer f.Close()

	gm.Write(f, "# Hello")

    // Body of /tmp/gm_example:
    // => <h1><a id="user-content-hello" class="anchor" href="#hello" aria-hidden="true"><span class="octicon octicon-link"></span></a>Hello</h1>
}
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
