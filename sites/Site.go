package sites

import "godl/structs"

type Site struct {
	Can      func(string) bool
	Index    func(*structs.Cache, string) []string
	Download func(*structs.Cache, string, []string) bool
}
