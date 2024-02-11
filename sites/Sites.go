package sites

import "godl/sites/coomer"

var SitesMap map[string]Site = map[string]Site{
	"coomer": {coomer.Can, coomer.Identify, coomer.Index, coomer.Download},
}

