package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
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
    if err != nil {
        log.Fatal(err)
    }

    pre2next := make(map[byte][]byte)
    next2pre := make(map[byte][]byte)

    for _, line := range lines {
        pre2next[line[0]] = append(pre2next[line[0]], line[1])
        next2pre[line[1]] = append(next2pre[line[1]], line[0])
    }

    var visited, visiting []byte
    isVisited := make(map[byte]bool)
    for node := range pre2next {
        if _, ok := next2pre[node]; !ok {
            visiting = insert(visiting, node)
        }
    }

    for ;; {
        if len(visiting) == 0 {
            break
        }
        node := visiting[0]
        visiting = visiting[1:]
        visited = append(visited, node)
        isVisited[node] = true
        candidates := pre2next[node]
        for ;; {
            if len(candidates) == 0 {
                break;
            }
            candidate := candidates[0]
            candidates = candidates[1:]
            canVisiting := true
            for _, preNode := range next2pre[candidate] {
                if _, ok := isVisited[preNode]; !ok {
                    canVisiting = false
                    candidates = append(candidates, preNode)
                }
            }
            // node can visiting if all of it's previous node is visited
            if canVisiting {
                visiting = insert(visiting, candidate)
            }
        }
    }

    var b strings.Builder
    _, err = b.Write(visited)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(b.String())
}
