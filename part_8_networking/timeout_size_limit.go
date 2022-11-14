package part8

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// TimeoutAndSizeLimit set timeout and size limit
func TimeoutAndSizeLimit() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*3000)
	defer cancel()

	req, e := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if e != nil {
		log.Fatal(e)
	}

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()

	const bytesToRead = 14
	r := io.LimitReader(resp.Body, bytesToRead)

	os.Remove("tmp1.txt")

	f, e := os.OpenFile("tmp1.txt", os.O_CREATE, os.ModeAppend)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	io.Copy(f, r)
}
