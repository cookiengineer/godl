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

			username := actions.Identify(url)

			if username == "" {
				console.Error("No Username!")
				os.Exit(1)
				// TODO: Generate random username from url hash?
			}

			cache := structs.NewCache(cwd+"/"+username)
			index := structs.NewIndex(cwd+"/"+username)

			if !index.Completed {
				actions.Index(&cache, &index, url)
			}

			if len(index.Downloads) > 0 {
				actions.Download(&cache, &index, url)
			}

		}

	}

}
