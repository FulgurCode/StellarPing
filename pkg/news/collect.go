package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CollectNews(day string) {
	var NEWS_API = os.Getenv("NEWS_API")
	var url = "https://spacenews.p.rapidapi.com/datenews/" + day

	var req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("x-rapidapi-key", NEWS_API)
	req.Header.Add("x-rapidapi-host", "spacenews.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response []News
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	var news []interface{}
	for _, n := range response {
		news = append(news, n)
	}

	AddNews(news)
}
