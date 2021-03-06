// Package handler provides with handler functions for handling
// the various HTTP Requests
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/springwiz/articlemaster/model"
	"github.com/springwiz/articlemaster/repository"
)

// GET /articles/{id}
// implements and returns the GET Article Handler function
func GetArticleByIdHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Article Id: ", vars["id"])
		var err1 error
		id, err1 := strconv.ParseUint(vars["id"], 10, 64)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		article, err1 := repository.GetInstance().GetArticle(id)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		articleBytes, err := json.Marshal(article)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		log.Println("Response published: ", string(articleBytes))
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(string(articleBytes))
	}
}

//  GET /tags/{tagName}/{date}
// implements and returns the GET Tags Handler function
func GetArticlesByTagDateHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Tag Name: ", vars["tagName"])
		log.Println("Tag Date: ", vars["date"])

		tag, err1 := repository.GetInstance().GetArticlesByTagDate(vars["tagName"], vars["date"])
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}

		tagBytes, err := json.Marshal(tag)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		log.Println("Response published: ", string(tagBytes))
		json.NewEncoder(w).Encode(string(tagBytes))
	}
}

//  POST /articles
// implements and returns the POST Article Handler function
func PostArticleHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		article := model.NewBlankArticle()
		err1 := json.NewDecoder(r.Body).Decode(&article)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		article.Date, _ = time.Parse("yyyy-MM-dd", article.DateString)
		err := repository.GetInstance().SaveArticle(article)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err.Error()))
			log.Printf("Error while parsing %s", err)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(model.NewException("", "Success"))
	}
}

//  PUT /articles
// implements and returns the PUT Article Handler function
func PutArticleHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Article Id: ", vars["id"])
		article := &model.Article{}
		err1 := json.NewDecoder(r.Body).Decode(article)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		article, err1 = model.NewArticle(vars["id"], article.Title, article.Body, article.Tags, article.DateString)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			log.Printf("Error while parsing %s", err1)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		err2 := repository.GetInstance().SaveArticle(article)
		if err2 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err2.Error()))
			log.Printf("Error while parsing %s", err2)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(string(errRes))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(model.NewException("", "Success"))
	}
}
