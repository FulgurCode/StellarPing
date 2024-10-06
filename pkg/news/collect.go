package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Data []News `json:"data"`
}

func CollectNews() {
	var NEWS_API = os.Getenv("NEWS_API")
	var url = fmt.Sprintf("http://api.mediastack.com/v1/news?access_key=%s&keywords=space&source=en&categories=science&language=en", NEWS_API)

	var req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response Response
	json.Unmarshal(body, &response)

	var news []interface{}
	for _, n := range response.Data {
		news = append(news, n)
	}

	AddNews(news)
}
