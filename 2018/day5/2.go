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
    var chars [26][]byte
    var units, length [26]int
    var i byte
    bytes := make([]byte, 1024)
    for ;; {
        n, err := f.Read(bytes)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
        for i = 0; i < 26; i += 1 {
            for _, c := range bytes {
                //blow char A
                if c < 65 {
                    break
                }
                if c == 65+i || c == 97+i {
                    continue
                }
                if units[i] > 0 && (c-chars[i][units[i]-1] == 32 || chars[i][units[i]-1]-c == 32)  {
                    units[i] -= 1
                } else {
                    if units[i] < length[i] {
                        chars[i][units[i]] = c
                        units[i] += 1
                    } else {
                        chars[i] = append(chars[i], c)
                        length[i] += 1
                        units[i] += 1
                    }
                }
            }
        }
    }
    minUnit := units[0]
    for _, unit := range units {
        if minUnit > unit {
            minUnit = unit
        }
    }
    fmt.Println(minUnit)
}
