package main

import (
	"testing"
)

var errFormat = `got "%s" want "%s"`

func TestParseError(t *testing.T) {
	_, err := UrlMaker("some_bad_string")
	expected := "this is not URL for git"
	if err.Error() != expected {
		t.Errorf(errFormat, err.Error(), expected)
	}
}

func TestParseGit(t *testing.T) {
	var expected string
	m, _ := UrlMaker("git://github.com/hoge/fuga.git")

	expected = "git"
	if m.Scheme != expected {
		t.Errorf(errFormat, m.Scheme, expected)
	}

	expected = ""
	if m.Username != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "github.com"
	if m.Host != expected {
		t.Errorf(errFormat, m.Host, expected)
	}

	expected = "hoge/fuga"
	if m.Path != expected {
		t.Errorf(errFormat, m.Path, expected)
	}
}

func TestParseSsh(t *testing.T) {
	var expected string
	m, _ := UrlMaker("ssh://git@github.com/hoge/fuga.git")

	expected = "ssh"
	if m.Scheme != expected {
		t.Errorf(errFormat, m.Scheme, expected)
	}

	expected = "git"
	if m.Username != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "github.com"
	if m.Host != expected {
		t.Errorf(errFormat, m.Host, expected)
	}

	expected = "hoge/fuga"
	if m.Path != expected {
		t.Errorf(errFormat, m.Path, expected)
	}
}

func TestParseSimpleSsh(t *testing.T) {
	var expected string
	m, _ := UrlMaker("git@github.com:/hoge/fuga.git")

	expected = ""
	if m.Scheme != expected {
		t.Errorf(errFormat, m.Scheme, expected)
	}

	expected = "git"
	if m.Username != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "github.com"
	if m.Host != expected {
		t.Errorf(errFormat, m.Host, expected)
	}

	expected = "/hoge/fuga"
	if m.Path != expected {
		t.Errorf(errFormat, m.Path, expected)
	}
}

func TestParseHttps(t *testing.T) {
	var expected string
	m, _ := UrlMaker("https://github.com/hoge/fuga.git")

	expected = "https"
	if m.Scheme != expected {
		t.Errorf(errFormat, m.Scheme, expected)
	}

	expected = ""
	if m.Username != expected {
		t.Errorf(errFormat, m.Username, expected)
	}

	expected = "github.com"
	if m.Host != expected {
		t.Errorf(errFormat, m.Host, expected)
	}

	expected = "hoge/fuga"
	if m.Path != expected {
		t.Errorf(errFormat, m.Path, expected)
	}
}
