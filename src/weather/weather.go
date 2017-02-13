package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

/**
 * goqueryでyahoo天気からそれとなくデータを取る
 */
func main() {
	doc, err := goquery.NewDocument("http://weather.yahoo.co.jp/weather/13/4410.html")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		alt, _ := s.Attr("alt")
		fmt.Println(alt)
	})
}
