package part6

import (
	"fmt"
	"net/http"
)

//goroutine - like thread but very light

// GetContentType GetContentType
func GetContentType(url string) (string, error) {
	response, e := http.Get(url)

	if e != nil {
		fmt.Printf("error: %s\n", e)
		return "", e
	}

	defer response.Body.Close()
	contentType := response.Header.Get("content-type")
	fmt.Printf("%s->%s\n", url, contentType)
	return contentType, nil
}

// SerialGetContentType SerialGetContentType
func SerialGetContentType(sites []string) ([]string, error) {
	contentTypes := []string{}
	for _, url := range sites {
		c, e := GetContentType(url)
		if e != nil {
			return nil, e
		}
		contentTypes = append(contentTypes, c)
	}

	return contentTypes, nil
}
