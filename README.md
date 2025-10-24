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

## Example

```bash
user@wp_plugin_downloader % go run . -pages=2 -out ./downloads -url=https://wordpress.org/themes/ 
Visiting:  https://wordpress.org/themes/
Visiting:  https://wordpress.org/themes/twentytwentyfive/
Visiting:  https://downloads.wordpress.org/theme/twentytwentyfive.1.3.zip
Visiting:  https://wordpress.org/themes/hello-elementor/
Visiting:  https://downloads.wordpress.org/theme/hello-elementor.3.4.4.zip
Visiting:  https://wordpress.org/themes/astra/
Visiting:  https://downloads.wordpress.org/theme/astra.4.11.13.zip
...snip....
Visiting:  https://downloads.wordpress.org/theme/envo-one.1.0.4.zip
Visiting:  https://wordpress.org/themes/kubio/
Visiting:  https://downloads.wordpress.org/theme/kubio.1.0.56.zip
Visiting:  https://wordpress.org/themes/yith-wonder/
Visiting:  https://downloads.wordpress.org/theme/yith-wonder.2.1.1.zip
Visiting:  https://wordpress.org/themes/futurio-storefront/
Visiting:  https://downloads.wordpress.org/theme/futurio-storefront.1.0.3.zip
Visiting:  https://wordpress.org/themes/twentynineteen/
Visiting:  https://downloads.wordpress.org/theme/twentynineteen.3.1.zip
user@wp_plugin_downloader % ls downloads 
astra.4.11.13.zip		hello-elementor.3.4.4.zip	twentyseventeen.3.9.zip
blocksy.2.1.17.zip		kadence.1.3.6.zip		twentytwenty.2.9.zip
bluehost-blueprint.1.0.0.zip	kubio.1.0.56.zip		twentytwentyfive.1.3.zip
envo-one.1.0.4.zip		neve.4.1.4.zip			twentytwentyfour.1.3.zip
extendable.2.0.29.zip		oceanwp.4.1.3.zip		twentytwentyone.2.6.zip
futurio-storefront.1.0.3.zip	popularfx.1.2.7.zip		twentytwentythree.1.6.zip
generatepress.3.6.0.zip		royal-elementor-kit.1.0.139.zip	twentytwentytwo.2.0.zip
hello-biz.1.2.0.zip		twentynineteen.3.1.zip		yith-wonder.2.1.1.zip
```