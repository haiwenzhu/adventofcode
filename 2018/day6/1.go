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

func readInput() (points [][2]int, err error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
	for scanner.Scan() {
        xy := strings.Split(scanner.Text(), ", ")
        x, _ := strconv.Atoi(xy[0]) 
        y, _ := strconv.Atoi(xy[1]) 
        points = append(points, [2]int{x, y})
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return points, nil
}

func getPointsEdge(points [][2]int) (minX, maxX, minY, maxY int) {
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
    return
}

func getNearby(point [2]int, distance int) (points [][2]int) {
    for i := 0; i <= distance; i += 1 {
        j := distance - i
        if i == 0 && j == 0 {
            points = append(points, [2]int{point[0]+i, point[1]+j})
        } else if i == 0 {
            points = append(points, [][2]int{
                [2]int{point[0]+i, point[1]+j},
                [2]int{point[0]+i, point[1]-j},
            }...)
        } else if j == 0 {
            points = append(points, [][2]int{
                [2]int{point[0]+i, point[1]+j},
                [2]int{point[0]-i, point[1]+j},
            }...)
        } else {
            points = append(points, [][2]int{
                [2]int{point[0]+i, point[1]+j},
                [2]int{point[0]-i, point[1]+j},
                [2]int{point[0]+i, point[1]-j},
                [2]int{point[0]-i, point[1]-j},
            }...)
        }
    }
    return points
}

func main() {
    points, err := readInput()
    if err != nil {
        log.Fatal(err)
    }

    minX, maxX, minY, maxY := getPointsEdge(points)

    distance := 0
    visited := make(map[int]map[int]int)
    areas := make([]int, len(points))
    for ;; {
        for _, point := range points {
            nearPoints := getNearby(point, distance)
            for _, p := range nearPoints {
                if _, ok := visited[p[0]]; !ok {
                    visited[p[0]] = make(map[int]int)
                }
                visited[p[0]][p[1]] += 1
            }
        }

        visitFinished := true
        for idx, point := range points {
            finish := true
            nearPoints := getNearby(point, distance)
            for _, p := range nearPoints {
                if visited[p[0]][p[1]] == 1 && areas[idx] >= 0 {
                    //meet edge
                    if p[0] <= minX || p[0] >= maxX || p[1] <= minY || p[1] >= maxY {
                        areas[idx] = -1 //infinity point
                    } else {
                        finish = false
                        areas[idx] += 1
                    }
                }
            }
            visitFinished = visitFinished && finish
        }
        if visitFinished {
            break
        }
        distance += 1
    }

    maxArea := 0
    for _, area := range areas {
        if area > 0 && maxArea < area {
            maxArea = area
        }
    }
    fmt.Println(maxArea)
}
