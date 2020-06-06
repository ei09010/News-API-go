package models

import (
	"time"
	"News-API-go/News-API-go/constants"
)

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
	Keyword  string
	Sources  []string
	Category Category
	Language Language
	Country  Country
	Page     int
	PageSize int
}
