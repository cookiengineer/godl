package coomer

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

			}

		}

	}

	return result

}
