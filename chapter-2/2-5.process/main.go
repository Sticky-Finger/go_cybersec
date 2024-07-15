package main

// # Command function

// import (
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	cmd := exec.Command("ls")
// 	// cmd := exec.Command("dir")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(output))
// }

// ## provide a command and parameters

// import (
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	cmd := exec.Command("ls", "-al", "/tmp")
// 	cmdOut, err := cmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(cmdOut))
// }

// ## use pipe

// import (
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	cmd := exec.Command("ping", "baidu.com")
// 	stdoutPipe, err := cmd.StdoutPipe()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		panic(err)
// 	}

// 	go func() {
// 		for {
// 			buf := make([]byte, 1024)
// 			n, err := stdoutPipe.Read(buf) // 从管道读取执行结果到buf中
// 			if err != nil {
// 				return
// 			}
// 			output := string(buf[:n])
// 			fmt.Print(output)
// 		}
// 	}()

// 	// 待命令执行完成，并检查命令的返回值以判断是否执行成功。
// 	err = cmd.Wait()
// 	if err != nil {
// 		fmt.Printf("Command finished with error: %v\n", err)
// 	} else {
// 		fmt.Println("Command finished successfully.")
// 	}
// }

// # LookPath function

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println("ls not found in PATH")
	} else {
		fmt.Println("ls is available at", path)
	}
}
