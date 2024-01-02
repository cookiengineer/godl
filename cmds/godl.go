package main

import "godl/console"
import "godl/actions"
import "godl/structs"
import "os"
import "strings"

func main() {

	var url string

	cwd, err := os.Getwd()

	if err == nil {

		cache := structs.NewCache(cwd)

		if len(os.Args) == 2 {

			value := os.Args[1]

			if strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "http://") {
				url = value
			}

		}

		console.Clear()
		console.Group("godl: Command-Line Arguments")
		console.Inspect(struct {
			URL string
		}{
			URL: url,
		})
		console.GroupEnd("")

		if url != "" {

			urls := actions.Index(&cache, url)

			if len(urls) > 0 {
				actions.Download(&cache, url, urls)
			}

		}

	}

}
