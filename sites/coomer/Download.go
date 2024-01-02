package coomer

import "godl/console"
import "godl/structs"
import "godl/utils"
import "net/url"
import "strconv"
import "time"

func Download(cache *structs.Cache, link string, downloads []string) bool {

	var result bool = false

	console.Group("coomer/Download")

	base, err := url.Parse(link)

	if err == nil {

		if base.Scheme == "https" && base.Host == "coomer.su" {

			username := toUsername(link)
			scraper := structs.NewScraper(cache, &map[string]string{
				"Referer": base.Scheme + "://" + base.Host + base.Path,
			})

			if username != "" && len(downloads) > 0 {

				for d := 0; d < len(downloads); d++ {

					download := downloads[d]
					filename := utils.ToFilename(download)

					if !cache.Exists("/" + username + "/" + filename) {

						buffer := scraper.Request(download)

						if len(buffer) > 0 {
							cache.Write("/" + username + "/" + filename, buffer)
						}

						time.Sleep(100 * time.Millisecond)

					}

					console.Progress(strconv.Itoa(d+1) + " of " + strconv.Itoa(len(downloads)) + " Downloads")

				}

				result = true

			}

		}

	}

	console.GroupEnd("coomer/Download")

	return result

}
