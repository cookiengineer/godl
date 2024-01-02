package utils

import "golang.org/x/net/html"

func IsBlockElement(node *html.Node) bool {

	var result bool = false

	if node.Type == html.ElementNode {

		block_elements := []string{
			"address",
			"article",
			"aside",
			"blockquote",
			"canvas",
			"dd",
			"div",
			"dl",
			"dt",
			"fieldset",
			"figcaption",
			"figure",
			"footer",
			"form",
			"h1", "h2", "h3", "h4", "h5", "h6",
			"header",
			"hr",
			"li",
			"main",
			"nav",
			"noscript",
			"ol",
			"output",
			"p",
			"pre",
			"section",
			"thead",
			"tfoot",
			"ul",
			"video",
		}

		for b := 0; b < len(block_elements); b++ {

			if block_elements[b] == node.Data {
				result = true
				break
			}

		}

	}

	return result

}
