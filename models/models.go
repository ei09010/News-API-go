package models

import (
	"News-API-go/constants"
	"time"
)

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	SourceName  Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content string `json:"content"`
}


// type AutoGenerated struct {
// 	Status       string `json:"status"`
// 	TotalResults int    `json:"totalResults"`
// 	Articles     []struct {
// 		Source struct {
// 			ID   string `json:"id"`
// 			Name string `json:"name"`
// 		} `json:"source"`
// 		Author      string    `json:"author"`
// 		Title       string    `json:"title"`
// 		Description string    `json:"description"`
// 		URL         string    `json:"url"`
// 		URLToImage  string    `json:"urlToImage"`
// 		PublishedAt time.Time `json:"publishedAt"`
// 		Content     string    `json:"content"`
// 	} `json:"articles"`
// }

type ArticlesResult struct {
	Status       string `json:"status"`
	//Error        Error
	TotalResults int  `json:"totalResults"`
	Articles     []Article `json:"articles"`
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
