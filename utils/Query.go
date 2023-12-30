package utils

import "github.com/ericchiang/css"
import "golang.org/x/net/html"
import "strings"

func Query(buffer []byte, query string) []*html.Node {

	var result []*html.Node = make([]*html.Node, 0)

	reader := strings.NewReader(string(buffer))
	document, err1 := html.Parse(reader)
	selector, err2 := css.Parse(query)

	if err1 == nil && err2 == nil {
		result = selector.Select(document)
	}

	return result

}
