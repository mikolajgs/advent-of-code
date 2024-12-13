package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")

	sum := getSum(string(f))
	log.Printf("%d\n", sum)

	// part 2
	dontsArr := strings.Split(string(f), "don't()")
	log.Printf("%d\n", len(dontsArr))
	sum2 := 0
	for j, v := range dontsArr {
		if j == 0 {
			sum2 += getSum(v)
			continue
		}
		dosArr := strings.Split(v, "do()")
		if len(dosArr) > 1 {
			left := ""
			for i:=1; i<=len(dosArr)-1; i++ {
				left += dosArr[i]
			}
			sum2 += getSum(left)
		}
	}
	log.Printf("%d\n", sum2)
}

func getSum(f string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	found := re.FindAllString(string(f), -1)
	
	sum := 0
	for _, v := range found {
		s := strings.Replace(v, "mul(", "", 1)
		s = strings.Replace(s, ")", "", 1)
		vals := strings.Split(s, ",")
		left, _ := strconv.Atoi(vals[0])
		right, _ := strconv.Atoi(vals[1])

		m := left * right
		sum += m
	}
	return sum
}
