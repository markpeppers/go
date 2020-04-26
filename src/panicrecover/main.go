package main

import "fmt"

func main() {
	fmt.Println(badreturn())
}

func badreturn() (ret string) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			ret = "haha!"
		default:
			panic(p)
		}
	}()

	panic(bailout{})
}
