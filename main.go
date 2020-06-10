package main

import (
	"News-API-go/models"
	"News-API-go/constants"
	"News-API-go/client"
	"fmt"
)

var newsClient = client.NewClient("https://newsapi.org", "7ba67bcc987240369c1915f6e27b4b3f")

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
		
}