package sites

import "godl/structs"

type Site struct {
	Can      func(string) bool
	Identify func(string) string
	Index    func(*structs.Cache, *structs.Index, string) bool
	Download func(*structs.Cache, *structs.Index, string) bool
}
