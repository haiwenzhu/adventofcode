package main

import (
    "os"
    "fmt"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func readInput(filename string) (input []int, err error) {
    f, err := os.OpenFile(filename, os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var line string
	for scanner.Scan() {
        line = scanner.Text()
        break
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    nums := strings.Split(line, " ")
    for _, num := range nums {
        num, _ := strconv.Atoi(num)
        input = append(input, num)
    }

    return input, nil
}

type node struct {
    parent *node
    children []*node
    meta []int
    childrenCount, metaCount, start, end int
}

func parseTree(input []int) node {
    var root node
    root.parent = nil
    root.childrenCount = input[0]
    root.metaCount = input[1]
    root.children = make([]*node, input[0])
    root.start = 0

    curNode := &root
    for ;; {
        if curNode == nil {
            break
        }
        if curNode.childrenCount > 0 {
            curNodeChanged := false
            for i, child := range curNode.children {
                if child == nil {
                    child = &node{}
                }
                if child.parent == nil {
                    child.parent = curNode
                    if i > 0 {
                        child.start = curNode.children[i-1].end
                    } else {
                        child.start = curNode.start + 2
                    }
                    child.childrenCount = input[child.start]
                    child.metaCount = input[child.start+1]
                    if child.childrenCount > 0 {
                        child.children = make([]*node, child.childrenCount)
                    }
                    curNode.children[i] = child
                    curNode = child
                    curNodeChanged = true
                    break
                }
            }
            if !curNodeChanged {
                curNode.end = curNode.children[curNode.childrenCount-1].end + curNode.metaCount
                curNode.meta = input[curNode.end-curNode.metaCount:curNode.end]
            } else {
                continue
            }
        } else {
            curNode.end = curNode.start + 2 + curNode.metaCount
            curNode.meta = input[curNode.start+2:curNode.end]
        }
        curNode = curNode.parent
    }

    return root
}

func sumMeta(root node) int {
    sum := 0
    for _, v := range root.meta {
        sum += v
    }
    for _, child := range root.children {
        sum += sumMeta(*child)
    }
    return sum
}

func main() {
    input, err := readInput("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    root := parseTree(input)
    fmt.Println(sumMeta(root))
}
