package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EntryText struct {
	command []string
	wordCount int
}
	

func main() {
        fmt.Print(`Which example do you want to run?
  1) fmt.Scan(...)
  2) bufio.Reader.ReadString(...)
  3) bufio.Reader.ReadByte(...)
  4) bufio.Reader.ReadRune()
  5) list(...)
Please enter 1..5 and press ENTER: `)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//entryText := scanner.Text()
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		//capturedText := scanner.Text()
		commandList := strings.Split(scanner.Text(), " ")
		//s := make([]string, 1)
		
		for i := range commandList {
			fmt.PrintLn(commandList[i])
		}
		fmt.Println(len(commandList))
		//fmt.PrintLn(capturedText)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
