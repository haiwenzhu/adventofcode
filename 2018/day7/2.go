package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "math"
)

func readInput() (lines [][2]byte, err error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
	for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, [2]byte{line[5:6][0], line[36:37][0]})
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return lines, nil
}

func insert(nodes []byte, node byte) []byte {
    n := len(nodes)
    if n == 0 {
        return []byte{node}
    } else if node < nodes[0] {
        return append([]byte{node}, nodes...)
    } else if node > nodes[n-1] {
        return append(nodes, node)
    } else {
        for i := 0; i < n-1; i += 1 {
            if nodes[i] < node && nodes[i+1] > node {
                return append(nodes[:i+1], append([]byte{node}, nodes[i+1:]...)...)
            }
        }
    }
    return nodes
}

func main() {
    lines, err := readInput()
    workersCount := 5
    stepTimeCost := 60
    if err != nil {
        log.Fatal(err)
    }

    type worker struct {
        finishTime int
        node byte
    }

    pre2next := make(map[byte][]byte)
    next2pre := make(map[byte][]byte)
    workers := make([]worker, workersCount)
    for _, line := range lines {
        pre2next[line[0]] = append(pre2next[line[0]], line[1])
        next2pre[line[1]] = append(next2pre[line[1]], line[0])
    }

    var waitWorkDo []byte
    isWorkDone := make(map[byte]bool)
    isWorkDoing := make(map[byte]bool)
    for node := range pre2next {
        if _, ok := next2pre[node]; !ok {
            waitWorkDo = insert(waitWorkDo, node)
        }
    }

    currentTime := 0
    hasWorkDoing := false
    for ;; {
        if len(waitWorkDo) == 0 && !hasWorkDoing {
            break
        }

        //time passby
        nextTime := math.MaxInt32
        hasIdleWorker := false
        for _, w := range workers {
            if w.finishTime > 0 && w.finishTime < nextTime && w.node > 0 {
                nextTime = w.finishTime
            }
            if w.node == 0 {
                hasIdleWorker = true
            }
        }
        if nextTime != math.MaxInt32 && !(hasIdleWorker && len(waitWorkDo) > 0) {
            //passby time when no idle worker of no work waiting to do
            currentTime = nextTime
        }

        var candidates []byte
        hasWorkDoing = false
        for i, w := range workers {
            if w.finishTime <= currentTime {
                isWorkDone[w.node] = true
                candidates = append(candidates, pre2next[w.node]...)
                workers[i].node = 0
            }
            //dispatch work
            if workers[i].node == 0 && len(waitWorkDo) > 0 {
                workers[i].node = waitWorkDo[0]
                isWorkDoing[workers[i].node] = true
                waitWorkDo = waitWorkDo[1:]
                workers[i].finishTime = currentTime + int(workers[i].node-64) + stepTimeCost
            }
            if workers[i].finishTime > currentTime {
                hasWorkDoing = true
            }
        }

        for ;; {
            if len(candidates) == 0 {
                break;
            }
            candidate := candidates[0]
            candidates = candidates[1:]
            canDoing := true
            for _, preNode := range next2pre[candidate] {
                if _, ok := isWorkDone[preNode]; !ok {
                    canDoing = false
                    candidates = append(candidates, preNode)
                }
            }
            // node can waitWorkDo if all of it's previous node is visited
            if canDoing {
                if _, ok := isWorkDoing[candidate]; !ok {
                    waitWorkDo = insert(waitWorkDo, candidate)
                }
            }
        }
    }

    for _, w := range workers {
        if w.finishTime > 0 && w.finishTime > currentTime {
            currentTime = w.finishTime
        }
    }
    fmt.Println(currentTime)
}
