package xchannels

func nilChannelWork() {
	var x chan int
	x <- 1 //panic

}
