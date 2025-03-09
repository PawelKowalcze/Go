
select {
	case v := <-ch:
		// use v
	default:
		// receiving from ch would block
		// so do something else
}


//Read only channels
func main(){
	ch := make(chan int)
	readCh(ch)
}

func readCh(ch <-chan int) {
	// ch can only be read from
	// in this function
}

//Write only channels
func writeCh(ch chan<- int) {
	// ch can only be written to
	// in this function
}