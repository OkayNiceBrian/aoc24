package main

import (
	"bufio"
	"sort"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./inputs/day1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var in1 []int
	var in2 []int
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), "   ")
		int1, err := strconv.Atoi(line[0])
		check(err)
		int2, err := strconv.Atoi(line[1])
		check(err)
		in1 = append(in1, int1)
		in2 = append(in2, int2)
	}
	check(scanner.Err())

	sort.Slice(in1, func(i, j int) bool {
		return in1[i] < in1[j]
	})
	sort.Slice(in2, func(i, j int) bool {
		return in2[i] < in2[j]
	})

	var distance int
	for i := range in1 {
		d := in1[i] - in2[i]
		if (d < 0) {
			d *= -1
		}
		distance += d
	}

	fmt.Println(distance) // Answer 1

	sim := map[int]int{}
	var score int
	for i := range in1 {
		v := in1[i]
		if value, ok := sim[v]; ok {
			score += value
		} else {
			count := 0
			for j := range in2 {
				if in2[j] == v {
					count++
				}
			}
			sim[v] = v * count
			score += v * count
		}
	}

	fmt.Println(score) // Answer 2
}