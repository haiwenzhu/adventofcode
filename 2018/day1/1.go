package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "log"
)

func main() {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(f)
    sum := 0 
	for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        sum += num
	}
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
    fmt.Println(sum)
}
