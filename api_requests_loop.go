package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Result struct {
	currentPage int           `json:"current_page"`
	data        []interface{} `json:"data"`
	from        int           `json:"from"`
	lastPage    int           `json:"last_page"`
	nextPageUrl string        `json:"next_page_url"`
	perPage     int           `json:"per_page"`
	prevPageUrl string        `json:"prev_page_url"`
	to          int           `json:"to"`
	total       int           `json:"total"`
}

func main() {
	for i := 1; i <= 254; i++ {
		concatenated := fmt.Sprint("{{url}}/ci/?per_page=70&page=1&status=OPE&show_ip=1&search=10.201.245.", i)
		call(concatenated, "GET")
	}
}
func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Authorization", "tuo token JWT")
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result Result
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(url, result.total)
	defer response.Body.Close()
	return nil
}
