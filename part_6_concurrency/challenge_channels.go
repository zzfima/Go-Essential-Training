package part6

//WaitForAllGetContentType wait for all content types using channels
func WaitForAllGetContentType(urls []string) ([]string, error) {
	contentTypes := []string{}
	ch := make(chan string)

	go func() {
		for _, url := range urls {
			c, e := GetContentType(url)
			if e != nil {
				return
			}
			ch <- c
		}
		close(ch)
	}()

	for s := range ch {
		contentTypes = append(contentTypes, s)
	}

	return contentTypes, nil
}
