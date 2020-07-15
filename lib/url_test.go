package lib

import (
	"testing"
)

func TestSetSiteUrl(t *testing.T) {
	SetSiteUrl("https://dictionary.sutta.org/")
	if rootPath != "/" {
		t.Error(rootPath)
	}

	SetSiteUrl("https://dictionary.sutta.org")
	if rootPath != "" {
		t.Error(rootPath)
	}
}

func TestSetCurrentLocale(t *testing.T) {
	SetCurrentLocale("zh_TW")
	if currentLocale != "zh_TW" {
		t.Error(currentLocale)
	}
}

func TestDeterminePageType(t *testing.T) {
	if DeterminePageType("/") != RootPage {
		t.Error("error root page type")
	}

	if DeterminePageType("/about/") != AboutPage {
		t.Error("error about page type")
	}

	if DeterminePageType("/browse/s/sacca/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/ā/āpadā/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/s/āpadā/") == WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/s/sacca") != NoSuchPage {
		t.Error("error no such page")
	}

	if DeterminePageType("/about/something") != NoSuchPage {
		t.Error("error no such page")
	}

	if DeterminePageType("/browse/s/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/āa/") == PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/ā/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8D/") != PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8/") == PrefixPage {
		t.Error("error prefix page type")
	}

	if DeterminePageType("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != WordPage {
		t.Error("error word page type")
	}

	if DeterminePageType("/browse/%E1%B8%8/%E1%B8%8Dibhi/") == WordPage {
		t.Error("error word page type")
	}
}

func TestIsValidPrefixUrlPath(t *testing.T) {
	if !IsValidPrefixUrlPath("/browse/ā/") {
		t.Error("/browse/ā/ should be true")
	}

	if IsValidPrefixUrlPath("/browse/āa/") {
		t.Error("/browse/āa/ should be false")
	}

	if IsValidPrefixUrlPath("/about/") {
		t.Error("/about/ should be false")
	}

	if !IsValidPrefixUrlPath("/browse/%E1%B8%8D/") {
		println("/browse/%E1%B8%8D/ (%E1%B8%8D is ḍ) should be true")
		t.Error("/browse/ḍ/ should be true")
	}

	if IsValidPrefixUrlPath("/browse/%E1%B8%8/") {
		println("/browse/%E1%B8%8/ should be false")
		t.Error("should be true")
	}
}

func TestIsValidWordUrlPath(t *testing.T) {
	if !IsValidWordUrlPath("/browse/ā/āpadā/") {
		t.Error("/browse/ā/āpadā/ should be true")
	}

	if IsValidWordUrlPath("/browse/ā/āpadā") {
		t.Error("/browse/ā/āpadā should be false")
	}

	if IsValidWordUrlPath("/about/") {
		t.Error("/about/ should be false")
	}

	if !IsValidWordUrlPath("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") {
		t.Error("/browse/ḍ/ḍibhi/ should be true")
	}

	if IsValidWordUrlPath("/browse/%E1%B8%8/%E1%B8%8Dibhi/") {
		t.Error("should be false")
	}
}

func TestGetPrefixFromUrlPath(t *testing.T) {
	if GetPrefixFromUrlPath("/browse/s/") != "s" {
		t.Error("/browse/s/ should return s")
	}

	if GetPrefixFromUrlPath("/browse/āā/") != "" {
		t.Error(`/browse/āā/ should return ""`)
	}

	if GetPrefixFromUrlPath("/browse/ā/") != "ā" {
		t.Error(`/browse/ā/ should return "ā"`)
	}

	if GetPrefixFromUrlPath("/about/") != "" {
		t.Error(`/about/ should return ""`)
	}

	if GetPrefixFromUrlPath("/browse/%E1%B8%8D/") != "ḍ" {
		println("/browse/%E1%B8%8D/ should return ḍ")
		t.Error("should return ḍ")
	}

	if GetPrefixFromUrlPath("/browse/%E1%B8%8/") != "" {
		println(`/browse/%E1%B8%8/ should return ""`)
		t.Error(`should return ""`)
	}
}

func TestGetWordFromUrlPath(t *testing.T) {
	if GetWordFromUrlPath("/browse/s/sacca/") != "sacca" {
		t.Error("/browse/s/sacca/ should return sacca")
	}

	if GetWordFromUrlPath("/browse/s/āpadā/") != "" {
		t.Error(`/browse/s/āpadā/ should return ""`)
	}

	if GetWordFromUrlPath("/about/") != "" {
		t.Error(`/about/ should return ""`)
	}

	if GetWordFromUrlPath("/browse/%E1%B8%8D/%E1%B8%8Dibhi/") != "ḍibhi" {
		println("/browse/%E1%B8%8D/%E1%B8%8Dibhi/ should return ḍibhi")
		t.Error("should return ḍibhi")
	}

	if GetWordFromUrlPath("/browse/%E1%B8%8/%E1%B8%8Dibhi/") != "" {
		println(`/browse/%E1%B8%8/%E1%B8%8Dibhi/ should ""`)
		t.Error(`should return ""`)
	}
}

func TestWordUrlPath(t *testing.T) {
	if WordUrlPath("sacca") != "/browse/s/sacca/" {
		t.Error("error sacca path url")
	}

	if WordUrlPath("āpadā") != "/browse/ā/āpadā/" {
		t.Error("error āpadā path url")
	}
}

func TestGetFirstCharacterOfWord(t *testing.T) {
	if GetFirstCharacterOfWord("sacca") != "s" {
		t.Error("sacca first char wrong")
	}

	if GetFirstCharacterOfWord("āpadā") != "ā" {
		t.Error("āpadā first char wrong")
	}

	if GetFirstCharacterOfWord("ṭakāra") != "ṭ" {
		t.Error("ṭakāra first char wrong")
	}

	if GetFirstCharacterOfWord("ṭ") != "ṭ" {
		t.Error("ṭ first char wrong")
	}

	if GetFirstCharacterOfWord("ḍ") != "ḍ" {
		t.Error("ḍ first char wrong")
	}
}

func TestPrefixUrlPath(t *testing.T) {
	if PrefixUrlPath("s") != "/browse/s/" {
		t.Error("prefix s url path wrong")
	}

	if PrefixUrlPath("ā") != "/browse/ā/" {
		t.Error("prefix ā url path wrong")
	}
}
