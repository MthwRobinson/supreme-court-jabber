// The main module to kick of the data ingest workflow
package main

import (
	"github.com/MthwRobinson/supreme-court-jabber/data/argument_fetcher"
)

func main() {
	fetch_arguments.GetOralArguments()
}
