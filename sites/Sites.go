package sites

import "godl/sites/coomer"
import "godl/structs"

type Site struct {
	Can      func(string) bool
	Download func(*structs.Cache, string) bool
}

var SitesMap map[string]Site = map[string]Site{
	"coomer": {coomer.Can, coomer.Download},
}

