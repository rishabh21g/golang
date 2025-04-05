package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello, World!")
	// result := sum(120, 120)
	var result int = sum(11, 4)
	fmt.Println(result)

}
