// The module that contains code for fetching content over http
package fetch_arguments

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/http/httputil"
)

const supremeCourtURL = "https://www.supremecourt.gov/oral_arguments/argument_transcript/2021"

func GetOralArguments() {
	client := &http.Client{}

	response, err := client.Get(supremeCourtURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	z := html.NewTokenizer(response.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			fmt.Println("Done!")
			return
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						fmt.Println(attr.Val)
					}
				}
			}

		}

	}

}
