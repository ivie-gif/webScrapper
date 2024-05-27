package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

// Fetches the HTML content of a given URL
func fetchURL(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Parses the HTML and extracts data
func parseHTML(response *http.Response) {
	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			t := tokenizer.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						fmt.Println("Link found:", attr.Val)
					}
				}
			}
		}
	}
}

func main() {
	url := "https://go.dev/"
	fmt.Println("Fetching URL:", url)

	response, err := fetchURL(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	parseHTML(response)
}
