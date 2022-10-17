package utils

import (
	"bookshelf-app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBookInfoByIsbn(ISBN string) *models.BookModel {
	apiUrl := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%v&jscmd=data&format=json", ISBN)
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}

	datas := data["ISBN:"+ISBN].(map[string]interface{})
	title := datas["title"]
	publish_date := datas["publish_date"]
	authors := datas["authors"].([]interface{})
	author := authors[0].(map[string]interface{})["name"]
	pages := datas["number_of_pages"]
	if pages == nil {
		pages = 0.0
	}

	return &models.BookModel{
		ISBN: ISBN, Title: title.(string),
		Author:    author.(string),
		Published: publish_date.(string),
		Pages:     pages.(float64),
	}
}
