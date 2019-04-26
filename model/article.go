// Package model contains all the entities
package model

import (
	"strconv"
	"sync/atomic"
	"time"
)

// The article entity
type Article struct {
	// article id
	Id uint64 `json:"-"`

	StringId string `json:"id"`

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

// article id counter
var LastReceivedArticleId uint64 = 0

// create new article
func NewArticle(strId string, title string, body string, tags []string, strDate string) (*Article, error) {
	id, err1 := strconv.ParseUint(strId, 10, 64)
	if err1 != nil {
		return nil, err1
	}
	date, err2 := time.Parse("2006-01-02", strDate)
	if err2 != nil {
		return nil, err2
	}
	newArticle := &Article{
		Id:         id,
		StringId:   strId,
		Title:      title,
		Body:       body,
		DateString: strDate,
		Date:       date,
		Tags:       tags,
	}
	return newArticle, nil
}

// create new blank article
func NewBlankArticle() *Article {
	newArticle := &Article{
		Id:       atomic.AddUint64(&LastReceivedArticleId, 1),
		StringId: strconv.FormatUint(LastReceivedArticleId, 10),
	}
	return newArticle
}
