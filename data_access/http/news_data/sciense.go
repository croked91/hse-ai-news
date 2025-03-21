package news_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/croked91/news-ai/domain"
)

const scienceNewsLink = "https://api.worldnewsapi.com/search-news?categories=science,sports&number=5&language=ru&api-key="

type RawNews struct {
	Digest domain.NewsList `json:"news"`
}

func (c *Controller) ScienceNewsS() {
	apiKey := c.apiKey
	ep := scienceNewsLink + apiKey

	fmt.Println(ep)

	resp, err := http.Get(ep)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var rawNews RawNews
	err = json.Unmarshal(body, &rawNews)
	if err != nil {
		fmt.Println(err)
	}

	c.llm.ProcessNews(rawNews.Digest)
}
