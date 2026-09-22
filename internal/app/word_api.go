package internal/app

import (

	"net/http"
	"net/url"
	"time"
	"fmt"
	"path"
	"strconv"
	"context"
	"encoding/json"
	"log/slog"
)


type APIService struct {
	client *http.Client
	baseUrl string
}


//Create a new client with 2min timeout incase word dictonary API is unreachable
func NewClient() *APIService{
	return &APIService{
		baseUrl: "https://random-word-api.herokuapp.com",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}


//Call API to get a single word with a user defined difficulty
func (s *APIService) GetWord(ctx context.Context, diffLvel int) (string,error) {
	

	//Perform validation on difficulty
	if (diffLvel <= 0 || diffLvel > 5){

		return "", fmt.Errorf("Difficulty must be between 1 to 5")	

	}

	//API request
	baseUrl, err := url.Parse(s.baseUrl)

	if err != nil {
		return "", fmt.Errorf("Something wrong with parsing URL:%w", err)
	}


	//path variable : /word
	endpoint := "word"
	baseUrl.Path = path.Join(baseUrl.Path,endpoint)


	//query params 
	query := url.Values{}
	query.Add("diff",strconv.Itoa(diffLvel))

	//Combine path and url params to prepare url for GET API Call
	baseUrl.RawQuery = query.Encode() //diff=0 

	slog.InfoContext(ctx, "prepared url","url", baseUrl.String())
	
	//Call GET API
	
	//Build request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		baseUrl.String(),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create request: %w",err)
	}

	//Call request and read response
	res, err := s.client.Do(req)


	if err != nil {
		return "", fmt.Errorf("failed to call word API: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status: %d", res.StatusCode)
	}
		
	//Deserialize JSON body and get the get the word
	var words []string
	
	err = json.NewDecoder(res.Body).Decode(&words)
	
	if err != nil {
		return "" , fmt.Errorf("failed to decode response: %w", err)
	}

	if len(words) == 0{
		return "", fmt.Errorf("API returned no words")
	}

	return words[0], nil
	
}
