package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func printer() {
	if r%2 == 0 {
		fmt.Println(stars)
	} else {
		fmt.Println(plus)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	plus := strings.Repeat("+", r)
	printer
}
