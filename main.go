package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Dest file required")
        os.Exit(1)
    }

    fmt.Println("Enter data:")
    scanner := bufio.NewScanner(os.Stdin)

    f, err := os.Create(os.Args[1])
    if err != nil {
        panic(err)
    }

    for {
        scanner.Scan()
        line := scanner.Bytes()

        _, err := f.Write(line)
        if err != nil {
            panic(err)
        }
    }
}
