package main

import (
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
}

func getPages() int {
	resp, err := http.Get(baseURL)
	chkErr(err)

	return 0
}

func chkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func chkCode(resp http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", resp.StatusCode)
	}
}
