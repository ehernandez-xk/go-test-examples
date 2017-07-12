package main

func facto(n int) (x int) {
	for i := n; i > 0; i-- {
		x += i
	}
	return
}

func main() {
	println("factorial of 8 is", facto(8))
}
