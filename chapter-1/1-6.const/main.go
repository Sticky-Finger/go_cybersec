package main

import "fmt"

// const PI = 3.14

// func main() {
// 	fmt.Println(PI)
// }

// # enum type

// const (
// 	RED    = 1
// 	ORANGE = 2
// 	YELLOW = 3
// 	GREEN  = 4
// 	BLUE   = 5
// 	PURPLE = 6
// )

// func main() {
// 	fmt.Println(RED)
// }

// # iota

const (
	RED    = 1
	ORANGE = 2
	YELLOW = 3
	GREEN  = iota
	BLUE   = 5
	PURPLE = 6
)

func main() {
	fmt.Println(GREEN)
}
