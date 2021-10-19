// The module that contains code for fetching content over http
package fetch_arguments

import (
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
}
