package memory

func Fill() [10 * 1024 * 1024]string {
	var buffer [10 * 1024 * 1024]string
	for e, _ := range buffer {
		buffer[e] = "string"
	}
	return buffer
}
