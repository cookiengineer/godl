package utils

import "golang.org/x/net/html"

func ToAttribute(node *html.Node, attribute string) string {

	var result string

	if node.Type == html.ElementNode {

		for a := 0; a < len(node.Attr); a++ {

			tmp := node.Attr[a]

			if tmp.Key == attribute {
				result = tmp.Val
				break
			}

		}

	}

	return result

}
