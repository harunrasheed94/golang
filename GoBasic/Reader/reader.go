package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hi this is haroon")
	p := make([]byte, 5)
	for {
		n, err := r.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[0:n])) //should handle any remaining bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(p[0:n]))
	}
}
