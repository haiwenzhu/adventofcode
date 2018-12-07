package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
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

func countDiff(s1, s2 string) (count int) {
    for i := range s1 {
        if s1[i]  != s2[i] {
            count += 1
            if count > 1 {
                return count
            }
        }
    }
    return count
}

func intersect(s1, s2 string) string {
    var b strings.Builder
    for i, c := range s1 {
        if s1[i] == s2[i] {
            b.WriteRune(c)
        }
    }
    return b.String()
}

func main() {
    strs, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    count := len(strs)
    for i, str := range strs {
        for j := i+1; j < count; j += 1 {
            if countDiff(str, strs[j]) == 1 {
                fmt.Println(intersect(str, strs[j]))
                return;
            }
        }
    }
}
