package coomer

import "godl/console"
import "godl/structs"
import "godl/utils"
import "net/url"
import "strconv"
import "strings"
import "time"

func Download(cache *structs.Cache, link string) bool {

	var result bool = false

	console.Group("coomer/Download")

	base, err := url.Parse(link)

	if err == nil {

		if base.Scheme == "https" && base.Host == "coomer.su" {

			username := toUsername(link)
			scraper := structs.NewScraper(cache, &map[string]string{
				"Referer": base.Scheme + "://" + base.Host + base.Path,
			})

			if username != "" {

				var downloads []string = make([]string, 0)
				var pages []string = make([]string, 0)
				var amount int = 0

				buffer := scraper.Request(base.Scheme + "://" + base.Host + base.Path)
				nodes := utils.Query(buffer, "div#paginator-top menu a:last-of-type")

				if len(nodes) == 1 {

					link := utils.ToAttribute(nodes[0], "href")

					if strings.Contains(link, "?o=") {
						amount = utils.ToInt(strings.TrimSpace(link[strings.Index(link, "?o=")+3:]))
					}

				}

				if amount > 0 {

					for offset := 0; offset <= amount; offset += 50 {

						if offset == 0 {
							pages = append(pages, base.Scheme + "://" + base.Host + base.Path)
						} else {
							pages = append(pages, base.Scheme + "://" + base.Host + base.Path + "?o=" + strconv.Itoa(offset))
						}

					}

				} else {
					pages = append(pages, base.Scheme + "://" + base.Host + base.Path)
				}

				if len(pages) > 0 {

					for p := 0; p < len(pages); p++ {

						buffer := scraper.Request(pages[p])
						nodes := utils.Query(buffer, "div.card-list__items article.post-card > a")

						if len(nodes) > 0 {

							for n := 0; n < len(nodes); n++ {

								link := utils.ToAttribute(nodes[n], "href")

								if strings.HasPrefix(link, "/") {

									post_buffer := scraper.Request(base.Scheme + "://" + base.Host + link)
									post_images := utils.Query(post_buffer, "div.post__files div.post__thumbnail figure a.fileThumb")
									post_attachments := utils.Query(post_buffer, "ul.post__attachments a.post__attachment-link")

									if len(post_images) > 0 {

										for i := 0; i < len(post_images); i++ {

											image_link := utils.ToAttribute(post_images[i], "href")

											if strings.Contains(image_link, "/data/") {
												downloads = append(downloads, image_link)
											}

										}

									}

									if len(post_attachments) > 0 {

										for a := 0; a < len(post_attachments); a++ {

											attachment_link := utils.ToAttribute(post_attachments[a], "href")

											if strings.Contains(attachment_link, "/data/") {
												downloads = append(downloads, attachment_link)
											}

										}

									}

									time.Sleep(100 * time.Millisecond)

								}

							}

						}

						time.Sleep(100 * time.Millisecond)

					}

				}

				if len(downloads) > 0 {

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

	}

	console.GroupEnd("coomer/Download")

	return result

}
