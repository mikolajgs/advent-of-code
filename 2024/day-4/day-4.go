package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fp, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(fp)
	scannedLines := 0
	foundXmas := 0
	lastLines := [4]string{}
	for scanner.Scan() {
		line := scanner.Text()
		scannedLines++
		if scannedLines > 0 && scannedLines < 4 {
			lastLines[scannedLines-1] = line // 0, 1, 2, 3
		}
		if scannedLines == 4 {
			lastLines[scannedLines-1] = line
		}
		if scannedLines > 4 {
			// move last 3 lines up and the latest as 4th
			lastLines[0] = lastLines[1]
			lastLines[1] = lastLines[2]
			lastLines[2] = lastLines[3]
			lastLines[3] = line
		}

		scanHorizontal(line, &foundXmas)
		if scannedLines > 3 {
			scanVertical(&lastLines, &foundXmas)
			scanDiagonal(&lastLines, &foundXmas)
		}
	}
	fp.Close()

	log.Printf("%d\n", foundXmas)

	// part 2
	fp, _ = os.Open("input.txt")
	defer fp.Close()
	scanner = bufio.NewScanner(fp)
	scannedLines = 0
	foundXmas = 0
	lastLines2 := [3]string{}
	for scanner.Scan() {
		line := scanner.Text()
		scannedLines++
		if scannedLines > 0 && scannedLines < 3 {
			lastLines2[scannedLines-1] = line // 0, 1, 2, 3
		}
		if scannedLines == 3 {
			lastLines2[scannedLines-1] = line
		}
		if scannedLines > 3 {
			// move last 3 lines up and the latest as 4th
			lastLines2[0] = lastLines2[1]
			lastLines2[1] = lastLines2[2]
			lastLines2[2] = line
		}
		// checking for X made of MAS
		if scannedLines > 2 {
			found := 0
			for col := 0; col < len(lastLines2[2])-2; col++ {
				if lastLines2[1][col+1] != []byte("A")[0] {
					continue
				}
				word1 := fmt.Sprintf("%c%c%c", lastLines2[2][col], lastLines2[1][col+1], lastLines2[0][col+2])
				word2 := fmt.Sprintf("%c%c%c", lastLines2[2][col+2], lastLines2[1][col+1], lastLines2[0][col])
				if (word1 == "SAM" || word1 == "MAS") && (word2 == "SAM" || word2 == "MAS") {
					found++
				}
			}
			if found > 0 {
				foundXmas += found
			}
		}
	}
	log.Printf("%d\n", foundXmas)
}

func scanHorizontal(line string, foundXmas *int) {
	cnt := strings.Count(line, "XMAS")
	cnt += strings.Count(line, "SAMX")
	*foundXmas += cnt
}

func scanVertical(lines *[4]string, foundXmas *int) {
	cnt := 0
	for col := 0; col < len(lines[3])-2; col++ {
		// scan upwards and downwards
		word := fmt.Sprintf("%c%c%c%c", lines[3][col], lines[2][col], lines[1][col], lines[0][col])
		if word == "XMAS" || word == "SAMX" {
			cnt++
		}
	}
	*foundXmas += cnt
}

func scanDiagonal(lines *[4]string, foundXmas *int) {
	for col := 0; col < len(lines[3]); col++ {
		if len(lines[3]) < 4 {
			break
		}

		// north right and left
		if col <= len(lines[3])-4 {
			word := fmt.Sprintf("%c%c%c%c", lines[3][col], lines[2][col+1], lines[1][col+2], lines[0][col+3])
			if word == "XMAS" || word == "SAMX" {
				*foundXmas++
			}
			word = fmt.Sprintf("%c%c%c%c", lines[3][col+3], lines[2][col+2], lines[1][col+1], lines[0][col])
			if word == "XMAS" || word == "SAMX" {
				*foundXmas++
			}
		}
	}
}
