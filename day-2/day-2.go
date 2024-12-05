package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	safeReports := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		safeReport, _ := check(numbers)
		if safeReport {
			safeReports++
			continue
		}
	}

	log.Printf("%d\n", safeReports)

	// part 2
	safeReports2 := 0
	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp)
	lineNo := 0
	for scanner.Scan() {
		lineNo++

		line := scanner.Text()
		numbers := strings.Split(line, " ")
		
		safeReport, problematicIndex := check(numbers)
		if safeReport {
			safeReports2++
			continue
		}

		numbersWithoutLeft := []string{}
		for i, num := range numbers {
			if i != problematicIndex {
				numbersWithoutLeft = append(numbersWithoutLeft, num)
			}
		}
		safeReport, _ = check(numbersWithoutLeft)
		if safeReport {
			safeReports2++
			continue
		}

		numbersWithoutRight := []string{}
		for i, num := range numbers {
			if i != problematicIndex+1 {
				numbersWithoutRight = append(numbersWithoutRight, num)
			}
		}
		safeReport, _ = check(numbersWithoutRight)
		if safeReport {
			safeReports2++
			continue
		}

		// let's try the one the left of the left
		if problematicIndex > 0 {
			numbersWithoutLeftLeft := []string{}
			for i, num := range numbers {
				if i != problematicIndex-1 {
					numbersWithoutLeftLeft = append(numbersWithoutLeftLeft, num)
				}
			}
			safeReport, _ = check(numbersWithoutLeftLeft)
			if safeReport {
				safeReports2++
				continue
			}		
		}
	}

	log.Printf("%d\n", safeReports2)
}

func check(numbers []string) (bool, int) {
	wasLastIncrease := false // hence it's decrease
	safeReport := false
	problematicIndex := -1
	for i:=0; i<len(numbers)-1; i++ {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[i+1])

		if a == b {
			problematicIndex = i
			break
		}

		if i == 0 {
			if b > a {
				wasLastIncrease = true
			}
		} else {
			if b > a && wasLastIncrease == false {
				problematicIndex = i
				break
			}
			if a > b && wasLastIncrease == true {
				problematicIndex = i
				break
			}
		}
		if wasLastIncrease && b-a > 3 {
			problematicIndex = i
			break
		}
		if !wasLastIncrease && a-b > 3 {
			problematicIndex = i
			break
		}

		if i == len(numbers)-2 {
			safeReport = true
		}
	}

	return safeReport, problematicIndex
}
