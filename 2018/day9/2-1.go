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

func main() {
    playersCount, worth, err := readInput("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    //worth *= 100

    scores := make(map[int]int)
    //double list
    marbles := make([]int, (worth+1)*2)
    marbles[0] = 0
    marbles[1] = 0
    for i := 0; i < playersCount; i += 1 {
        scores[i] = 0
    }
    idx := 0
    player := 0
    for w := 1; w <= worth; w += 1 {
        if w % 23 == 0 {
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            idx = marbles[idx*2]
            scores[player] += w + idx
            marbles[marbles[idx*2]*2+1] = marbles[idx*2+1]
            marbles[marbles[idx*2+1]*2] = marbles[idx*2]
            idx = marbles[idx*2+1]
        } else {
            marbles[w*2] = marbles[idx*2+1]
            marbles[w*2+1] = marbles[marbles[idx*2+1]*2 + 1]
            marbles[marbles[w*2+1]*2] = w
            marbles[marbles[w*2]*2+1] = w
            idx = w
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
