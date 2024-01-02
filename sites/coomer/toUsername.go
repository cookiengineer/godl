package coomer

import "godl/structs"
import "godl/utils"
import "net/url"
import "strings"

func toUsername(link string) string {

	var result string

	base, err := url.Parse(link)

	if err == nil {

		if base.Scheme == "https" && base.Host == "coomer.su" {

			if strings.HasPrefix(base.Path, "/onlyfans/user/") {

				tmp := strings.Split(base.Path[15:], "/")

				if len(tmp) >= 1 {
					result = strings.TrimSpace(tmp[0])
				}

			} else if strings.HasPrefix(base.Path, "/fansly/user/") {

				scraper := structs.NewScraper(nil, &map[string]string{
					"Referer": base.Scheme + "://" + base.Host + base.Path,
				})
				buffer := scraper.Request(base.Scheme + "://" + base.Host + base.Path)
				nodes := utils.Query(buffer, "header.user-header span[itemprop=\"name\"]")

				if len(nodes) == 1 {

					tmp := utils.ToText(nodes[0])

					if tmp != "" {
						result = strings.TrimSpace(tmp)
					}

				}

			}

		}

	}

	return result

}
