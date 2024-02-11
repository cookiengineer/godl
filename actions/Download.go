package actions

import "godl/console"
import "godl/structs"
import "godl/sites"

func Download(cache *structs.Cache, index *structs.Index, base_url string) bool {

	var result bool = false
	var found bool = false

	for _, site := range sites.SitesMap {

		if site.Can(base_url) {

			result = site.Download(cache, index, base_url)
			found = true
			break

		}

	}

	if found == false {
		console.Warn("Cannot find any Site Adapter for URL \"" + base_url + "\"")
	}

	return result

}
