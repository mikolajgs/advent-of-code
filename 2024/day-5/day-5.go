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
	updatesInWrongOrder := []string{}
	updatesInFixedOrder := []string{}
	var middleSum int64
	var middleSumOfFixed int64
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
			continue
		}

		updatesInWrongOrder = append(updatesInWrongOrder, line)

		updatesFixed := []string{}
		updatesAppended := map[int]bool{}
		// order items
		for len(updatesFixed) < len(updates) {
			for i := 0; i < len(updates); i++ {
				if _, ok := updatesAppended[i]; ok {
					continue
				}

				appendToFixed := false
				for j := 0; j < len(updates); j++ {
					_, ok := updatesAppended[j]
					if i == j || ok {
						continue
					}
					var key string
					if i < j {
						key = fmt.Sprintf("%s|%s", updates[i], updates[j])
					} else {
						key = fmt.Sprintf("%s|%s", updates[j], updates[i])
					}
					// rule ordering differently found, i cannot be put on the left\
					log.Printf("%d %d %s", i, j, key)
					if _, ok := pageOrderingRules[key]; ok {
						appendToFixed = true
						break
					}
				}
				if appendToFixed {
					updatesFixed = append(updatesFixed, updates[i])
					updatesAppended[i] = true
				}
			}
			log.Printf("it: %v", updatesFixed)
		}

		updatesInFixedOrder = append(updatesInFixedOrder, strings.Join(updatesFixed, ","))
		if len(updates)%2 == 1 {
			middleNum, _ := strconv.ParseInt(updatesFixed[(len(updatesFixed)-1)/2], 10, 64)
			middleSumOfFixed += middleNum
		}
	}

	log.Printf("%v\n", updatesInRightOrder)
	log.Printf("%v\n", updatesInWrongOrder)
	log.Printf("%v\n", updatesInFixedOrder)
	log.Printf("%d\n", middleSum)
	log.Printf("%d\n", middleSumOfFixed)
}
