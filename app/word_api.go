package app


import (

	"net/http"
	"net/url"
	"time"
	"fmt"
	"path"
	"strconv"
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
func (s *APIService) GetWord(diffLvel int) (string,error) {
	

	//Perform validation on difficulty
	if (diffLvel <= 0 || diffLvel > 4){

		return "", fmt.Errorf("Difficulty must be between 1 to 5")	

	}

	//API request
	baseUrl, err := url.Parse(s.baseUrl)

	if err != nil {
		fmt.Errorf("Something wrong with parsing URL:%w", err)
	}


	//path variable : /word
	endpoint := "word"
	baseUrl.Path = path.Join(baseUrl.Path,endpoint)


	//query params 
	query := url.Values{}
	query.Add("diff",strconv.Itoa(diffLvel))
	baseUrl.RawQuery = query.Encode() //diff=0 

	fmt.Println("Prepared url: %s",baseUrl.String())

	
	//Combine path and url params to prepare url for GET API Call
	


	

	
}




