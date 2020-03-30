# covid19api.com Go SDK

This package provides an API and some example commands for interacting with the [Free API for Data on the Coronavirus COVID19](https://covid19api.com/).

This is a work in progress, so be wary of things broken or not performant enough.

## Installation

As simple as:

```
go get github.com/inkel/covid19/covid19api
```

## Example

I've included an example `covid19` command that you can install and execute like follows:

```
$ go install github.com/inkel/covid19/cmd/covid19

$ covid19
```

The code of this command is:

```go
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/inkel/covid19/covid19api"
)

func main() {
	s, err := covid19api.GetSummary()
	if err != nil {
		fmt.Println("Cannot fetch summary from covid19api.com: %v", err)
		os.Exit(1)
	}

	showCountry := func(slug string) bool {
		if len(os.Args) == 1 {
			return true
		}

		for _, s := range os.Args[1:] {
			if s == slug {
				return true
			}
		}

		return false
	}

	fmt.Printf("COVID-19 Summary at %v (%v ago)\n", s.Date, time.Since(s.Date).Truncate(time.Second))
	fmt.Println("You can filter by passing each country's slug you want to include in the output.")
	fmt.Println("Format for cases is new/total.\n")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "Country\tSlug\tConfirmed\tRecovered\tDeaths\t")

	for _, c := range s.Countries {
		if showCountry(c.Slug) {
			fmt.Fprintf(w, "%s\t%s\t%d/%d\t%d/%d\t%d/%d\t\n", c.Country, c.Slug, c.NewConfirmed, c.TotalConfirmed, c.NewRecovered, c.TotalRecovered, c.NewDeaths, c.TotalDeaths)
		}
	}

	w.Flush()
}
```

## License

See [LICENSE](LICENSE).
