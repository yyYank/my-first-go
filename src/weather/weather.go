package weather

import (
	"net/http"
)


/*
 */

func main() {
	url := "http://yyyank.blogspot.com/feeds/posts/default?max-results=999"
	response, _ := http.Get(url)
	defer response.Body.Close()
}
