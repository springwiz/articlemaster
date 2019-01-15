// Package model contains all the entities
package model

import (
	"time"
)

// The article entity
type Article struct {
	// article id
	Id uint64 `json:"id"`

	// article title
	Title string `json:"title"`

	// article description
	Body string `json:"body"`

	// article tags
	Tags []string `json:"tags"`

	// last updated article date
	Date time.Time `json:"-"`

	DateString string `json:"date"`
}

// create new article
func NewArticle(title string, body string, tags []string) *Article {
	newArticle := &Article{
		Id:    LastReceivedArticleId,
		Title: title,
		Body:  body,
		Date:  time.Now(),
		Tags:  make([]string, 0),
	}
	LastReceivedArticleId++
	return newArticle
}

// create new blank article
func NewBlankArticle() *Article {
	newArticle := &Article{
		Id: LastReceivedArticleId,
	}
	LastReceivedArticleId += 1
	return newArticle
}

// article id counter
var LastReceivedArticleId uint64 = 1
