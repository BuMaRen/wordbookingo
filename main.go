package main

import "fmt"

func main() {
	rb := ReadWordBook("words.json")
	if rb == nil {
		fmt.Println("nil table")
		return
	}
	fmt.Println(rb)
}
