package utils

import "strings"

func ToFilename(value string) string {

	var result string

	tmp := strings.Split(value, "/")

	if len(tmp) > 0 {

		filename := tmp[len(tmp)-1]

		if strings.Contains(filename, "?") {
			filename = filename[0:strings.Index(filename, "?")]
		}

		if strings.Contains(filename, ".") {
			result = filename
		}

	}

	return result

}
