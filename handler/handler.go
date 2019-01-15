// Package handler provides with handler functions for handling
// the various HTTP Requests
package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/springwiz/articlemaster/model"
	"github.com/springwiz/articlemaster/repository"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// logger instance
var Logger *log.Logger

// initialize the logger for the handler
func init() {
	Logger = log.New(os.Stdout, "handler", log.Ldate|log.Ltime)
}

// GET /articles/{id}
// implements and returns the GET Article Handler function
func GetArticleByIdHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		Logger.Println("Article Id: ", vars["id"])
		var err1 error
		id, err1 := strconv.ParseUint(vars["id"], 10, 64)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		article, err1 := repository.GetInstance().GetArticle(id)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		articleBytes, err := json.Marshal(article)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		Logger.Println("Response published: ", string(articleBytes))
		json.NewEncoder(w).Encode(string(articleBytes))
	}
}

//  GET /tags/{tagName}/{date}
// implements and returns the GET Tags Handler function
func GetArticlesByTagDateHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		Logger.Println("Tag Name: ", vars["tagName"])
		Logger.Println("Tag Date: ", vars["date"])

		tag, err1 := repository.GetInstance().GetArticlesByTagDate(vars["tagName"], vars["date"])
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}

		tagBytes, err := json.Marshal(tag)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00001", "Parse Error"))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		Logger.Println("Response published: ", string(tagBytes))
		json.NewEncoder(w).Encode(string(tagBytes))
	}
}

//  POST /articles
// implements and returns the POST Article Handler function
func PostArticleHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		Logger.Println("Article Id: ", id)
		article := model.NewBlankArticle()
		err1 := json.NewDecoder(r.Body).Decode(&article)
		if err1 != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err1.Error()))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		article.Date, _ = time.Parse("yyyy-MM-dd", article.DateString)
		err := repository.GetInstance().SaveArticle(article)
		if err != nil {
			errRes, _ := json.Marshal(model.NewException("PE00002", err.Error()))
			Logger.Println("Error while parsing: ", errRes)
			json.NewEncoder(w).Encode(errRes)
			return
		}
		json.NewEncoder(w).Encode(model.NewException("", "Success"))
	}
}
