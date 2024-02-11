package actions

import "godl/console"
import "godl/sites"

func Identify(base_url string) string {

	var result string

	for _, site := range sites.SitesMap {

		if site.Can(base_url) {

			result = site.Identify(base_url)

			break

		}

	}

	if result == "" {
		console.Warn("Cannot find any Site Adapter for URL \"" + base_url + "\"")
	}

	return result

}
