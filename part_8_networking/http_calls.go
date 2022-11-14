package part8

import (
	"net/http"
	"strconv"
)

// GetHTTPContentLength get body of "https://httpbin.org/get"
func GetHTTPContentLength() (int, error) {
	response, e := http.Get("https://httpbin.org/get")
	if e != nil {
		return 0, e
	}
	defer response.Body.Close()

	return strconv.Atoi(response.Header["Content-Length"][0])
}
