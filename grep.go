package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	line int
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println(`Usage:
		go run grep.go <pattern> < text.file`)
		return
	}
	word := os.Args[1]
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		line++

		wl := strings.Fields(in.Text())

		// check if word exists, if so print it

		for _, w := range wl {
			if strings.Contains(w, word) {
				fmt.Printf("#%d: %v\n", line, w)
			}
		}

	}

	if err := in.Err(); err != nil {
		fmt.Println("#### Err:", err)
	}

}
