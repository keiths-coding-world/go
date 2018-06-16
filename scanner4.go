package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type userCommand struct{
	commandType string
	commandParams []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		results := strings.Split(scanner.Text(), " ")
		uiCommand := userCommand{}
		uiCommandCount := len(results)
		fmt.Println(uiCommandCount)
		uiParams := make([]string, uiCommandCount)
		for i := range results {
			if i == 0{
				uiCommand.commandType = results[i]
			} else {
				uiParams[i] = results[i]	
			}
                }

		fmt.Println(uiCommand.commandType)

		if uiCommandCount > 0{
			uiCommand.commandParams = uiParams
			for i := range uiParams {
				fmt.Println(uiParams[i])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
