package models

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Article struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Url     string        `bson: "url"`
	Title   string
	Content template.HTML `bson: "content"`
	Score   int64         `bson: "score"`
	ViewSet mgo.DBRef     `bson: "viewCollection"`
}

// prepare, call, and read Readability API
func SetReadableContent(article *Article) error {
	urlString := article.Url
	// construct the Readability request url
	queryString := "https://readability.com/api/content/v1/parser?url=" +
		url.QueryEscape(urlString) +
		"&token=" + os.Getenv("JUST_READ_READABILITY_TOKEN")

	// send the request
	res, err := http.Get(queryString)
	if err != nil {
		return err
	}

	// convert response into a byte array
	json_array, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// close response body
	res.Body.Close()

	// convert byte array response into object map
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(json_array, &objmap)
	if err != nil {
		return err
	}

	var readabilityResponse ReadabilityResponse
	err = json.Unmarshal(json_array, &readabilityResponse)
	if err != nil || readabilityResponse.Error == "true" {
		return err
	}

	article.Title = readabilityResponse.Title
	article.Content = template.HTML(readabilityResponse.Content)

	return nil
}
