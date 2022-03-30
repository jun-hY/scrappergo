package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

// getPages is check the pages can response
func getPages() int {
	pages := 0
	resp, err := http.Get(baseURL)
	chkErr(err)
	chkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	chkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// chkErr is check the Error
func chkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// chkCode is check the http status Code
func chkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", resp.StatusCode)
	}
}

//
