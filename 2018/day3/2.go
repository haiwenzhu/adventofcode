package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func readInput() ([][4]int, error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var lines [][4]int
	for scanner.Scan() {
        lines = append(lines, parseRect(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return lines, nil
}

//rect{left, top, right, bottom}
func parseRect(str string) (rect [4]int) {
    strs := strings.Split(str, " ")
    var err error
    str1 := strings.Split(strings.Trim(strs[2], ":"), ",")
    rect[0], err = strconv.Atoi(str1[0]) 
    rect[1], err = strconv.Atoi(str1[1])
    str2 := strings.Split(strs[3], "x")
    rect[2], err = strconv.Atoi(str2[0])
    rect[3], err = strconv.Atoi(str2[1])
    if err != nil {
        log.Fatal(err)
    }
    rect[2] += rect[0]
    rect[3] += rect[1]
    return rect
}

func intersect(a, b [4]int) (rect [4]int, area int) {
    if a[0] >= b[2] || a[2] <= b[0] || a[1] >= b[3] || a[3] <= b[1] {
        //no intersect
        return rect, 0
    } 
    if a[0] >= b[0] {
        rect[0] = a[0]
    } else {
        rect[0] = b[0]
    }
    if a[1] >= b[1] {
        rect[1] = a[1]
    } else {
        rect[1] = b[1]
    }
    if a[2] <= b[2] {
        rect[2] = a[2]
    } else {
        rect[2] = b[2]
    }
    if a[3] <= b[3] {
        rect[3] = a[3]
    } else {
        rect[3] = b[3]
    }
    area = (rect[2]-rect[0]) * (rect[3]-rect[1])
    return rect, area
}

func main() {
    rects, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    var isIntersect bool
    for i, rect := range rects {
        isIntersect = false
        for j, rect1 := range rects {
            if i != j {
                _, area := intersect(rect, rect1)
                if area > 0 {
                    isIntersect = true
                    break
                }
            }
        }
        if !isIntersect {
            fmt.Println(i+1)
            break
        }
    }
}
