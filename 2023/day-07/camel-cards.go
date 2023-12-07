package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type HandType int

const (
	High_card HandType = iota
	One_pair
	Two_pair
	Three_of_a_kind
	Full_house
	Four_of_a_kind
	Five_of_a_kind
)

type Card int

const (
	Zero Card = iota // this is a sentinal value and should never be used
	One              // this is a bogus value to make enum Two start at 2
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Parsed struct {
	Hand []Card
	Bid  int
	Type HandType
}

func charToCard(c rune) (Card, error) {
	switch c {
	case '0':
		return Zero, nil
	case '1':
		return One, nil
	case '2':
		return Two, nil
	case '3':
		return Three, nil
	case '4':
		return Four, nil
	case '5':
		return Five, nil
	case '6':
		return Six, nil
	case '7':
		return Seven, nil
	case '8':
		return Eight, nil
	case '9':
		return Nine, nil
	case 'T':
		return Ten, nil
	case 'J':
		return Jack, nil
	case 'Q':
		return Queen, nil
	case 'K':
		return King, nil
	case 'A':
		return Ace, nil
	default:
		return Zero, fmt.Errorf("invalid character: %c", c)
	}
}

func removePunctuationAndWhitespace(s string) string {
	f := func(r rune) rune {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			return -1
		}
		return r
	}
	return strings.Map(f, s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var x Parsed

		line := scanner.Text()
		line = strings.TrimSpace(line)
		lastSpace := strings.LastIndex(line, " ")
		left := line[:lastSpace]
		right := line[lastSpace+1:]
		left = removePunctuationAndWhitespace(left)
		hand := strings.ToUpper(left)

		i, err := strconv.Atoi(right)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}

		fmt.Println("hand  = ", hand)
		fmt.Println("Bid   = ", x.Bid)

		x.Bid = i

		/*
			p := Player{
				Hand: []Card{One, Two, Three},
				Bid:  5,
			}
		*/
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
