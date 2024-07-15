package main

import "fmt"

func main() {
	// // name of variables
	// var year = 2022
	// var age, name = 10, "小明"
	// fmt.Printf("%d 年%s %d 岁", year, name, age)

	// // # boolean variables
	// var a = true
	// var b = false
	// // ## simplified statements
	// a := true
	// b := false
	// // ## only declare variables
	// var a bool // default value is false
	// a = true // assign value

	// // # number variables
	// var a = 1
	// var b = 65535
	// // ## simplified statements
	// a := 1
	// b := 65535
	// // ## only declare variables
	// var a int // default value is 0
	// a = 1 // assign value

	// // # string variables
	// var a = "助安社区"
	// var b = "君师"
	// // ## simplified statements
	// a := "助安社区"
	// b := "君师"
	// // ## only declare variables
	// var a string // default value is ""
	// a = "助安社区" // assign value

	// // # array variables
	// var a = [9]int{} // fixed length array
	// var b = [4]bool{true, true} // assign values to the head 2 elements, the rest are default value. The value of the array is {true, true, false, false}
	// var c = [...]int{1, 2, 3, 101, 102} // mutable length array
	// ## access arrays
	// var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(a[4]) // print forth element, output 4
	// a[0] = 101

	// // # slice variables
	// var a = make([]string, 3) // initialize slice, length can be set to 0. Slice is scalable
	// var b = []int{1, 2, 3, 4} // initialize and assign
	// // ## simplified statements
	// a := make([]string, 3)
	// b := []int{1, 2, 3, 4}
	// // ## only declare variables
	// var a []string // default value is []
	// // ## append new element to a slice
	// var a = make([]int, 0)
	// a = append(a, 1)
	// a = append(a, 2)
	// // ## copy slice value
	// var a = []int{1, 2, 3}
	// var b = make([]int, 3) // pay attention to the length of b
	// copy(b, a)
	// fmt.Printf("a = %#v\n", a)
	// fmt.Printf("b = %#v\n", b)
	// // ## cut out slice
	// var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(a[0])
	// fmt.Println(a[:3])
	// fmt.Println(a[4:])
	// fmt.Println(a[4:6])
	// // 从0开始，左包含右不包含

	// // # pointer type
	// var a = "助安社区"
	// var b *string
	// b = &a
	// // ## simplified statements
	// a := "助安社区"
	// b := &a
	// fmt.Println(a)
	// fmt.Println(b)
	// // ## what is the value of b?
	// var b = 101
	// var bp = &b
	// *bp = 102
	// fmt.Println(b)

	// // # struct type
	// type people struct {
	// 	name     string
	// 	age      int
	// 	birthday string
	// 	gender   string
	// }

	// // assign value along with definition
	// var a = people{
	// 	name:     "小明",
	// 	age:      10,
	// 	birthday: "1992-05-15",
	// 	gender:   "男",
	// }
	// fmt.Printf("a = %+v\n", a)

	// // define first then assign value
	// var b = people{}
	// b.name = "君诗"
	// b.age = 30
	// b.birthday = "1982-12-25"
	// b.gender = "女"
	// fmt.Printf("b = %+v\n", b)

	// # Map type
	// declare map variable
	var a map[string]string

	// create map, the code above just defined not created
	a = make(map[string]string)

	// assign value to map variable
	a["name"] = "小米"
	a["birthday"] = "10/10"
	a["gender"] = "女"
	fmt.Printf("a = %+v\n", a)

	// declare variable and assign value
	var b = map[string]string{
    "name":     "小雷",
    "birthday": "11/11",
    "gender":   "男",
  }
  fmt.Printf("b = %+v\n", b)
}
