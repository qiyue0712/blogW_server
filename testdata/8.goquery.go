package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	// 创建文档
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	selection := doc.Find("title")
	doc.Find("head").AppendHtml("<meta name=\"keyword1\" content=\"枫枫知道,网站,开发,程序员,golang\">")
	fmt.Println(selection.Text())
	selection.SetText("王佳阅0712")
	selection.SetAttr("", "")
	html, err := doc.Html()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(html)

	// 寻找标签

	// 修改标签
}
