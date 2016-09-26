package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"strconv"
)

func main() {
	url := "http://yyyank.blogspot.com/feeds/posts/default?max-results=999"
	response, _ := http.Get(url)
	defer response.Body.Close()
	byteArr, _ := ioutil.ReadAll(response.Body)
	v := Result{}
	err := xml.Unmarshal(byteArr, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for i := range v.Entries {
		fmt.Printf("Title: %q\n", v.Entries[i])
	}
	fmt.Println("ブログの投稿数=====>>>" + strconv.Itoa(len(v.Entries)))
}

type Result struct {
	Entries []Entry `xml:"entry"`
}
type Entry struct {
	Titles []Title `xml:"title"`
}
type Title struct {
	Title string `xml:",chardata"`
}