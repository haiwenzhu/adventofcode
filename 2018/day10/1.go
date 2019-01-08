package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strconv"
    "regexp"
    "math"
)

func readInput(filename string) (inputs [][4]int, err error) {
    f, err := os.OpenFile(filename, os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return inputs, err
    }

    re := regexp.MustCompile("-?[0-9]+")
    scanner := bufio.NewScanner(f)
	for scanner.Scan() {
        line := re.FindAllString(scanner.Text(), -1)
        x, _ := strconv.Atoi(line[0])
        y, _ := strconv.Atoi(line[1])
        vx, _ := strconv.Atoi(line[2])
        vy, _ := strconv.Atoi(line[3])
        inputs = append(inputs, [4]int{x, y, vx, vy})
	}
	if err := scanner.Err(); err != nil {
        return inputs, err
	}
    
    return inputs, nil
}

func draw(points [][4]int, sec int) {
    maxX := math.MinInt32
    maxY := math.MinInt32
    minX := math.MaxInt32
    minY := math.MaxInt32
    x2y := make(map[int]map[int]bool)
    for _, p := range points {
        x := p[0] + sec * p[2]
        y := p[1] + sec * p[3]
        if maxX < x {
            maxX = x
        }
        if minX > x {
            minX = x
        }
        if maxY < y {
            maxY = y
        }
        if minY > y {
            minY = y
        }
        if _, ok := x2y[y]; !ok {
            x2y[y] = make(map[int]bool)
        }
        x2y[y][x] = true
    }
    for y := minY; y <= maxY; y++ {
        if _, ok := x2y[y]; !ok {
            return
        }
    }
    for y := minY; y <= maxY; y++ {
        for x := minX; x <= maxX; x++ {
            if _, ok := x2y[y][x]; !ok {
                fmt.Print(".")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Println("")
    }
}

//getMaxYPoint return the point with max y value and velocity
func getMaxYPoint(points [][4]int) (point [4]int) {
    maxY := math.MinInt32
    maxVY := math.MinInt32
    for _, p := range points {
        if maxVY < p[3] {
            maxVY = p[3]
            maxY = p[1]
            point = p
        } else if maxVY == p[3] {
            if maxY < p[1] {
                maxY = p[1]
                point = p
            }
        }
    }
    return point
}

func main() {
    points, err := readInput("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    count := len(points)

    loop := 0
    maxYPoint := getMaxYPoint(points)
    for ;; {
        maxY := math.MinInt32
        minY := math.MaxInt32
        loop++
        for _, p := range points {
            y := p[1] + loop * p[3]
            if maxY < y {
                maxY = y
            }
            if minY > y {
                minY = y
            }
        }
        if maxY - minY + 1 <= count {
            draw(points, loop)
        }
        //if the point with max y axis value and velocity became maxY, it's time to break
        if maxYPoint[1] + loop*maxYPoint[3] == maxY {
            break;
        }
    }
}
