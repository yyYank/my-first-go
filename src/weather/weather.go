package weather

import (
	"net/http"
)


/*
このあたりを参考にした
 http://qiita.com/taizo/items/c397dbfed7215969b0a5
 http://cuto.unirita.co.jp/gostudy/post/standard-library-xml/
 */

func main() {
	url := "http://yyyank.blogspot.com/feeds/posts/default?max-results=999"
	response, _ := http.Get(url)
	defer response.Body.Close()
}
