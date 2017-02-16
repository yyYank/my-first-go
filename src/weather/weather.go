package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

/**
 * goqueryでyahoo天気からそれとなくデータを取る
 */
func main() {
	doc, err := goquery.NewDocument("http://weather.goo.ne.jp/")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find(".overview").Each(func(_ int, s *goquery.Selection) {
		var res = s.Find("p").Text()
		fmt.Println(res)
	})
}
