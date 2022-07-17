package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var article entity.Article
	json.Unmarshal(requestBody, &article)

	err := database.Connector.Create(article).Error
	if err != nil {
		boom.BadData(w, err)
	}

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(article)
	}
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article entity.Article
	err := database.Connector.First(&article, key).Error
	if err != nil {
		boom.NotFound(w, err)
	}

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(article)
	}

}
