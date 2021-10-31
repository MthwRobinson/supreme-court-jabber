// The module that contains code for fetching content over http
package fetch_arguments

import (
	"fmt"
	"log"
	http "net/http"
	"strings"

	goquery "github.com/PuerkitoBio/goquery"
)

const supremeCourtURL = "https://www.supremecourt.gov/oral_arguments/argument_transcript/2020"

type OralArgument struct {
	number string
	name   string
	date   string
	link   string
}

// Pulls the orgal arguments from a page on the Supreme Court website
func GetOralArguments() []OralArgument {
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

	var oralArguments []OralArgument
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			argumentLink, _ := s.Attr("href")
			if strings.HasSuffix(argumentLink, ".pdf") {
				argumentNumber := s.Text()

				parent := s.Parent()
				argumentName := parent.Next().Text()
				argumentDate := parent.Parent().Next().Text()

				oralArgument := OralArgument{
					number: argumentNumber,
					name:   argumentName,
					date:   argumentDate,
					link:   argumentLink,
				}
				oralArguments = append(oralArguments, oralArgument)
			}
		})

	})
	return oralArguments
}
