package libs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"yurii-lib/internal/models/dto"
)

const GoogleAPI = "https://www.googleapis.com/books/v1/volumes?q=isbn:"

func GetBookWithGoogleAPI(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	URL := GoogleAPI + isbn

	req, err := http.Get(URL)
	if err != nil {
		// some error info?
		errs <- err
		return
	}

	body, _ := io.ReadAll(req.Body)

	var info = struct {
		Items      int           `json:"totalItems"`
		VolumeInfo []interface{} `json:"items"`
	}{
		Items:      0,
		VolumeInfo: nil,
	}

	if err = json.Unmarshal(body, &info); err != nil {
		errs <- err
		return
	}

	if info.Items == 0 {
		errs <- fmt.Errorf("no books with %v ISBN found", isbn)
		return
	}

	volumeInfo := info.VolumeInfo[0].(map[string]interface{})["volumeInfo"].(map[string]interface{})

	var (
		lang string
		book dto.BookInfo
		ok   bool
	)

	book.Title, ok = volumeInfo["title"].(string)
	if !ok {
		fmt.Println("no title data provided")
		book.Title = ""
	}

	// Получение информации об авторах
	temp, exists := volumeInfo["authors"].([]interface{})
	if !exists {
		fmt.Println("no authors data provided")
		fmt.Printf("type of authors: %T\n", volumeInfo["authors"])
		book.Authors = []string{""}
	}

	for _, author := range temp {
		book.Authors = append(book.Authors, author.(string))
	}

	// Получение информации об описании

	var desc string
	desc, ok = volumeInfo["description"].(string)
	if !ok {
		fmt.Println("no description provided")
	}

	book.Description = &desc

	// Получение информации о языке
	lang, ok = volumeInfo["language"].(string)
	if !ok {
		fmt.Println("no lang data provided")
	} else {
		switch lang {
		case "ru":
			book.IsForeign = false
		default:
			book.IsForeign = true
		}
	}

	books <- book
}
