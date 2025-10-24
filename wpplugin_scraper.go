package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

func scrape(plugin_collector *colly.Collector, pages *int, url *string, out_dst *string) {
	//finds new plugins or themes, matches on plugin entries and download button
	plugin_collector.OnHTML("h3[class=entry-title], div[class=plugin-actions], li.status-publish, a[id=wporg-theme-button-download]", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		if link == "" {
			link = e.Attr("href")
		}
		plugin_collector.Visit(link)
	})

	plugin_collector.OnResponse(func(e *colly.Response) {
		fmt.Println("Visiting: ", e.Request.URL)
		//fmt.Println("Debug: ", string(e.Body))
		if !strings.HasSuffix(e.Request.URL.Path, ".zip") {
			//not a package, ignore
			return
		}

		file_slice := strings.Split(e.Request.URL.Path, "/")
		file := file_slice[len(file_slice)-1]
		path := path.Join(*out_dst, file)
		err := os.WriteFile(path, e.Body, 0644)
		if err != nil {
			return
		}
	})
	plugin_collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	if *pages == 0 {
		plugin_collector.Visit(*url)
	} else {
		for i := 1; i <= *pages; i++ {
			//fmt.Println(fmt.Sprintf("Checking: %spage/%d/", *url, i))
			target := strings.TrimSuffix(*url, "/")
			plugin_collector.Visit(fmt.Sprintf("%s/page/%d/", target, i))
		}
	}

}

func main() {

	cwd, _ := os.Getwd()

	pages := flag.Int("pages", 5, "Number of pages to scrape.")
	url := flag.String("url", "", "Wordpress url to scrape for plugins/themes. This can be something like https://wordpress.org/plugins/browse/popular or https://wordpress.org/plugins/search/testing/")
	out_dst := flag.String("out", cwd, "Path for storing downloads")

	flag.Parse()

	if *url == "" {
		fmt.Println("Error: url argument is mandatory.")
		flag.Usage()
		os.Exit(1)
	}

	plugin_collector := colly.NewCollector(
		colly.AllowedDomains("wordpress.org", "downloads.wordpress.org"),
	)
	plugin_collector.MaxBodySize = 0

	scrape(plugin_collector, pages, url, out_dst)

}
