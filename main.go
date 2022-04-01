package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summery  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	//var c chan string
	pages := getPages()
	fmt.Println(pages)
	for i := 0; i < pages; i++ {
		getPage(i)
	}
}

// getPage get elements in page
func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting this url :", pageURL)
	resp, err := http.Get(pageURL)
	chkErr(err)
	chkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	chkErr(err)

	doc.Find(".result").Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find(".jobTitle").Text()
		location := card.Find(".companyLocation").Text()
		salary, _ := card.Find(".salary-snippet").Attr("aria-label")
		if salary != "" {
			fmt.Println(id)
			fmt.Println(title)
			fmt.Println(location)
			fmt.Println(salary)
		}
	})
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
