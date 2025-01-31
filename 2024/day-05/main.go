package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadParseInput() (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter page ordering rules (blank line to end):")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "invalid input format")
			continue
		}
		before, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			fmt.Fprintln(os.Stderr, "invalid integer value")
			continue
		}
		rules[before] = append(rules[before], after)
	}

	fmt.Println("Enter updates (Ctrl+D to end):")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		var pageNumbers []int
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid integer value")
				continue
			}
			pageNumbers = append(pageNumbers, num)
		}
		updates = append(updates, pageNumbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return rules, updates
}

func isCorrectOrder(rules map[int][]int, update []int) bool {
	for i, pageNumber := range update {
		rule := rules[pageNumber]
		for j := 0; j < i; j++ {
			if contains(rule, update[j]) {
				return false
			}
		}
	}
	return true
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func middlePageNumber(pages []int) int {
	if len(pages)%2 == 0 {
		fmt.Fprintln(os.Stderr, "array does not have an odd number of elements")
		return -1
	}
	middleIndex := len(pages) / 2
	return pages[middleIndex]
}

func reorderUntilCorrect(rules map[int][]int, update []int) []int {
	for i, pageNumber := range update {
		rule := rules[pageNumber]
		for j := 0; j < i; j++ {
			if contains(rule, update[j]) {
				for k := j; k < len(update)-1; k++ {
					tooEarly := update[k]
					update[k] = update[k+1]
					update[k+1] = tooEarly
					if isCorrectOrder(rules, update) {
						break
					}
				}
			}
		}
	}
	return update
}

func main() {
	rules, updates := loadParseInput()

	middleCorrectSums := 0
	middleIncorrectSums := 0

	for _, update := range updates {
		if isCorrectOrder(rules, update) {
			middleCorrectSums += middlePageNumber(update)
		} else {
			newOrder := reorderUntilCorrect(rules, update)
			fmt.Println("Reordered to correct:", newOrder)
			middleIncorrectSums += middlePageNumber(newOrder)
		}
	}
	fmt.Print("Middle page number correct sums: ")
	fmt.Println(middleCorrectSums)
	fmt.Print("Middle page number incorrect sums: ")
	fmt.Println(middleIncorrectSums)
}
