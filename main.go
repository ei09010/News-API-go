package main

import (
	"News-API-go/models"
	"News-API-go/constants"
	"News-API-go/client"
	"fmt"
)

var newsClient = client.NewClient("https://newsapi.org", "XXXXXXXXXXXXXXXXXXXXXXXXXXXX")

func main(){

	topHeadlineRequest := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
	}

	response, err := newsClient.GetTopHeadlines(topHeadlineRequest)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("we had the following result amount: ", response.TotalResults)

	for _, article := range response.Articles{
	
		fmt.Println(article.SourceName)

		fmt.Println(article.Author)

		fmt.Println(article.Title)

		fmt.Println(article.Description)

		fmt.Println(article.Url)

		fmt.Println(article.UrlToImage)

		fmt.Println(article.PublishedAt)

		fmt.Println(`
		
		*******************
		
		`)
	}
		
}