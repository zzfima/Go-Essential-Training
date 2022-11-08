package part6

import (
	"fmt"
	"net/http"
)

//goroutine - like thread but very light

// GetContentTypeType GetContentTypeType
func GetContentTypeType(url string) (string, error) {
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
