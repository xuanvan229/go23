package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Data struct {
	Title string
	Link  string
}

func writeFile(data []Data) {

	csvFile, err := os.Create("data.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	_ = csvwriter.Write([]string{"Title", "Link"})

	for _, empRow := range data {
		row := []string{empRow.Title, empRow.Link}
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvFile.Close()

}

func main() {
	baseUrl := "https://vnexpress.net"
	result := []Data{}
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("vnexpress.net"),
	)

	c.OnHTML("ul.parent > li >  a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(baseUrl + link))
	})

	c.OnHTML(".title-news > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		result = append(result, Data{
			Title: strings.TrimSuffix(e.Text, "\n"),
			Link:  link,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://vnexpress.net/the-gioi")
	c.Wait()
	writeFile(result)
	fmt.Println("Done!!!")
}
