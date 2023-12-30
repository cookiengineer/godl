package actions

import "godl/console"
import "godl/structs"
import "godl/sites"

func Download(cache *structs.Cache, url string) bool {

	var result bool = false

	for _, site := range sites.SitesMap {

		if site.Can(url) {

			result = site.Download(cache, url)

			break

		}

	}

	if result == false {
		console.Warn("Cannot find any Site Adapter for URL \"" + url + "\"")
	}

	return result

}
