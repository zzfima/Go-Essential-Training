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

//ChannelMultipleRunWithClose multiple chanel operation with close operation
func ChannelMultipleRunWithClose() int {
	ch := make(chan int)
	sum := 0

	go func() {
		for i := 0; i < 5; i++ {
			ch <- 5
		}
		close(ch) //close channel
	}()

	//iterate on channel until not closed
	for i := range ch {
		sum += i
	}

	return sum
}

//BufferedChannel using channel buffered initialized by 1
func BufferedChannel() int {
	ch := make(chan int, 1)
	ch <- 55
	return <-ch
}
