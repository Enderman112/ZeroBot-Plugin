package nbnhhsh

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	zero.OnRegex(`^[?？]{1,2} ?([a-z0-9]+)$`).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			keyword := ctx.State["regex_matched"].([]string)[1]
			ctx.SendChain(message.Text(keyword + ": " + strings.Join(getValue(keyword), ", ")))
		})
}

func getValue(text string) []string {
	urlValues := url.Values{}
	urlValues.Add("text", text)
	resp, _ := http.PostForm("https://lab.magiconch.com/api/nbnhhsh/guess", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	json := gjson.ParseBytes(body)
	res := make([]string, 0)
	var jsonPath string
	if json.Get("0.trans").Exists() {
		jsonPath = "0.trans"
	} else {
		jsonPath = "0.inputting"
	}
	for _, value := range json.Get(jsonPath).Array() {
		res = append(res, value.String())
	}
	return res
}
