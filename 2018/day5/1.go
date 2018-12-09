package main

import (
    "os"
    "fmt"
    "io"
    "log"
)

func main() {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    if err != nil {
        log.Fatal(err)
    }
    var chars []byte
    var units, length int
    bytes := make([]byte, 1024)
    for ;; {
        n, err := f.Read(bytes)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
        for _, c := range bytes {
            //blow char A
            if c < 65 {
                break
            }
            if units > 0 && (c-chars[units-1] == 32 || chars[units-1]-c == 32)  {
                units -= 1
            } else {
                if units < length {
                    chars[units] = c
                    units += 1
                } else {
                    chars = append(chars, c)
                    length += 1
                    units += 1
                }
            }
        }
    }
    fmt.Println(units)
}
