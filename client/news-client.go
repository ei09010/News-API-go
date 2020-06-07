package main

import (
	"News-API-go/constants"
	"News-API-go/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Client to aggregate basic information regarding http calls
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	ApiKey     string
	httpClient *http.Client
}

// GetTopHeadlines to request top headlines
func (c *Client) GetTopHeadlines(req models.TopHeadlinesRequest) (*models.ArticlesResult, error) {

	relativePath := &url.URL{Path: "/top-headlines?"}

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

	if req.Page > 0 {
		queryParams = append(queryParams, "page="+string(req.Page))
	}

	if req.PageSize > 0 {
		queryParams = append(queryParams, "pageSize="+string(req.PageSize))
	}

	queryString := strings.Join(queryParams[:], "&")

	relativePath.Path += queryString

	fmt.Println("relative path: ", relativePath.Path)

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

	httpResponse, err := c.httpClient.Do(customReq)

	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	var articleResults models.ArticlesResult

	if httpResponse != nil {
		err = json.NewDecoder(httpResponse.Body).Decode(&articleResults)

		if articleResults.Status != http.StatusOK {
			err = errors.New("The API returned an error code that wasn't expected: " + string(articleResults.Status))

			articleResults.Error = models.Error{
				Error:   constants.UnexpectedError,
				Message: "The API returned an error code that wasn't expected: " + string(articleResults.Status)}
		}

	} else {
		articleResults.Status = http.StatusInternalServerError

		articleResults.Error = models.Error{
			Error:   constants.UnexpectedError,
			Message: "The API returned an empty response. Are you connected to the internet?"}
	}

	return &articleResults, err
}
