package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		results := strings.Split(scanner.Text(), " ")
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		var a = make([]string, 1)
		for i := range results {
			a = append(a, results[i])
                	fmt.Println(results[i])
                }

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
