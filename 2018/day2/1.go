package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
)

func readInput() ([]string, error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var strs []string
	for scanner.Scan() {
        strs = append(strs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return strs, nil
}

func countRepeat(s string) (count2, count3 int) {
    cmap := make(map[rune]int)
    for _, c := range s {
        if _, ok := cmap[c]; ok {
            cmap[c] += 1
        } else {
            cmap[c] = 1
        }
    }
    for _, count := range cmap {
        if count == 2 {
            count2 = 1
            if count3 > 0 {
                break
            }
        } else if count == 3 {
            count3 = 1
            if count2 > 0 {
                break
            }
        }
    }
    return count2, count3
}

func main() {
    strs, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    var sum2, sum3 int
    for _, str := range strs {
        count2, count3 := countRepeat(str)
        sum2 += count2
        sum3 += count3
    }

    fmt.Println(sum2*sum3)
}
