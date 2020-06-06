package client

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

)


// type Source struct {
// 	Id   string
// 	Name string
// }

// type Article struct {
// 	SourceName  Source
// 	Author      string
// 	Title       string
// 	Description string
// 	Url         string
// 	UrlToImage  string
// 	PublishedAt time.Time
// }

// type ArticlesResult struct {

// 	// buch os status
// 	// bunch of error
// 	TotalResults int
// 	Articles     []Article
// }

type TopHeadlinesRequest struct {
	Keyword  string
	Sources  []string
	Category Category
	Language Language
	Country  Country
	Page     int
	PageSize int
}

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *Client) GetTopHeadlines(req TopHeadlinesRequest) (*ArticlesResult, error) {

	rel := &url.URL{Path: "/topheadlines"}

	queryParams := []string{}

	if req.Keyword != "" {
		queryParams = append(queryParams, "q="+req.Keyword)
	}

	if len(req.Sources) > 0 {
		queryParams = append(queryParams, "sources="+strings.Join(req.Sources[:], ","))
	}

	// to add further query parameters

	queryString := strings.Join(queryParams[:], "&")

	rel.Path += queryString

	u := c.BaseURL.ResolveReference(rel)

	customReq, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}
	customReq.Header.Set("Accept", "application/json")
	customReq.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(customReq)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var articleResults ArticlesResult

	err = json.NewDecoder(resp.Body).Decode(&articleResults)

	return &articleResults, err

}
