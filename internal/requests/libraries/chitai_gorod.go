package libraries

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"yurii-lib/internal/models/dto"

	"github.com/PuerkitoBio/goquery"
)

const ChitaiGorod = "https://www.chitai-gorod.ru/search?phrase="

func Cleanify(s string) string {
	var left int

	i := 0
	for i < len(s) {
		if s[i] == 10 || s[i] == 32 {
			i++
		} else {
			break
		}
	}

	if i == len(s) {
		return ""
	}

	left = i
	i = len(s) - 1

	for i > -1 {
		if s[i] == 10 || s[i] == 32 {
			i--
		} else {
			break
		}
	}

	return s[left : i+1]
}

func GetBookWithChitaiGorod(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	url := ChitaiGorod + isbn

	req, err := http.Get(url)
	if err != nil {
		errs <- err
		return
	}
	defer req.Body.Close()

	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		errs <- err
		return
	}

	// Проверка на пустой результат
	val, _ := doc.Find("h4").Attr("class")
	if val == "catalog-empty-result__header" {
		errs <- fmt.Errorf("no books with isbn %s found", isbn)
		return
	}

	var (
		link string
		book dto.BookInfo
	)

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Find("a").Attr("href")
		if exists {
			link = value
		}
	})

	bookUrl := fmt.Sprintf("https://www.chitai-gorod.ru%v", link)

	req, err = http.Get(bookUrl)
	if err != nil {
		errs <- err
		return
	}

	doc, err = goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		errs <- err
		return
	}

	innerWG := sync.WaitGroup{}

	// Получение информации о названии
	innerWG.Add(1)
	go func(innerWG *sync.WaitGroup) {
		defer innerWG.Done()

		doc.Find("h1").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			if class == "detail-product__header-title" {
				book.Title = Cleanify(s.Text())
			}
		})
	}(&innerWG)

	// Получение информации об авторах
	innerWG.Add(1)
	go func(innerWG *sync.WaitGroup) {
		defer innerWG.Done()

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			itemprop, _ := s.Attr("class")

			if itemprop == "product-info-authors__author" {
				author, _ := strings.CutSuffix(Cleanify(s.Text()), ",")

				book.Authors = append(book.Authors, author)
			}
		})
	}(&innerWG)

	// Получение информации об обложке
	innerWG.Add(1)
	go func(innerWG *sync.WaitGroup) {
		defer innerWG.Done()

		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			if class == "product-info-gallery__poster" {
				src, _ := s.Attr("src")

				book.CoverURL = &src
			}
		})
	}(&innerWG)

	// Получение информации об описании книги
	innerWG.Add(1)
	go func(innerWG *sync.WaitGroup) {
		defer innerWG.Done()

		desc := doc.Find("article")

		if desc != nil {
			text := Cleanify(desc.Text())

			book.Description = &text
		}
	}(&innerWG)

	innerWG.Wait()

	books <- book
}
