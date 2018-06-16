package main

import "fmt"
import "time"

func main() {
	ch := make(chan string)
	go func(ch chan string) {
		/* Uncomment this block to actually read from stdin
		reader := bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')
			if err != nil { // Maybe log non io.EOF errors, if you want
				close(ch)
				return
			}
			ch <- s
		}
		*/
		// Simulating stdin
		ch <- "A line of text"
		close(ch)
	}(ch)

stdinloop:
	for {
		select {
		case stdin, ok := <-ch:
			if !ok {
				break stdinloop
			} else {
				fmt.Println("Read input from stdin:", stdin)
			}
		case <-time.After(1 * time.Second):
			// Do something when there is nothing read from stdin
		}
	}
	fmt.Println("Done, stdin must be closed")
}
