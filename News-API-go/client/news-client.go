package client

import (
	"net/http"
	"net/url"
	"time"
)

type Category int

const (
	Business Category = iota
	Entertainment
	Health
	Science
	Sports
	Technology
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type Source struct {
	Id   string
	Name string
}

type Article struct {
	SourceName  Source
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt time.Time
}

type ArticlesResult struct {

	// buch os status
	// bunch of error
	TotalResults int
	Articles     []Article
}

type TopHeadlinesRequest struct {
	Keyword string
	Sources []string
}

// func (c *Client) GetTopHeadlines() (*ArticlesResult, error) {

// }
