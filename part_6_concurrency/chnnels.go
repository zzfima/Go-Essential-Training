package part6

//ChannelRun ChannelRun
func ChannelRun() int {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	return <-ch
}
