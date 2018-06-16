package main

import (
        "fmt"
        "strings"
	"encoding/json"
)


func main() {
    var lines []string
    for {
        var line string
        fmt.Scanln(&line)
        if line == "" {
            break
        }
        lines = append(lines, "["+line+"]")
    }
    all := "[" + strings.Join(lines, ",") + "]"
    inputs := [][]float64{}
    if err := json.Unmarshal([]byte(all), &inputs); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(inputs)
}
