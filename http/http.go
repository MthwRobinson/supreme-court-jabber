// The module that contains code for fetching content over http
package fetch_arguments

import (
	"fmt"
	"log"
	http "net/http"

	goquery "github.com/PuerkitoBio/goquery"
)

const supremeCourtURL = "https://www.supremecourt.gov/oral_arguments/argument_transcript/2020"

// Pulls the argument number from the row in the transcript table
func getArgumentNumber(s *goquery.Selection) {
	argumentLink := s.Find("a")
	argumentNumber := argumentLink.Text()

	if argumentNumber != "" {
		link, _ := argumentLink.Attr("href")
		fmt.Printf("%s: %s\n", argumentNumber, link)
	}
}

func GetOralArguments() {
	client := &http.Client{}

	response, err := client.Get(supremeCourtURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		getArgumentNumber(s)

	})
}
