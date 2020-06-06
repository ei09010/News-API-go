package models

import (
	"News-API-go/constants"
	"time"
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
	Status       int
	Error        Error
	TotalResults int
	Articles     []Article
}

type TopHeadlinesRequest struct {
	Keyword  string
	Sources  []string
	Category constants.Category
	Language constants.Language
	Country  constants.Country
	Page     int
	PageSize int
}

type Error struct {
	Error   constants.ErrorCode
	Message string
}
