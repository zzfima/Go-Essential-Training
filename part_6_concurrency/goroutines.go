package part6

import (
	"fmt"
	"net/http"
	"sync"
)

//goroutine - like thread but very light

// GetContentType get content type of single url
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

// SerialGetContentType get content type of multiple urls
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

// SerialGetContentTypeConcurrent get content type of multiple urls in asynchronous manner
func SerialGetContentTypeConcurrent(sites []string) ([]string, error) {
	contentTypes := []string{}
	var wg sync.WaitGroup //wait group - like semaphore

	for _, url := range sites {
		wg.Add(1) //*** enter to semaphore (add 1 to wg)
		go func(u string) {
			c, e := GetContentType(u)
			if e != nil {
				return
			}
			contentTypes = append(contentTypes, c)
			wg.Done() //*** out from semaphore (sub 1 from wg)
		}(url)
	}
	wg.Wait() //*** wait until all finished ()until wg = 0)
	return contentTypes, nil
}
