// Package repository contains all the repository spec and implementations
package repository

import (
	"fmt"
	"github.com/springwiz/articlemaster/model"
	"log"
	"os"
	"strings"
	"sync"
)

// Pure memory based implementation of repository
type MemoryRepository struct {
	// Article storage map
	ArticleMap map[uint64]model.Article

	// Tag storage map index over article map
	TagMap map[string]map[uint64]struct{}

	// Locking
	Lock sync.RWMutex

	// Logger
	Logger *log.Logger
}

// Singleton instance
var instance Repository

// Singleton access function
func GetInstance() Repository {
	return instance
}

// Initialize the Memory Repository Singleton
func init() {
	instance = MemoryRepository{
		ArticleMap: make(map[uint64]model.Article),
		TagMap:     make(map[string]map[uint64]struct{}),
		Logger:     log.New(os.Stdout, "MemoryRepository", log.Ldate|log.Ltime),
	}
}

// Get method for retrieving the articles
func (m MemoryRepository) GetArticle(Id uint64) (*model.Article, error) {
	m.Logger.Println("Entry: Get article id: ", Id)
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	article, exists := m.ArticleMap[Id]
	if exists {
		return &article, nil
	} else {
		return nil, fmt.Errorf("article %d not found", Id)
	}
}

// Save method for saving the articles
func (m MemoryRepository) SaveArticle(article *model.Article) error {
	m.Logger.Println("Entry: Save article id: ", article.Id)
	var exists = struct{}{}
	m.Lock.Lock()
	defer m.Lock.Unlock()
	m.ArticleMap[article.Id] = *article
	for _, tag := range article.Tags {
		if m.TagMap[tag] == nil {
			m.TagMap[tag] = make(map[uint64]struct{})
		}
		m.TagMap[tag][article.Id] = exists
	}
	return nil
}

// Query method for querying tags
func (m MemoryRepository) GetArticlesByTagDate(tagString string, date string) (*model.Tag, error) {
	m.Logger.Println("Entry: Get articles by Tag: ", tagString)
	year := date[0:4]
	month := date[4:6]
	day := date[6:8]
	tag := new(model.Tag)
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	val, exists := m.TagMap[tagString]
	exists1 := struct{}{}

	articleDayCount := 0
	articleMap := make(map[uint64]struct{})
	for _, v := range m.TagMap {
		for id, _ := range v {
			article := m.ArticleMap[id]
			dateEle := strings.Split(article.DateString, "-")
			if dateEle[0] == year && dateEle[1] == month && dateEle[2] == day {
				tag.Count++
				if articleDayCount < 10 {
					articleMap[id] = exists1
				}
				articleDayCount++
			}
		}
	}
	for k, _ := range articleMap {
		tag.Articles = append(tag.Articles, k)
	}
	if exists {
		tag.Tag = tagString
		for key := range val {
			article, _ := m.ArticleMap[key]
			if !exists {
				delete(val, key)
			}
			dateEle := strings.Split(article.DateString, "-")
			if dateEle[0] == year && dateEle[1] == month && dateEle[2] == day {
				for _, articleTag := range article.Tags {
					tagExists := false
					for _, tagStr := range tag.RelatedTags {
						if articleTag == tagStr {
							tagExists = true
						}
					}
					if !tagExists {
						tag.RelatedTags = append(tag.RelatedTags, articleTag)
					}
				}
			}
		}
		return tag, nil
	} else {
		return nil, fmt.Errorf("tag %s not found", tagString)
	}
}
