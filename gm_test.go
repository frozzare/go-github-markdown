package gm

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestFetch(t *testing.T) {
	s := strings.Contains(Fetch("# Hello"), "Hello</h1>")
	assert.Equal(t, true, s)
}

func TestWrite(t *testing.T) {
	file := "/tmp/gm_test"
	f, err := os.Create(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	Write(f, "# Hello")

	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	s := strings.Contains(string(data), "Hello</h1>")
	assert.Equal(t, true, s)
}
