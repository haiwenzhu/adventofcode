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

func abs(x int) int {
    if x > 0 {
        return x
    } else {
        return -x
    }
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

func distanceOf2Point(a, b [2]int) int {
    return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func distanceOfPoints(a [2]int, b [][2]int) int {
    distance := 0
    for _, point := range b {
        distance += distanceOf2Point(a, point)
    }
    return distance
}

func distanceOfPointsX(x int, b [][2]int) int {
    distance := 0
    for _, point := range b {
        distance += abs(point[0] - x)
    }
    return distance
}

func distanceOfPointsY(y int, b [][2]int) int {
    distance := 0
    for _, point := range b {
        distance += abs(point[1] - y)
    }
    return distance
}

func main() {
    points, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    countOfPoints := len(points)
    distance := 10000

    minX, maxX, minY, maxY := getPointsEdge(points)
    distanceX := make(map[int]int)
    distanceY := make(map[int]int)
    minDistanceX := math.MaxInt32
    minDistanceY := math.MaxInt32
    for x := minX; x <= maxX; x += 1 {
        distanceX[x] = distanceOfPointsX(x, points)
        if minDistanceX > distanceX[x] {
            minDistanceX = distanceX[x]
        }
    }
    for y := minY; y <= maxY; y += 1 {
        distanceY[y] = distanceOfPointsY(y, points)
        if minDistanceY > distanceY[y] {
            minDistanceY = distanceY[y]
        }
    }

    var left, right, top, bottom int
    if distance-distanceX[minX]-minDistanceY > 0 {
        left = minX - (distance-distanceX[minX]-minDistanceY)/countOfPoints
        for x := left; x < minX; x += 1 {
            distanceX[x] = distanceX[minX] + (minX-x)*countOfPoints
        }
    } else {
        left = minX
    }
    if distance-distanceX[maxX]-minDistanceY > 0 {
        right = maxX + (distance-distanceX[maxX]-minDistanceY)/countOfPoints
        for x := right; x > maxX; x -= 1 {
            distanceX[x] = distanceX[maxX] + (x-maxX)*countOfPoints
        }
    } else {
        right = maxX
    }
    if distance-distanceY[minY]-minDistanceX > 0 {
        top = minY - (distance-distanceX[minY]-minDistanceX)/countOfPoints
        for y := top; y < minY; y += 1 {
            distanceY[y] = distanceY[minY] + (minY-y)*countOfPoints
        }
    } else {
        top = minY
    }
    if distance-distanceY[maxY]-minDistanceX > 0 {
        bottom = maxY + (distance-distanceY[maxY]-minDistanceX)/countOfPoints
        for y := bottom; y > maxY; y -= 1 {
            distanceY[y] = distanceY[maxY] + (y-maxY)*countOfPoints
        }
    } else {
        bottom = maxY
    }

    area := 0
    for x := left; x <= right; x += 1 {
        for y := top; y <= bottom; y += 1 {
            if distanceX[x] + distanceY[y] < distance {
                area += 1
            }
        }
    }
    fmt.Println(area)
}
