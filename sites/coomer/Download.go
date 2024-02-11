package coomer

import "godl/console"
import "godl/structs"
import "godl/utils"
import "net/url"
import "strconv"

func Download(cache *structs.Cache, index *structs.Index, link string) bool {

	var result bool = false

	console.Group("coomer/Download")

	base, err := url.Parse(link)

	if err == nil {

		if base.Scheme == "https" && base.Host == "coomer.su" {

			scraper := structs.NewScraper(cache, &map[string]string{
				"Referer": base.Scheme + "://" + base.Host + base.Path,
			})
			scraper.Throttled = true

			count := 0

			for url, _ := range index.Downloads {

				filename := utils.ToFilename(url)
				count++

				if !index.Exists(url) && !cache.Exists("/"+filename) {

					buffer := scraper.Request(url)

					if len(buffer) > 0 {
						cache.Write("/"+filename, buffer)
					}

					index.Set(url)
					index.Write()

				} else if cache.Exists("/" + filename) {

					index.Set(url)
					index.Write()

				}

				console.Progress(strconv.Itoa(count) + " of " + strconv.Itoa(len(index.Downloads)) + " Downloads")

			}

			result = true

		}

	}

	console.GroupEnd("coomer/Download")

	return result

}
