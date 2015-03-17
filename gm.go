package gm

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	url = "https://api.github.com/markdown/raw"
)

// Fetch the generated Markdown from GitHub API.
func Fetch(markdown string) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(markdown)))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("User-Agent", "https://github.com/frozzare/go-github-markdown")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

// Write the generated Markdown from GitHub API to a writer.
func Write(w io.Writer, markdown string) {
	html := Fetch(markdown)
	w.Write([]byte(html))
}
