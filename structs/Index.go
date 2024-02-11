package structs

import "encoding/json"
import "os"
import "strings"
import "time"

type Index struct {
	Folder    string            `json:"folder"`
	Completed bool              `json:"completed"`
	Downloads map[string]string `json:"downloads"`
}

func NewIndex(folder string) Index {

	var index Index

	index.Downloads = make(map[string]string)

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	stat, err1 := os.Stat(folder)

	if err1 == nil && stat.IsDir() {

		index.Folder = folder

		buffer, err2 := os.ReadFile(index.Folder + "/godl.json")

		if err2 == nil {
			json.Unmarshal(buffer, &index)
			// In case the folder has been renamed
			index.Folder = folder
		}

	} else {

		err2 := os.MkdirAll(folder, 0750)

		if err2 == nil {
			index.Folder = folder
		}

	}

	return index

}

func (index *Index) Exists(url string) bool {

	_, ok := index.Downloads[url]

	if ok == true {

		if index.Downloads[url] != "" {
			return true
		}

	}

	return false

}

func (index *Index) Create(url string) {

	_, ok := index.Downloads[url]

	if ok == false {
		index.Downloads[url] = ""
	}

}

func (index *Index) Set(url string) {

	_, ok := index.Downloads[url]

	if ok == true {
		index.Downloads[url] = time.Now().Format(time.RFC3339)
	} else {
		index.Downloads[url] = time.Now().Format(time.RFC3339)
	}

}

func (index *Index) Write() bool {

	var result bool = false

	buffer, err1 := json.MarshalIndent(index, "", "\t")

	if err1 == nil {

		err2 := os.WriteFile(index.Folder + "/godl.json", buffer, 0666)

		if err2 == nil {
			result = true
		}

	}

	return result

}
