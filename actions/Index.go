package actions

import "godl/console"
import "godl/structs"
import "godl/sites"

func Index(cache *structs.Cache, base_url string) []string {

	var result []string

	for _, site := range sites.SitesMap {

		if site.Can(base_url) {

			result = site.Index(cache, base_url)

			break

		}

	}

	if len(result) == 0 {
		console.Warn("Cannot find any Site Adapter for URL \"" + base_url + "\"")
	}

	return result

}
