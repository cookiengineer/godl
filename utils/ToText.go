package utils

import "golang.org/x/net/html"
import "strings"

func ToText(node *html.Node) string {

	var result string

	if node != nil {

		if node.Type == html.DocumentNode {

			for child := node.FirstChild; child != nil; child = child.NextSibling {

				var chunk = ToText(child)

				if chunk != "" {
					result = result + chunk
				}

			}

		} else if node.Type == html.DoctypeNode {

			result = ""

		} else if node.Type == html.ElementNode && node.Data == "pre" {

			for child := node.FirstChild; child != nil; child = child.NextSibling {

				var chunk = ToText(child)

				if chunk != "" {
					result = result + "\n" + chunk
				}

			}

		} else if node.Type == html.ElementNode {

			if node.Data == "br" {

				result = "\n"

			} else {

				for child := node.FirstChild; child != nil; child = child.NextSibling {

					var chunk = ToText(child)

					if chunk != "" {
						result = result + chunk
					}

				}

			}

		} else if node.Type == html.TextNode {

			var chunk = strings.TrimSpace(node.Data)

			if chunk != "" {

				if result != "" {

					if IsBlockElement(node.Parent) {
						result = result + "\n" + chunk
					} else if IsInlineElement(node.Parent) {
						result = result + " " + chunk
					} else {
						result = result + " " + chunk
					}

				} else {

					result = chunk

				}

			}

		}

	}

	return result

}
