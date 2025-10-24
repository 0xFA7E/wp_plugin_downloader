# wp_plugin_downloader
A simple scraper for downloading wordpress plugins and themes. Built with the idea to mirror a large amount of plugins for security testing.

# Usage

```
  -out string
    	Path for storing downloads (default CWD)
  -pages int
    	Number of pages to scrape. (default 5)
  -url string
    	Wordpress url to scrape for plugins/themes. This can be something like https://wordpress.org/plugins/browse/popular or https://wordpress.org/plugins/search/testing/
```