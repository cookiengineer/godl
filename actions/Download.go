package actions

import "godl/console"
import "godl/structs"
import "godl/sites"

func Download(cache *structs.Cache, base_url string, media_urls []string) bool {

	var result bool = false

	for _, site := range sites.SitesMap {

		if site.Can(base_url) {

			result = site.Download(cache, base_url, media_urls)

			break

		}

	}

	if result == false {
		console.Warn("Cannot find any Site Adapter for URL \"" + base_url + "\"")
	}

	return result

}
