package client

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *Client) GetTopHeadlines(req TopHeadlinesRequest) (ArticlesResult, error) {

	rel := &url.URL{Path: "/topheadlines"}

	queryParams := []string{}

	if req.KeyWord != "" {
		append(queryParams, "q="+req.KeyWord)
	}

	if len(req.Sources) > 0 {
		append(queryParams, "sources="+strings.Join(req.Sources[:], ","))
	}

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
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err

}
