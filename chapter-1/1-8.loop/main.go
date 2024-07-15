package main

import "fmt"

// # for loop

// ## form 1

// func main() {
// 	for i := 1; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// ## form 2

// func main() {
// 	i:= 0
// 	for i < 10 {
// 		i++
// 		fmt.Println(i)
// 	}
// }

// ## form 3

// func main() {
// 	userPassMap := map[string]string{
// 		"admin":  "admin",
// 		"root":   "123456",
// 		"ubuntu": "12345678",
// 		"test":   "passwd",
// 	}
// 	fmt.Println("用户\t密码")
// 	for user, pass := range userPassMap {
// 		fmt.Printf("%s\t%s\n", user, pass)
// 	}

// 	for user := range userPassMap {
// 		fmt.Printf("%s\t%s\n", user, userPassMap[user])
// 	}

// 	for _, pass := range userPassMap {
// 		fmt.Printf("%s\n", pass)
// 	}
// }

// # for loop embedding

// func main() {
// 	users := []string{
// 		"admin",
// 		"root",
// 		"ubuntu",
// 		"test",
// 	}
// 	passwords := []string{
// 		"admin",
// 		"123456",
// 		"password",
// 		"root",
// 	}
// 	fmt.Println("用户\t密码")
// 	for i := range users {
// 		for x := range passwords {
// 			fmt.Printf("%s\t%s\n", users[i], passwords[x])
// 		}
// 	}
// }

// # loop control

// ## goto statement

func main() {
	fmt.Println("代码片段1")
	goto JUMP
	fmt.Println("代码片段2")

JUMP:
	{
		fmt.Println("代码片段3")
	}
}
