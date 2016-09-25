package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	url := "http://yyyank.blogspot.com/feeds/posts/default?max-results=1"
	response, _ := http.Get(url)
	defer response.Body.Close()
	byteArray, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}
