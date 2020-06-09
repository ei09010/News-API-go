package main

import (
	"News-API-go/constants"
	"News-API-go/models"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var newsClient Client
var urlObj *url.URL

func TestMain(m *testing.M) {

	newsClient.ApiKey = "testKey"
	os.Exit(m.Run())
}

func TestGetTopHeadlines_StandardRequest_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 1,"articles": [{"source": {"id": "the-wall-street-journal","name": "The Wall Street Journal"},"author": "Alison Sider","title": "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal","description": "U.S. carriers are planning to operate smaller companies with fewer flights and employees","url": "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600","urlToImage": "https://images.wsj.net/im-194495/social","publishedAt": "2020-06-07T14:11:39Z","content": "Federal stimulus money for airlines is keeping them afloat through the coronavirus pandemic, but its not proving to be enough to sustain the industry at its pre-pandemic size.Carriers say they will… [+320 chars]"}]}`
	expectedCorrectRequest := `/top-headlines?category=business&language=en&country=ae`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		if req.URL.Path == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 1
	expectedFirstArticleSourceId := "the-wall-street-journal"
	expectedFirstArticleSourceName := "The Wall Street Journal"
	expectedFirstArticleAuthor := "Alison Sider"
	expectedFirstArticleTitle := "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal"
	expectedFirstArticleDescription := "U.S. carriers are planning to operate smaller companies with fewer flights and employees"
	expectedFirstArticleUrl := "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600"
	expectedFirstArticleUrlToImage := "https://images.wsj.net/im-194495/social"
	expectedPublishedAt := "2020-06-07 14:11:39 +0000 UTC"
	
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
		//Page:     1,
		//PageSize: 1,
		// Sources:  []string{"bla", "blo"},
	}

	newsClient.BaseURL, _ = url.Parse(ts.URL)

	// Act
	// execute client method and save result
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Articles[0].Url != expectedFirstArticleUrl {
		t.Errorf("handler returned unexpected firt article url: got %v want %v",
		response.Articles[0].Url, expectedFirstArticleUrl)
	}

	if response.Articles[0].UrlToImage != expectedFirstArticleUrlToImage {
		t.Errorf("handler returned unexpected firt article urlToImage: got %v want %v",
		response.Articles[0].UrlToImage, expectedFirstArticleUrlToImage)
	}

	if response.Articles[0].PublishedAt.String() != expectedPublishedAt {
		t.Errorf("handler returned unexpected firt article publishedAt: got %v want %v",
		response.Articles[0].PublishedAt, expectedPublishedAt)
	}

	if response.Articles[0].Content == "" {
		t.Errorf("handler returned empty content")
	}

	articleLength := len(response.Articles)

	if response.TotalResults != articleLength {
		t.Errorf("handler returned unexpected number of articles: got %v want %v",
		response.TotalResults, articleLength)
	}	
}


func TestGetTopHeadlines_StandardRequestWithPageInfo_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 1,"articles": [{"source": {"id": "the-wall-street-journal","name": "The Wall Street Journal"},"author": "Alison Sider","title": "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal","description": "U.S. carriers are planning to operate smaller companies with fewer flights and employees","url": "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600","urlToImage": "https://images.wsj.net/im-194495/social","publishedAt": "2020-06-07T14:11:39Z","content": "Federal stimulus money for airlines is keeping them afloat through the coronavirus pandemic, but its not proving to be enough to sustain the industry at its pre-pandemic size.Carriers say they will… [+320 chars]"}]}`
	expectedCorrectRequest := `/top-headlines?category=business&language=en&country=ae&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		if req.URL.Path == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 1
	expectedFirstArticleSourceId := "the-wall-street-journal"
	expectedFirstArticleSourceName := "The Wall Street Journal"
	expectedFirstArticleAuthor := "Alison Sider"
	expectedFirstArticleTitle := "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal"
	expectedFirstArticleDescription := "U.S. carriers are planning to operate smaller companies with fewer flights and employees"
	expectedFirstArticleUrl := "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600"
	expectedFirstArticleUrlToImage := "https://images.wsj.net/im-194495/social"
	expectedPublishedAt := "2020-06-07 14:11:39 +0000 UTC"
	
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
		Page:     1,
		PageSize: 1,
		// Sources:  []string{"bla", "blo"},
	}

	newsClient.BaseURL, _ = url.Parse(ts.URL)

	// Act
	// execute client method and save result
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Articles[0].Url != expectedFirstArticleUrl {
		t.Errorf("handler returned unexpected firt article url: got %v want %v",
		response.Articles[0].Url, expectedFirstArticleUrl)
	}

	if response.Articles[0].UrlToImage != expectedFirstArticleUrlToImage {
		t.Errorf("handler returned unexpected firt article urlToImage: got %v want %v",
		response.Articles[0].UrlToImage, expectedFirstArticleUrlToImage)
	}

	if response.Articles[0].PublishedAt.String() != expectedPublishedAt {
		t.Errorf("handler returned unexpected firt article publishedAt: got %v want %v",
		response.Articles[0].PublishedAt, expectedPublishedAt)
	}

	if response.Articles[0].Content == "" {
		t.Errorf("handler returned empty content")
	}

	articleLength := len(response.Articles)

	if response.TotalResults != articleLength {
		t.Errorf("handler returned unexpected number of articles: got %v want %v",
		response.TotalResults, articleLength)
	}	
}