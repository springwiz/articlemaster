package test

import (
	"github.com/springwiz/articlemaster/model"
	"github.com/springwiz/articlemaster/repository"
	"testing"
	"time"
)

func TestGetArticle(t *testing.T) {
	savedArticle := &model.Article{
		Id:         1,
		Title:      "latest science shows that potato chips are better for you than sugar",
		Body:       "some text, potentially containing simple markup about how potato chips are great",
		Date:       time.Now(),
		Tags:       []string{"health", "fitness", "science"},
		DateString: "2019-01-15",
	}
	repository.GetInstance().SaveArticle(savedArticle)
	getArticle, _ := repository.GetInstance().GetArticle(1)
	if getArticle.Id != savedArticle.Id {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", savedArticle.Id, getArticle.Id)
	}
}

func TestSaveArticle(t *testing.T) {
	savedArticle := &model.Article{
		Id:         1,
		Title:      "latest science shows that potato chips are better for you than sugar",
		Body:       "some text, potentially containing simple markup about how potato chips are great",
		Date:       time.Now(),
		Tags:       []string{"health", "fitness", "science"},
		DateString: "2019-01-15",
	}
	err := repository.GetInstance().SaveArticle(savedArticle)
	if err != nil {
		t.Errorf("Test failed, error thrown while saving %s", err.Error())
	}
}

func TestGetArticlesByTagDate(t *testing.T) {
	savedArticle := &model.Article{
		Id:         1,
		Title:      "latest science shows that potato chips are better for you than sugar",
		Body:       "some text, potentially containing simple markup about how potato chips are great",
		Date:       time.Now(),
		Tags:       []string{"health", "fitness", "science"},
		DateString: "2019-01-15",
	}
	repository.GetInstance().SaveArticle(savedArticle)
	tag, _ := repository.GetInstance().GetArticlesByTagDate("health", "20190115")
	test := false
	for _, id := range tag.Articles {
		if id == savedArticle.Id {
			test = true
		}
	}
	if !test {
		t.Errorf("Test failed, article id %d not found", savedArticle.Id)
	}
}
