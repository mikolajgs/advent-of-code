package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	leftColumn := []int{}
	rightColumn := []int{}

	re := regexp.MustCompile(`[0-9]+`)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		found := re.FindAllString(line, -1)
		left, _ := strconv.Atoi(found[0])
		right, _ := strconv.Atoi(found[1])
		leftColumn = append(leftColumn, left)
		rightColumn = append(rightColumn, right)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)
	
	distance := 0
	for i := 0; i < len(leftColumn); i++ {
		d := leftColumn[i] - rightColumn[i]
		if d < 0 {
			d = d * -1
		}
		distance += d
	}

	log.Printf("%d\n", distance)

	// part 2
	occurrencesOnTheRight := map[int]int{}
	lastNumber := -1
	currentCount := 0
	for i := 0; i < len(rightColumn); i++ {
		currentNumber := rightColumn[i];

		if currentNumber != lastNumber {
			// change number, write to hashmap, reset the counter
			if lastNumber != -1 {
				occurrencesOnTheRight[lastNumber] = currentCount
			}
			currentCount = 0
			currentCount++
			lastNumber = currentNumber
			continue
		}
		currentCount++
		// last one so write to hashmap
		if i == len(rightColumn) - 1 {
			occurrencesOnTheRight[lastNumber] = currentCount
		}
	}

	similarityScore := 0;
	for i := 0; i < len(leftColumn); i++ {
		occOnTheRight, ok := occurrencesOnTheRight[leftColumn[i]]
		if ok {
			similarityScore += leftColumn[i] * occOnTheRight
		}
	}

	log.Printf("%d\n", similarityScore)
}
