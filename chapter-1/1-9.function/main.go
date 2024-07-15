package main

import "fmt"

// // # function definition

// func sub(x int, y int) int {
// 	return x + y
// }

// // # function involving

// func main() {
// 	z := sub(10, 5)
// 	fmt.Println(z)
// }

// # function with multiple return values

// func swap(x, y int) (int, int) {
// 	return y, x
// }

// func main() {
// 	x, y := swap(10, 15)
// 	fmt.Println(x, y)
// }

// # function parameters

// ## value delivering

// func autoincrement(x int) int {
// 	x += 2
// 	return x
// }

// func main() {
// 	x := 15
// 	fmt.Println(x)
// 	c := autoincrement(x)
// 	fmt.Println(x, c)
// }

// ## reference delivering

// func increment(x *int) {
// 	*x++
// }

// func main() {
// 	number := 1
// 	increment(&number)
// 	fmt.Println(number)
// }

// # anonymous functions
func main() {
	func() {
		fmt.Println("anonymous function")
	}()
}
