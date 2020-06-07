package main

import (
	"News-API-go/constants"
	"News-API-go/models"
	"encoding/json"
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
	expected := `{"status": "ok","totalResults": 70,"articles": [{"source": {"id": "the-wall-street-journal","name": "The Wall Street Journal"},"author": "Alison Sider","title": "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal","description": "U.S. carriers are planning to operate smaller companies with fewer flights and employees","url": "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600","urlToImage": "https://images.wsj.net/im-194495/social","publishedAt": "2020-06-07T14:11:39Z","content": "Federal stimulus money for airlines is keeping them afloat through the coronavirus pandemic, but its not proving to be enough to sustain the industry at its pre-pandemic size.Carriers say they will… [+320 chars]"}]}`

	handler := func(w http.ResponseWriter, req *http.Request) {

		if req.URL.String() == `/top-headlines?country=us&apiKey=testKey&language=en&page=1&pageSize=1&category=business` {
			io.WriteString(w, expected)
		} else {
			io.WriteString(w, "Bad request")
		}
	}

	req, _ := http.NewRequest("GET", "/top-headlines", nil)

	responseRecorder := httptest.NewRecorder()

	// ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

	// 	if req.URL.String() == `/top-headlines?country=us&apiKey=testKey&language=en&page=1&pageSize=1&category=business` {
	// 		io.WriteString(w, expected)
	// 	} else {
	// 		io.WriteString(w, "Bad request")
	// 	}
	// }))

	//defer ts.Close()

	//resp := responseRecorder.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	// Check the response body is what we expect.

	// check with postman what expected json should be returned

	topHeadlinesReq := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
		Page:     1,
		PageSize: 1,
		// Sources:  []string{"bla", "blo"},
	}

	newsClient.BaseURL, _ = url.Parse(httptest.)

	// Act
	// execute client method and save result
	response, _ := newsClient.GetTopHeadlines(topHeadlinesReq)

	// Assert
	// Assert result against expected url output
	responseBs, _ := json.Marshal(response)

	if string(responseBs) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response, expected)
	}

}
