package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.olevod.com/" // Specify the URL you want to crawl
	links, err := fetchLinks(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

func fetchLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	var links []string
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						links = append(links, a.Val)
						break
					}
				}
			}
		}
	}
}
