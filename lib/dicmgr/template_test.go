package dicmgr

import (
	"net/http"
	"strings"
	"testing"

	"github.com/siongui/gopalilib/lib"
)

func HttpWordJsonPath(word string) string {
	return "https://siongui.github.io/xemaauj9k5qn34x88m4h/" + word + ".json"
}

func TestGetWordDefinitionHtml(t *testing.T) {
	resp, err := http.Get(HttpWordJsonPath("sacca"))
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error(`resp.StatusCode != 200`)
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	setting := lib.GetDefaultPaliSetting()

	html := GetWordDefinitionHtml(wi, setting, `en-US,en,zh-TW`)
	if !strings.Contains(html, "<p>巴利文-漢文佛學名相辭匯 翻譯：張文明</p>") {
		t.Error(html)
		return
	}

	setting.P2zh = false
	html = GetWordDefinitionHtml(wi, setting, `en-US,en,zh-TW`)
	if strings.Contains(html, "<p>巴利文-漢文佛學名相辭匯 翻譯：張文明</p>") {
		t.Error(html)
		return
	}
	//t.Log(html)
}

func TestGetSuggestedWordsHtml(t *testing.T) {
	html := GetSuggestedWordsHtml(GetSuggestedWords("sacc", 10))
	if !strings.Contains(html, "<div>sacca</div>") {
		t.Error(html)
		return
	}
	//t.Log(html)
}

func TestGetWordPreviewHtml(t *testing.T) {
	word := "sacca"
	resp, err := http.Get(HttpWordJsonPath(word))
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error(`resp.StatusCode != 200`)
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	setting := lib.GetDefaultPaliSetting()

	html := GetWordPreviewHtml(word, wi, setting, `en-US,en,zh-TW`)
	if !strings.Contains(html, "<span>巴利文-漢文佛學名相辭匯 翻譯：張文明</span>") {
		t.Error(html)
		return
	}

	setting.P2zh = false
	html = GetWordDefinitionHtml(wi, setting, `en-US,en,zh-TW`)
	if strings.Contains(html, "<span>巴利文-漢文佛學名相辭匯 翻譯：張文明</span>") {
		t.Error(html)
		return
	}
	//t.Log(html)
}
