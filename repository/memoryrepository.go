// Package repository contains all the repository spec and implementations
package repository

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"github.com/springwiz/articlemaster/model"
)

// Pure memory based implementation of repository
type MemoryRepository struct {
	// Article storage map
	ArticleMap map[uint64]model.Article

	// Tag storage map index over article map
	TagMap map[string]map[string]map[uint64]struct{}

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
		TagMap:     make(map[string]map[string]map[uint64]struct{}),
		Logger:     log.New(os.Stdout, "github.com/springwiz/repository/MemoryRepository:", log.Ldate|log.Ltime),
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
		tempDateString := strings.ReplaceAll(article.DateString, "-", "")
		if m.TagMap[tag] == nil {
			m.TagMap[tag] = make(map[string]map[uint64]struct{})
		}
		if m.TagMap[tag][tempDateString] == nil {
			m.TagMap[tag][tempDateString] = make(map[uint64]struct{})
		}
		m.TagMap[tag][tempDateString][article.Id] = exists
	}
	return nil
}

// Query method for querying tags
func (m MemoryRepository) GetArticlesByTagDate(tagString string, date string) (*model.Tag, error) {
	m.Logger.Println("Entry: Get articles by Tag: ", tagString)
	tag := new(model.Tag)
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	val, exists := m.TagMap[tagString][date]
	tag.Count = uint64(len(m.TagMap[tagString][date]))
	if exists {
		tag.Tag = tagString
		sliceKeys := make([]uint64, 0)
		for key := range val {
			sliceKeys = append(sliceKeys, key)
			article, _ := m.ArticleMap[key]
			if !exists {
				delete(val, key)
			}
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
		sort.Slice(sliceKeys, func(i, j int) bool { return sliceKeys[i] < sliceKeys[j] })
		var higherIndx int
		if len(sliceKeys) >= 10 {
			higherIndx = 10
		} else {
			higherIndx = len(sliceKeys)
		}
		tag.Articles = sliceKeys[0:higherIndx]
		return tag, nil
	}
	return nil, fmt.Errorf("tag %s not found", tagString)
}
