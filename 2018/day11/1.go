package main

import (
    "fmt"
)

func calcPowerLevel(x, y, serial int) int {
    rackId := x + 10
    return (rackId * (rackId*y + serial) % 1000) / 100 - 5
}

func main() {
    var grid [301][301]int
    var power [301][301]int
    serial := 4151
    for y :=1; y <= 300; y++ {
        for x := 1; x <= 300; x++ {
            grid[y][x] = calcPowerLevel(x, y, serial)
        }
    }
    power[1][1] = grid[1][1] + grid[1][2] + grid[1][3] + grid[2][1] + grid[2][2] + grid[2][3] + grid[3][1] + grid[3][2] + grid[3][3]
    maxPower := power[1][1]
    maxX := 1
    maxY := 1
    for x := 2; x <= 298; x++ {
        power[1][x] = power[1][x-1] + grid[1][x+2] + grid[2][x+2] + grid[3][x+2] - grid[1][x-1] - grid[2][x-1] - grid[3][x-1]
        if maxPower < power[1][x] {
            maxPower = power[1][x]
            maxX = x
        }
    }
    for y :=2; y <= 298; y++ {
        for x := 1; x <= 298; x++ {
            power[y][x] = power[y-1][x] + grid[y+2][x] + grid[y+2][x+1] + grid[y+2][x+2] - grid[y-1][x] - grid[y-1][x+1] - grid[y-1][x+2]
            if maxPower < power[y][x] {
                maxPower = power[y][x]
                maxX = x
                maxY = y
            }
        }
    }
    fmt.Println(maxX, maxY)
}
