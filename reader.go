package main

import (
	"bufio"
	"fmt"
	"os"
)

//"strings

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



reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString(' ')
//result := strings.Split(data, ",")


  //  for i := range result {
   //     fmt.Println(result[i])
    //}

//fmt.Println(len(result))
fmt.Println(text)
	
}
