package pkg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"time"
)

func FileForPage(Directory,FileName string) error{
	Way := fmt.Sprintf("/home/zhora/Desktop/goolang-book/L2_WB/develop/dev09/%v/%v",Directory,FileName)
	file, err := os.Create(Way)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func GetAllLinks(address string) map[string]bool {
	links := make(map[string]bool)

	parsed, _ := url.Parse(address)
	host := parsed.Hostname()

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(address)
	if err != nil || resp == nil {
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Поиск всех возможных ссылок

	doc.Find("a").Each(func(index int, element *goquery.Selection){

		href, _ := element.Attr("href")
		parsed, err = url.Parse(href)

		if err != nil || parsed.Path == "" {
			return
		}
		protocol := "https"

		newLink := fmt.Sprintf("%s://%s%s", protocol, host, parsed.Path)

		links[newLink] = true

	})
		return links
}