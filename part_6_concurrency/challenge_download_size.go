/* Calculate total download size for NYC taxi data for 2020
For each month, we have two files: green and yellow. For example:
	https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2020-03.csv
	https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2020-03.csv
Turn the below sequential algorithm to a concurrent one using goroutine per file.
*/

package part6

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	urlTemplate = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2020-%02d.parquet"
	colors      = []string{"green", "yellow"}
)

func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

// RunAllDownloadSize run download size on a lot of sites
func RunAllDownloadSize() (int, time.Duration, error) {
	start := time.Now()
	size := 0

	var chNumber = make(chan int)
	var chError = make(chan error)

	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			u := fmt.Sprintf(urlTemplate, color, month)
			fmt.Println(u)

			go downloadSizeChannels(u, chNumber, chError)

			n := <-chNumber
			err := <-chError

			if err != nil {
				log.Fatal(err)
				return 0, 0, err
			}
			size += n
		}
	}

	duration := time.Since(start)
	fmt.Println(size, duration)

	return size, duration, nil
}

func downloadSizeChannels(url string, chNumber chan int, chError chan error) {
	n, e := downloadSize(url)
	chNumber <- n
	chError <- e
}
