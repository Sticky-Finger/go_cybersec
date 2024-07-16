package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input() string {
	if len(os.Args) < 2 {
		log.Fatalf("too few arguments")
	}

	return os.Args[1]
}

func read_input_file(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 检查扫描期间是否出错
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	return lines
}

func main() {
	filePath := read_input()
	lines := read_input_file(filePath)
	for _, line := range lines {
		fmt.Println(line)
	}
}
