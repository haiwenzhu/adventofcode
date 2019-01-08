package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func readInput(filename string) (players, worth int, err error) {
    f, err := os.OpenFile(filename, os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return 0, 0, err
    }

    scanner := bufio.NewScanner(f)
    var line string
	for scanner.Scan() {
        line = scanner.Text()
        break
	}
	if err := scanner.Err(); err != nil {
        return 0, 0, err
	}

    nums := strings.Split(line, " ")
    players, _ = strconv.Atoi(nums[0])
    worth, _ = strconv.Atoi(nums[6])

    return players, worth, nil
}

func insert(lists []int, index int, val int) []int {
    return append(lists[:index], append([]int{val}, lists[index:]...)...)
}

func del(lists []int, index int) []int {
    return append(lists[:index], lists[index+1:]...)
}

func main() {
    playersCount, worth, err := readInput("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scores := make(map[int]int)
    marbles := []int{0}
    for i := 0; i < playersCount; i += 1 {
        scores[i] = 0
    }
    index := 0
    player := 0
    length := 1
    for number := 1; number <= worth; number += 1 {
        if number % 23 == 0 {
            index = (index - 7 + length) % length
            scores[player] += number + marbles[index]
            marbles = del(marbles, index)
            length -= 1
        } else {
            index = (index+2) % length
            if index == 0 {
                index = length
            }
            marbles = insert(marbles, index, number)
            length += 1
        }
        player = (player+1) % playersCount
    }
    maxScore := 0
    for _, score := range scores {
        if maxScore < score {
            maxScore = score
        }
    }
    fmt.Println(maxScore)
}
