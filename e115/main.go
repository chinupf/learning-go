package main

import "fmt"

func add5value(count int) {
	count += 5
	fmt.Println("add5Value:", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point :", *count)
}

func main() {

	var count int

	add5value(count)
	fmt.Println("add5Value post:", count)

	add5Point(&count)
	fmt.Println("add5Point post:", count)

}
