package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	update := false
	pageOrderingRules := map[string]bool{}
	updatesInRightOrder := []string{}
	var middleSum int64
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			update = true
			continue
		}

		if !update {
			pageOrderingRules[line] = true
			continue
		}

		// just in case of empty lines at the end
		if line == "" && update {
			continue
		}

		// got update line
		updates := strings.Split(line, ",")
		invalid := false
		for i := 0; i < len(updates)-1; i++ {
			for j := i + 1; j < len(updates); j++ {
				key := fmt.Sprintf("%s|%s", updates[i], updates[j])
				if !pageOrderingRules[key] {
					invalid = true
					break
				}
				key = fmt.Sprintf("%s|%s", updates[j], updates[i])
				if pageOrderingRules[key] {
					invalid = true
					break
				}
			}
			if invalid {
				break
			}
		}
		if !invalid {
			updatesInRightOrder = append(updatesInRightOrder, line)
			if len(updates)%2 == 1 {
				middleNum, _ := strconv.ParseInt(updates[(len(updates)-1)/2], 10, 64)
				middleSum += middleNum
			}
		}
	}

	log.Printf("%v\n", updatesInRightOrder)
	log.Printf("%d\n", middleSum)
}
