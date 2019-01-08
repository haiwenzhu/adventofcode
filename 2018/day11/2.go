package main

import (
    "fmt"
)

func calcPowerLevel(x, y, serial int) int {
    rackId := x + 10
    return (rackId * (rackId*y + serial) % 1000) / 100 - 5
}

func max(grid [301][301]int, size int) (x, y, max int) {
    var power [301][301]int
    power[1][1] = 0
    for x := 1; x <= size; x++ {
        for y :=1; y <= size; y++ {
            power[1][1] += grid[x][y]
        }
    }
    maxPower := power[1][1]
    maxX := 1
    maxY := 1
    for x := 2; x <= 301 - size; x++ {
        power[1][x] = power[1][x-1]
        for y := 1; y <= size; y++ {
            power[1][x] += grid[y][x+size-1] - grid[y][x-1]
        }
        if maxPower < power[1][x] {
            maxPower = power[1][x]
            maxX = x
        }
    }
    for y :=2; y <= 301-size; y++ {
        for x := 1; x <= 301-size; x++ {
            power[y][x] = power[y-1][x]
            for s := 0; s < size; s++ {
                power[y][x] += grid[y+size-1][x+s] - grid[y-1][x+s]
            }
            if maxPower < power[y][x] {
                maxPower = power[y][x]
                maxX = x
                maxY = y
            }
        }
    }
    return maxX, maxY, maxPower
}

func main() {
    var grid [301][301]int
    serial := 4151
    for y :=1; y <= 300; y++ {
        for x := 1; x <= 300; x++ {
            grid[y][x] = calcPowerLevel(x, y, serial)
        }
    }
    var maxX, maxY, maxPower, maxSize int
    for size := 1; size <= 300; size++ {
        x, y, power := max(grid, size)
        if maxPower < power {
            maxPower = power
            maxX = x
            maxY = y
            maxSize = size
        }
    }
    fmt.Println(maxX, maxY, maxSize)
}
