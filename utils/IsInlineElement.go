package utils

import "golang.org/x/net/html"

func IsInlineElement(node *html.Node) bool {

	var result bool = false

	if node.Type == html.ElementNode {

		inline_elements := []string{
			"a",
			"abbr",
			"acronym",
			"b",
			"bdo",
			"big",
			"br", // special case, actually block element behaviour
			"button",
			"cite",
			"code",
			"data",
			"del",
			"details",
			"dfn",
			"em",
			"i",
			"img",
			"input",
			"ins",
			"kbd",
			"label",
			"map",
			"object",
			"option",
			"q",
			"ruby",
			"rp",
			"rt",
			"s",
			"samp",
			"script",
			"select",
			"small",
			"span",
			"strong",
			"sub",
			"sup",
			"summary",
			"textarea",
			"time",
			"tt",
			"u",
			"var",
		}

		for i := 0; i < len(inline_elements); i++ {

			if inline_elements[i] == node.Data {
				result = true
				break
			}

		}

	}

	return result

}
