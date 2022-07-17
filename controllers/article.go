package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/darahayes/go-boom"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var article entity.Article
	json.Unmarshal(requestBody, &article)

	err := database.Connector.Create(article).Error
	if err != nil {
		boom.BadData(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {

}
