package coomer

import "net/url"

func Can(link string) bool {

	var result bool = false

	tmp, err := url.Parse(link)

	if err == nil {

		if tmp.Scheme == "https" && tmp.Host == "coomer.su" {
			result = true
		}

	}

	return result

}
