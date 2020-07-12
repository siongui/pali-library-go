package lib

import (
	"strings"
	"unicode/utf8"
)

// Type of the webpage, determined according to path of URL
//go:generate stringer -type=PageType
type PageType int

const (
	RootPage PageType = iota
	AboutPage
	PrefixPage
	WordPage
	NoSuchPage
)

// Do not handle utf8 string verion of DeterminePageType
func DeterminePageTypeNoUtf(urlpath string) PageType {
	if urlpath == "/" {
		return RootPage
	}
	if urlpath == "/about/" {
		return AboutPage
	}
	if IsValidPrefixUrlPathNoUtf(urlpath) {
		return PrefixPage
	}
	if IsValidWordUrlPath(urlpath) {
		return WordPage
	}

	return NoSuchPage
}

// Do not handle utf8 string verion of IsValidPrefixUrlPath
func IsValidPrefixUrlPathNoUtf(urlpath string) bool {
	ss := strings.Split(urlpath, "/")

	if len(ss) != 4 {
		return false
	}

	if ss[0] != "" {
		return false
	}

	if ss[1] != "browse" {
		return false
	}

	if ss[3] != "" {
		return false
	}

	return true
}

// Do not handle utf8 string verion of GetPrefixFromUrlPath
func GetPrefixFromUrlPathNoUtf(urlpath string) string {
	if IsValidPrefixUrlPathNoUtf(urlpath) {
		ss := strings.Split(urlpath, "/")
		return ss[2]
	}

	return ""
}

// DeterminePageType determines the type of the webpage according to path of
// URL.
func DeterminePageType(urlpath string) PageType {
	if urlpath == "/" {
		return RootPage
	}
	if urlpath == "/about/" {
		return AboutPage
	}
	if IsValidPrefixUrlPath(urlpath) {
		return PrefixPage
	}
	if IsValidWordUrlPath(urlpath) {
		return WordPage
	}

	return NoSuchPage
}

// IsValidPrefixUrlPath will return true if the path of the url is a possible
// prefix of Pāli words.
func IsValidPrefixUrlPath(urlpath string) bool {
	ss := strings.Split(urlpath, "/")

	if len(ss) != 4 {
		return false
	}

	if ss[0] != "" {
		return false
	}

	if ss[1] != "browse" {
		return false
	}

	if ss[3] != "" {
		return false
	}

	if ss[2] != GetFirstCharacterOfWord(ss[2]) {
		return false
	}

	return true
}

// IsValidWordUrlPath will return true if the path of the url is a possible Pāli
// word.
func IsValidWordUrlPath(urlpath string) bool {
	ss := strings.Split(urlpath, "/")

	if len(ss) != 5 {
		return false
	}

	if ss[0] != "" {
		return false
	}

	if ss[1] != "browse" {
		return false
	}

	if ss[4] != "" {
		return false
	}

	if !strings.HasPrefix(ss[3], ss[2]) {
		return false
	}

	return true
}

// GetPrefixFromUrlPath will return the prefix string embedded in the path of
// url if url path is valid. Otherwise return empty string. Note that this
// method do not check if the prefix string is a valid prefix. Use with caution.
//
// For example,
//
// "/browse/s/" will return "s"
//
// "/browse/āā/" will return ""
func GetPrefixFromUrlPath(urlpath string) string {
	if IsValidPrefixUrlPath(urlpath) {
		ss := strings.Split(urlpath, "/")
		return ss[2]
	}

	return ""
}

// GetWordFromUrlPath will return the word string embedded in the path of url if
// url path is valid. Otherwise return empty string. Note that this method do
// not check if the word string is a valid word. Use with caution.
//
// For example,
//
// "/browse/s/sacca/" will return "sacca"
//
// "/browse/s/āpadā/" will return ""
func GetWordFromUrlPath(urlpath string) string {
	if IsValidWordUrlPath(urlpath) {
		ss := strings.Split(urlpath, "/")
		return ss[3]
	}

	return ""
}

// WordUrlPath will return the url path of the given Pāli word.
//
// Example:
//
// URL path of word ``sacca`` is:
//
//   /browse/s/sacca/
//
// URL path of word ``āpadā`` is:
//
//   /browse/ā/āpadā/
//
// Note that this method do not check the validity of the word. Use with
// caution.
func WordUrlPath(word string) string {
	return "/browse/" + GetFirstCharacterOfWord(word) + "/" + word + "/"
}

// GetFirstCharacterOfWord returns first character of the word. For example,
// āpadā will return ā
//
// FIXME: this method will return incorrect output if compiled to JavaScript
// via GopherJS. ḍ (%E1%B8%8D) will return % in JavaScript environment.
//
// Google search: gopherjs utf8
//
//   Unable to correctly handle certain unicode/utf-8 characters
//   https://github.com/gopherjs/gopherjs/issues/319
//
//   GopherJS utf8 encoding problem
//   https://gist.github.com/cryptix/054b955e55f144428f97/0ef91a71e286cc7f7334ac1c99ec78dc629db784
func GetFirstCharacterOfWord(word string) string {
	runeValue, _ := utf8.DecodeRuneInString(word)
	return string(runeValue)
}

// PrefixUrlPath will return the url path of the given prefix.
//
// Example:
//
// URL path of prefix ``s`` is:
//
//   /browse/s/
//
// URL path of prefix ``ā`` is:
//
//   /browse/ā/
//
// Note that this method do not check the validity of the prefix. Use with
// caution.
func PrefixUrlPath(prefix string) string {
	return "/browse/" + prefix + "/"
}
