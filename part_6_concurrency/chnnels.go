package part6

//ChannelSingleRun single chanel operation
func ChannelSingleRun() int {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	return <-ch
}

//ChannelMultipleRun multiple chanel operation
func ChannelMultipleRun() int {
	ch := make(chan int)
	count := 3
	sum := 0

	go func() {
		for i := 0; i < count; i++ {
			ch <- 5
		}
	}()

	for i := 0; i < count; i++ {
		sum += <-ch
	}

	return sum
}
