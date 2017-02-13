package weather

import (
	"net/http"
)


/*
 * 天気情報の取得
 */
func main() {
	url := "http://yyyank.blogspot.com/feeds/posts/default?max-results=999"
	response, _ := http.Get(url)
	defer response.Body.Close()
}
