package libraries

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"yurii-lib/internal/models/dto"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "https://www.livelib.ru/find/books/"

func GetBookWithLivelib(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	endpoint := fmt.Sprintf("%s%s", baseUrl, isbn)

	resp, err := http.Get(endpoint)
	if err != nil {
		errs <- err
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		errs <- err
		return
	}

	// check if there is no results

	var book dto.BookInfo

	innerWG := sync.WaitGroup{}

	// Получение информации о названии и авторе
	innerWG.Add(1)
	go func(innerWg *sync.WaitGroup) {
		defer innerWG.Done()

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			switch class {
			case "title":
				book.Title = s.Text()

			case "description":
				book.Authors = append(book.Authors, s.Text())
			}
		})
	}(&innerWG)

	// Получение информации об изображении и описании книги
	innerWG.Add(1)
	go func(innerWg *sync.WaitGroup) {
		defer innerWG.Done()

		doc.Find("span").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			switch class {
			case "description":
				text := s.Text()
				book.Description = &text

			case "object-cover":
				stl, _ := s.Attr("style")

				// cutting prefix and suffix to get cover url
				stl, _ = strings.CutPrefix(stl, "background:url(")
				stl, _ = strings.CutSuffix(stl, ") no-repeat;")

				book.CoverURL = &stl
			}
		})
	}(&innerWG)

	innerWG.Wait()

	books <- book
}
