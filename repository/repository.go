// Package repository contains all the repository spec and implementations
package repository

import (
	"github.com/springwiz/articlemaster/model"
)

// Defines the repository specification
type Repository interface {
	// Get method for retrieving the articles
	GetArticle(Id uint64) (*model.Article, error)

	// Save method for saving the articles
	SaveArticle(article model.Article) error

	// Query method for querying tags
	GetArticlesByTagDate(tag string, date string) (*model.Tag, error)
}
