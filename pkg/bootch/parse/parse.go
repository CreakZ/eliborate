package parse

import (
	"eliborate/internal/models/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	googleAPI   = "GoogleAPI"
	chitaiGorod = "Читай Город"
	livelib     = "Livelib"

	GoogleAPIUrl   = "https://www.googleapis.com/books/v1/volumes?q=isbn:"
	ChitaiGorodUrl = "https://www.chitai-gorod.ru/search?phrase="
	LivelibUrl     = "https://www.livelib.ru/find/books/"
)

func ParseBookInfoFromGoogleBookApi(isbn string) (dto.BookInfo, error) {
	URL := GoogleAPIUrl + isbn

	resp, err := http.Get(URL)
	if err != nil {
		return dto.BookInfo{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	info := struct {
		Items      int           `json:"totalItems"`
		VolumeInfo []interface{} `json:"items"`
	}{
		Items:      0,
		VolumeInfo: nil,
	}

	if err = json.Unmarshal(body, &info); err != nil {
		return dto.BookInfo{}, err
	}

	if info.Items == 0 {
		return dto.BookInfo{}, newErrorWithCaller(googleAPI, fmt.Sprintf("no books with %s isbn found", isbn))
	}

	volumeInfo, ok := info.VolumeInfo[0].(map[string]interface{})["volumeInfo"].(map[string]interface{})
	if !ok {
		return dto.BookInfo{}, newErrorWithCaller(googleAPI, "response json has unexpected structure")
	}

	var book dto.BookInfo

	book.Title, ok = volumeInfo["title"].(string)
	if !ok {
		book.Title = ""
	}

	// Получение информации об авторах
	temp, exists := volumeInfo["authors"].([]interface{})
	if !exists {
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
	} else {
		book.Description = &desc
	}

	return book, nil
}

func ParseBookInfoFromChitaiGorod(isbn string) (dto.BookInfo, error) {
	url := ChitaiGorodUrl + isbn

	resp, err := http.Get(url)
	if err != nil {
		return dto.BookInfo{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return dto.BookInfo{}, err
	}

	// Проверка на пустой результат
	val, _ := doc.Find("h4").Attr("class")
	if val == "catalog-empty-result__header" {
		return dto.BookInfo{}, newErrorWithCaller(chitaiGorod, fmt.Sprintf("no books with '%s' isbn found", isbn))
	}

	var (
		link string
		book dto.BookInfo
	)

	doc.Find("article").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			value, exists := s.Find("a").Attr("href")
			if exists {
				link = value
				return false
			}
			return true
		})

	bookUrl := "https://www.chitai-gorod.ru" + link

	resp, err = http.Get(bookUrl)
	if err != nil {
		return dto.BookInfo{}, err
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return dto.BookInfo{}, err
	}

	// Получение информации о названии
	doc.Find("h1").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			class, _ := s.Attr("class")

			if class == "detail-product__header-title" {
				book.Title = strings.TrimSpace(s.Text())
				return false
			}
			return true
		})

	// Получение информации об авторах
	doc.Find("a").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			itemprop, _ := s.Attr("class")

			if itemprop == "product-info-authors__author" {
				author, _ := strings.CutSuffix(strings.TrimSpace(s.Text()), ",")
				book.Authors = append(book.Authors, author)

				return false
			}
			return true
		})

	// Получение информации об обложке
	doc.Find("img").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			if s.AttrOr("class", "") == "product-info-gallery__poster" {
				src, _ := s.Attr("src")
				book.CoverUrls = append(book.CoverUrls, src)

				return false
			}
			return true
		})

	// Получение информации об описании книги
	if desc := doc.Find("article"); desc != nil {
		text := strings.TrimSpace(desc.Text())
		book.Description = &text
	}

	return book, nil
}

func ParseBookInfoFromLivelib(isbn string) (dto.BookInfo, error) {
	url := LivelibUrl + isbn

	resp, err := http.Get(url)
	if err != nil {
		return dto.BookInfo{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return dto.BookInfo{}, err
	}

	var book dto.BookInfo

	// Получение информации о названии и авторе
	doc.Find("a").
		Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			switch class {
			case "title":
				book.Title = s.Text()

			case "description":
				book.Authors = append(book.Authors, s.Text())
			}
		})

	// Получение информации об изображении и описании книги
	doc.Find("span").
		Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			switch class {
			case "description":
				text := s.Text()
				book.Description = &text

			case "object-cover":
				stl, _ := s.Attr("style")

				leftBound := strings.IndexByte(stl, '(')
				rightBound := strings.IndexByte(stl, ')')

				if leftBound != -1 && rightBound != -1 && leftBound < rightBound {
					book.CoverUrls = append(book.CoverUrls, stl[leftBound+1:rightBound])
				}
			}
		})

	return book, nil
}

func newErrorWithCaller(caller, msg string) error {
	return fmt.Errorf("%s: %s", caller, msg)
}
