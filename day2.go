package main

import (
	"bufio"
	"strconv"
	"strings"
	"os"
)

func day2() {
	file, err := os.Open("./inputs/day2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), " ")
		var nums = []int{}
		for _, i := range line {
			j, err := strconv.Atoi(i)
			check(err)
			nums = append(nums, j)
		}
		isSafe := true
		skip := -1
		for {
			x := 0
			isIncreasing := false
			for idx, i := range nums {
				if (skip == idx) {
					continue
				}
				if ((idx == 1 || (skip == 1 && idx == 2)) && i > x) {
					isIncreasing = true
				}
				if ((idx > 0 || (skip == 0 && idx > 1)) && i == x) {
					isSafe = false
					break
				} else if ((idx > 0 || (skip == 0 && idx > 1)) && isIncreasing && i - x > 3) {
					isSafe = false
					break
				} else if ((idx > 0 || (skip == 0 && idx > 1)) && !isIncreasing && x - i > 3) {
					isSafe = false
					break
				}
				if ((idx > 1 || (skip == 1 && idx > 2)) && isIncreasing && i < x) {
					isSafe = false
					break
				} else if ((idx > 1 || (skip == 1 && idx > 2)) && !isIncreasing && i > x) {
					isSafe = false
					break
				}
				x = i
			}
			skip++
			if (isSafe || skip == len(nums)) {
				break
			}
			isSafe = true
		}
		
		if (isSafe) {
			safe++
		}
	}
	check(scanner.Err())

	println(safe)
}