package main

func main() {
	var a = increment()

	println("inc:\tValue Of[", a, "]\tAddr Of[", &a, "]")
}

//go:noinline
func increment() int {
	num := 11
	println("inc:\tValue Of[", num, "]\tAddr Of[", &num, "]")
	//esp(&num)
	//println("inc:\tValue Of[", num, "]\tAddr Of[", &num, "]")

	return num
}

//go:noinline
func esp(nptr *int) *int {
	println("inc:\tValue Of[", *nptr, "]\tAddr Of[", nptr, "]")
	return nptr
}
