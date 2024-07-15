package main

// # Client

// ## Get request

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// resp, err := http.Get("http://go-hacker-code.lab.secself.com/robots.txt")
// 	resp, err := http.Get("https://www.bing.com.cn")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("状态：%d\n", resp.StatusCode)
// 	fmt.Printf("长度：%d\n", resp.ContentLength)
// 	fmt.Printf("内容：\n%s", string(body))
// }

// ## Post request

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	data := url.Values{
// 		"user": {"sechelper"},
// 	}

// 	resp, err := http.PostForm("http://go-hacker-code.lab.secself.com", data)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(resp.StatusCode)
// }

// ### submit json data

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"log"
// )

// type Login struct {
// 	User     string
// 	Password string
// }

// func main() {
// 	login := Login{
// 		"sechelper",
// 		"123456",
// 	}

// 	loginJson, _ := json.Marshal(login)

// 	resp, err := http.Post("http://go-hacker-code.lab.secself.com", "application/json", bytes.NewBuffer(loginJson))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println(resp.StatusCode)
// }

// ## set header

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", "http://go-hacker-code.lab.secself.com", nil)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:108.0) Gecko/20100101 Firefox/108.0")

// 	resp, err := client.Do(req)
// 	if err!= nil {
//     log.Fatal(err)
//   }

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(body))
// 	fmt.Println(resp.StatusCode)
// }

// # Server side

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Secself!")
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
