package main

import "fmt"

func main() {

	count := 5

	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)

	name := "Johannes"
	name += " Scharfenberg"

	fmt.Println("Hello", name)

}
