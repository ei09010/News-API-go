package client

import (
	"News-API-go/models"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Client to aggregate basic information regarding http calls
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

// GetTopHeadlines to request top headlines
func (c *Client) GetTopHeadlines(req models.TopHeadlinesRequest) (*models.ArticlesResult, error) {

	rel := &url.URL{Path: "/topheadlines"}

	queryParams := []string{}

	if req.Keyword != "" {
		queryParams = append(queryParams, "q="+req.Keyword)
	}

	if len(req.Sources) > 0 {
		queryParams = append(queryParams, "sources="+strings.Join(req.Sources[:], ","))
	}

	if string(req.Category) != "" {
		queryParams = append(queryParams, "category="+strings.ToLower(string(req.Category)))
	}

	if string(req.Language) != "" {
		queryParams = append(queryParams, "language="+strings.ToLower(string(req.Language)))
	}

	if string(req.Country) != "" {
		queryParams = append(queryParams, "country="+strings.ToLower(string(req.Country)))
	}

	//page

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

	var articleResults models.ArticlesResult

	err = json.NewDecoder(resp.Body).Decode(&articleResults)

	return &articleResults, err

}
