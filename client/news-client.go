package main

import (
	"News-API-go/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"strconv"
)

// Client to aggregate basic information regarding http calls
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	ApiKey     string
	httpClient *http.Client
}

func main() {

}

// GetTopHeadlines to request top headlines
func (c *Client) GetTopHeadlines(req models.TopHeadlinesRequest) (*models.ArticlesResult, error) {

	relativePath := &url.URL{Path: "/top-headlines"}

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

	//page information

	if req.Page > 0 {
		queryParams = append(queryParams, "page="+strconv.Itoa(req.Page))
	}

	if req.PageSize > 0 {
		queryParams = append(queryParams, "pageSize="+strconv.Itoa(req.PageSize))
	}

	queryString := strings.Join(queryParams[:], "&")

	if len(queryParams) > 0{
		relativePath.Path += "?"
	}

	relativePath.Path += queryString

	urlAbsoluteReference := c.BaseURL.ResolveReference(relativePath)

	return c.makeRequest(urlAbsoluteReference, queryString)

}

func (c *Client) makeRequest(url *url.URL, queryString string) (*models.ArticlesResult, error) {

	customReq, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, err
	}

	customReq.Header.Set("accept", "application/json")
	customReq.Header.Set("user-agent", c.UserAgent)
	customReq.Header.Set("x-api-key", c.ApiKey)

	c.httpClient = &http.Client{}

	httpResponse, err := c.httpClient.Do(customReq)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	var articleResults models.ArticlesResult

	if httpResponse != nil {

		myResponse, err := ioutil.ReadAll(httpResponse.Body)

		if err != nil {
			return nil, err
		}

		httpResponse.Body.Close()

		err = json.Unmarshal(myResponse, &articleResults)

		if articleResults.Status != string(http.StatusOK) {

			err = json.Unmarshal(myResponse, &articleResults.Error)

			return &articleResults, err
		}

	} else {
			return &articleResults, err
	}

	return &articleResults, err
}
