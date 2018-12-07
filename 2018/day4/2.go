package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
    "sort"
)

func readInput() ([]string, error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var lines []string
	for scanner.Scan() {
        lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}
    sort.Strings(lines)

    return lines, nil
}

func parseLine(line string) (date string, minute int, action string) {
    tokens := strings.Split(line, " ")
    date = tokens[0][1:]
    minute, _ = strconv.Atoi(tokens[1][3:5])
    switch tokens[2] {
    case "Guard":
        action = tokens[3][1:]
    case "falls":
        action = "sleep"
    case "wakes":
        action = "wake"
    }
    return date, minute, action
}

func sum(arr [60]int) (sum int) {
    for _, v := range arr {
        sum += v
    }
    return sum
}

func main() {
    sleepRecords := make(map[string][60]int)
    lines, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    preDate, preMinute, guard := parseLine(lines[0])
    for _, line := range lines[1:] {
        date, minute, action := parseLine(line)
        if date != preDate {
            preMinute = 0
        }
        switch action {
        case "sleep":
        case "wake":
            sleepRecord := sleepRecords[guard]
            for m := preMinute; m < minute; m += 1 {
                sleepRecord[m] += 1
            }
            sleepRecords[guard] = sleepRecord
        default:
            guard = action
        }
        preMinute = minute
        preDate = date
    }
    var mostSleepGuard string
    var maxMinutes, mostSleepMinute int
    for guard, sleepRecord := range sleepRecords {
        for minute, minutes := range sleepRecord {
            if maxMinutes < minutes {
                maxMinutes = minutes
                mostSleepMinute = minute
                mostSleepGuard = guard
            }
        }
    }
    mostSleepGuardInt, _ := strconv.Atoi(mostSleepGuard)
    fmt.Println(mostSleepGuardInt * mostSleepMinute)
}
