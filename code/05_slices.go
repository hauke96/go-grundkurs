func initSha1(a []byte) {
	a[0] = 0x01
	a[1] = 0x23
	a[2] = 0x45
	// ...
	fmt.Println(a)
}

func main() {
	var x [20]byte
	initSha1(x[:]) // x[:] is a slice (view) into x
	fmt.Println(x) // x[:] is short-hand for x[0:20]
}
