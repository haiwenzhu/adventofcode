package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "log"
)

func readInput() ([]int, error) {
    f, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(f)
    var nums []int
	for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return nil, err
        }
        nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
        return nil, err
	}

    return nums, nil
}

func main() {
    nums, err := readInput()
    if err != nil {
        log.Fatal(err)
    }
    count := len(nums)
    var sum []int
    i := 0
    for ;; {
        idx := i % count
        for j := i-1; j >= 0; j -= 1 {
            sum[j] += nums[idx]  
            if sum[j] == 0 {
                if j > 0 {
                    fmt.Println(sum[0]+nums[idx])
                } else {
                    fmt.Println(0)
                }
                return
            }
        }
        sum = append(sum, nums[idx])
        i += 1
    }
}
