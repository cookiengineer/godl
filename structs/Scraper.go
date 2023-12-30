package structs

import "godl/console"
import "net/http"
import net_url "net/url"
import "io/ioutil"
import "strconv"
import "strings"
import "time"

var CONTENT_TYPES []string = []string{
	"application/gzip",
	"application/json",
	"application/ld+json",
	"application/octet-stream",
	"application/rss+xml",
	"application/x-bzip2",
	"application/x-gzip",
	"application/xml",
	"application/zip",
	"image/jpeg",
	"image/png",
	"text/csv",
	"text/html",
	"text/plain",
	"text/xml",
	"video/mp4",
	"video/x-m4v",
}

type Callback func([]byte)

type ScraperTask struct {
	Url      string
	Callback Callback
}

type Scraper struct {
	Cache     *Cache
	Busy      bool
	Limit     int
	Tasks     []ScraperTask
	Headers   map[string]string
	Throttled bool
}

func processRequests(scraper *Scraper) {

	var filtered []ScraperTask
	var limit int = scraper.Limit

	if scraper.Throttled == true {
		limit = 1
	}

	for t := 0; t < len(scraper.Tasks); t++ {

		if len(filtered) < limit {
			filtered = append(filtered, scraper.Tasks[t])
		} else {
			break
		}

	}

	if len(filtered) > 0 {

		for f := 0; f < len(filtered); f++ {

			var task = filtered[f]

			buffer := scraper.Request(task.Url)

			task.Callback(buffer)

		}

		scraper.Tasks = scraper.Tasks[len(filtered):]

		if len(scraper.Tasks) > 0 {

			if scraper.Throttled == true {

				console.Log(strconv.Itoa(len(scraper.Tasks)) + " Request Tasks left...")

				time.AfterFunc(10*time.Second, func() {
					processRequests(scraper)
				})

			} else {

				time.AfterFunc(1*time.Second, func() {
					processRequests(scraper)
				})

			}

		} else {

			scraper.Busy = false

		}

	}

}

func NewScraper(cache *Cache, headers *map[string]string) Scraper {

	var scraper Scraper

	scraper.Cache = cache
	scraper.Busy = false
	scraper.Limit = 8
	scraper.Tasks = make([]ScraperTask, 0)
	scraper.Throttled = false

	scraper.Headers = map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"Accept-Encoding":           "identity",
		"Cache-Control":             "no-cache",
		"Connection":                "keep-alive",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/119.0",
	}

	if headers != nil {

		for key, val := range *headers {
			scraper.SetHeader(key, val)
		}

	}

	return scraper

}

func (scraper *Scraper) DeferDownload(url string, file string) {

	if scraper.Cache.Exists(file) == false {

		scraper.Tasks = append(scraper.Tasks, ScraperTask{
			Url: url,
			Callback: func(buffer []byte) {

				if len(buffer) > 0 {
					scraper.Cache.Write(file, buffer)
				}

			},
		})

		if scraper.Busy == false {

			scraper.Busy = true

			time.AfterFunc(1*time.Second, func() {
				processRequests(scraper)
			})

		}

	} else {

		console.Log("Skip download of \"" + url + "\"")

	}

}

func (scraper *Scraper) DeferRequest(url string, callback Callback) {

	scraper.Tasks = append(scraper.Tasks, ScraperTask{
		Url:      url,
		Callback: callback,
	})

	if scraper.Busy == false {

		scraper.Busy = true

		time.AfterFunc(1*time.Second, func() {
			processRequests(scraper)
		})

	}

}

func (scraper *Scraper) Download(url string, file string) bool {

	var result bool = false

	if scraper.Cache.Exists(file) == true {

		length := len(scraper.Cache.Read(file))

		if length > 0 {

			size := scraper.Stat(url)

			if size != length {

				buffer := scraper.Request(url)

				if len(buffer) > 0 {
					result = scraper.Cache.Write(file, buffer)
				}

			} else {
				console.Log("Skip \"" + url + "\"")
			}

		} else {

			buffer := scraper.Request(url)

			if len(buffer) > 0 {
				result = scraper.Cache.Write(file, buffer)
			}

		}

	} else {

		buffer := scraper.Request(url)

		if len(buffer) > 0 {
			result = scraper.Cache.Write(file, buffer)
		}

	}

	return result

}

func (scraper *Scraper) Send(url string, parameters map[string]string) []byte {

	var buffer []byte
	var content_type string
	var status_code int

	data := net_url.Values{}

	for key, val := range parameters {
		data.Set(key, val)
	}

	client := &http.Client{}
	client.CloseIdleConnections()

	request, err1 := http.NewRequest("POST", url, strings.NewReader(data.Encode()))

	if err1 == nil {

		for key, val := range scraper.Headers {
			request.Header.Set(key, val)
		}

		response, err2 := client.Do(request)

		if err2 == nil {

			status_code = response.StatusCode

			if status_code == 200 || status_code == 304 {

				if len(response.Header["Content-Type"]) > 0 {
					content_type = response.Header["Content-Type"][0]
				} else {
					content_type = "application/octet-stream"
				}

				var valid bool = false

				for c := 0; c < len(CONTENT_TYPES); c++ {

					if strings.Contains(content_type, CONTENT_TYPES[c]) {
						valid = true
						break
					}

				}

				if valid == true {

					data, err2 := ioutil.ReadAll(response.Body)

					if err2 == nil {
						buffer = data
					}

				}

			}

		}

	}

	if len(buffer) > 0 {

		console.Log("Send \"" + url + "\"")

	} else {

		console.Error("Send \"" + url + "\"")

		if content_type != "" {
			console.Error("Unsupported Content-Type \"" + content_type + "\"")
		}

		if status_code != 0 {
			console.Error("Unsupported Status Code \"" + strconv.Itoa(status_code) + "\"")
		}

	}

	return buffer

}

func (scraper *Scraper) Request(url string) []byte {

	var buffer []byte
	var content_type string
	var status_code int

	client := &http.Client{}
	client.CloseIdleConnections()

	request, err1 := http.NewRequest("GET", url, nil)

	if err1 == nil {

		for key, val := range scraper.Headers {
			request.Header.Set(key, val)
		}

		url_parameters, err := net_url.Parse(url)

		if err == nil {

			hostname := strings.TrimSpace(url_parameters.Hostname())

			if hostname != "" {
				request.Header.Set("Hostname", hostname)
			}

		}

		response, err2 := client.Do(request)

		if err2 == nil {

			status_code = response.StatusCode

			if status_code == 200 || status_code == 304 {

				if len(response.Header["Content-Type"]) > 0 {
					content_type = response.Header["Content-Type"][0]
				} else {
					content_type = "application/octet-stream"
				}

				var valid bool = false

				for c := 0; c < len(CONTENT_TYPES); c++ {

					if strings.Contains(content_type, CONTENT_TYPES[c]) {
						valid = true
						break
					}

				}

				if valid == true {

					data, err2 := ioutil.ReadAll(response.Body)

					if err2 == nil {
						buffer = data
					}

				}

			}

		}

	}

	if len(buffer) > 0 {

		console.Log("Request \"" + url + "\"")

	} else {

		console.Error("Request \"" + url + "\"")

		if content_type != "" {
			console.Error("Unsupported Content-Type \"" + content_type + "\"")
		}

		if status_code != 0 {
			console.Error("Unsupported Status Code \"" + strconv.Itoa(status_code) + "\"")
		}

	}

	return buffer

}

func (scraper *Scraper) Stat(url string) int {

	var result int = 0

	client := &http.Client{}
	client.CloseIdleConnections()

	request, err1 := http.NewRequest("HEAD", url, nil)

	if err1 == nil {

		for key, val := range scraper.Headers {
			request.Header.Set(key, val)
		}

		url_parameters, err := net_url.Parse(url)

		if err == nil {

			hostname := strings.TrimSpace(url_parameters.Hostname())

			if hostname != "" {
				request.Header.Set("Hostname", hostname)
			}

		}

		response, err2 := client.Do(request)

		if err2 == nil {

			status_code := response.StatusCode

			if status_code == 200 || status_code == 304 {

				if len(response.Header["Content-Length"]) > 0 {

					tmp, err3 := strconv.Atoi(response.Header["Content-Length"][0])

					if err3 == nil {
						result = tmp
					}

				}

			}

		}

	}

	if result > 0 {
		console.Log("Stat \"" + url + "\" (" + strconv.Itoa(result) + " Bytes)")
	} else {
		console.Error("Stat \"" + url + "\"")
	}

	return result

}

func (scraper *Scraper) SetHeader(key string, val string) {
	scraper.Headers[key] = val
}

func (scraper *Scraper) SetLimit(value int) {

	if value >= 1 {
		scraper.Limit = value
	}

}

func (scraper *Scraper) SetThrottled(value bool) {
	scraper.Throttled = value
}
