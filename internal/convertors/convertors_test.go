package convertors_test

import (
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentFromBookPlacement(t *testing.T) {
	books := []domain.BookPlacement{}

	// TODO
	_ = books
}

func TestBookFromMeiliDocument(t *testing.T) {
	a := assert.New(t)

	book1 := map[string]interface{}{
		"id":    1,
		"title": "Phimosis",
		"authors": []interface{}{
			"Kolya Gl1nom3s",
			"Titov5 Da",
		},
		"description": "Book about all types of phimosis",
		"category":    "Something",
	}
	book2 := map[string]interface{}{
		"id":    2,
		"title": "Как правильно выпивать",
		"authors": []string{
			"Филипп Федотов",
		},
		"description": "Я выпиваю уже более 5 лет, и вот что со мной случилось",
		"category":    "Обучение",
	}
	book3 := map[string]interface{}{
		"authors":     []string{"Федор Достоевский"},
		"category":    "Художественная отечественная литература",
		"description": "Очень интересная книга",
		"id":          3,
		"title":       "Преступление и наказание",
	}

	if !a.Equal(
		domain.BookSearch{
			ID:    1,
			Title: "Phimosis",
			Authors: []string{
				"Kolya Gl1nom3s",
				"Titov5 Da",
			},
			Description: "Book about all types of phimosis",
			Category:    "Something",
		},
		convertors.BookFromMeiliDocument(book1),
	) {
		t.Errorf("failed in assertion 1")
	}

	if !a.Equal(
		domain.BookSearch{
			ID:    2,
			Title: "Как правильно выпивать",
			Authors: []string{
				"Филипп Федотов",
			},
			Description: "Я выпиваю уже более 5 лет, и вот что со мной случилось",
			Category:    "Обучение",
		},
		convertors.BookFromMeiliDocument(book2),
	) {
		t.Errorf("failed in assertion 2")
	}

	if !a.Equal(
		domain.BookSearch{
			ID:    3,
			Title: "Преступление и наказание",
			Authors: []string{
				"Федор Достоевский",
			},
			Description: "Очень интересная книга",
			Category:    "Художественная отечественная литература",
		},
		convertors.BookFromMeiliDocument(book3),
	) {
		t.Errorf("failed in assertion 3")
	}
}
