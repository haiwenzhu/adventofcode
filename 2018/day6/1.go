package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
    "math"
)

func readInput() (points [][2]int, error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var lines []string
	for scanner.Scan() {
        xy := strings.Split(scanner.Text(), ",")
        x, _ := strconv.Atoi(xy[0]) 
        y, _ := strconv.Atoi(xy[1]) 
        points = append(points, [2]int{x, y})
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return points, nil
}

func findInfinitePoints([][2]int) []bool {
}

func main() {
    points, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    var minX, maxX, minY, maxY int
    minX = math.MaxInt32
    minY = math.MaxInt32
    for _, point := range points {
        if minX > point[0] {
            minX = point[0]
        }
        if maxX < point[0] {
            maxX = point[0]
        }
        if minY > point[1] {
            minY = point[1]
        }
        if maxY < point[1] {
            maxY = point[1]
        }
    }

    distance := 1
    visited := make(map[int]map[int]int)
    areas := make([]int, len(points))
    var x, y int
    for ;; {
        for i := 1; i <= distance; i += 1 {
            for j := distance-i; j <= distance; j += 1 {
                for _, point := range points {
                    x = point[0] + i
                    y = point[1] + j
                    if _, ok := visited[x]; !ok {
                        visited[x] = make(map[int]int)
                    }
                    visited[x][y] += 1
                }
            }
        }

        visitFinish = true
        for idx, point := range points {
            for i := 1; i <= distance; i += 1 {
                finish = true
                for j := distance-i; j <= distance; j += 1 {
                    x = point[0] + i
                    y = point[1] + j
                    if visited[x][y] == 1 && areas[idx] >= 0 {
                        if x <= minX || x >= maxX || y <= minY || y >= maxY {
                            areas[idx] = -1
                        } else {
                            finish = false
                            areas[idx] += 1
                        }
                    }
                }
                visitFinish &= finish
            }
        }
        if visitFinish {
            break
        }
    }
    maxArea := 0
    for _, area := range areas {
        if area > 0 && maxArea < area {
            maxArea = area
        }
    }
    fmt.Println(maxArea)
}
