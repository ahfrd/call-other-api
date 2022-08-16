package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestDataTopUp struct {
	Username string `json:"username"`
	Ballance int    `json:"ballance"`
}

type ResponseApi struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyUrl  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}
type GeneralResponse struct {
	Code string      `json:"response_code"`
	Msg  string      `json:"response_msg"`
	Data interface{} `json:"response_data"`
}

func FormatCallApiResult(url string) ([]ResponseApi, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var resApi []ResponseApi
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(responseData), &resApi)
	return resApi, err
}
func FormatCallApiRow(url string) (ResponseApi, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var resApi ResponseApi
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(responseData), &resApi)

	return resApi, err
}
func TopUpApiCall(url string, username string, ballance int) (GeneralResponse, error) {
	bodyBeforeJson := &RequestDataTopUp{
		Username: username,
		Ballance: ballance,
	}
	postBody, _ := json.Marshal(bodyBeforeJson)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var resApi GeneralResponse
	json.Unmarshal([]byte(body), &resApi)

	return resApi, err
}
