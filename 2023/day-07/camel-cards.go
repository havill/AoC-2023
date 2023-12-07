package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func IsFiveOfAKind(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
		if counts[card] == 5 {
			return true
		}
	}
	return false
}

func IsFourOfAKind(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
		if counts[card] == 4 {
			return true
		}
	}
	return false
}

func IsFullHouse(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
	}

	var pair, threeOfAKind bool
	for _, count := range counts {
		if count == 2 {
			pair = true
		} else if count == 3 {
			threeOfAKind = true
		}
	}

	return pair && threeOfAKind
}

func IsThreeOfAKind(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
		if counts[card] == 3 {
			return true
		}
	}
	return false
}

func IsTwoPair(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
	}

	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}

	return pairs == 2
}

func IsOnePair(hand []Card) bool {
	counts := make(map[Card]int)
	for _, card := range hand {
		counts[card]++
	}

	for _, count := range counts {
		if count == 2 {
			return true
		}
	}

	return false
}

func less(a, b []Card) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func totalWinnings(hands []Parsed) int {
	total := 0
	for rank, hand := range hands {
		total += hand.Bid * rank
	}
	return total
}

func main() {
	var hands []Parsed = make([]Parsed, 1)
	// prefix with sentinel value so first real hand is rank 1
	hands[0] = Parsed{Hand: []Card{Zero, Zero, Zero, Zero, Zero}, Bid: 0, Type: -1}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var x Parsed

		line := scanner.Text()
		//fmt.Fprintf(os.Stderr, "line = '%s'\n", line)

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
		x.Bid = i
		for _, c := range hand {
			card, err := charToCard(c)
			if err != nil {
				// handle error
			}
			x.Hand = append(x.Hand, card)
		}
		if IsFiveOfAKind(x.Hand) {
			x.Type = Five_of_a_kind
		} else if IsFourOfAKind(x.Hand) {
			x.Type = Four_of_a_kind
		} else if IsFullHouse(x.Hand) {
			x.Type = Full_house
		} else if IsThreeOfAKind(x.Hand) {
			x.Type = Three_of_a_kind
		} else if IsTwoPair(x.Hand) {
			x.Type = Two_pair
		} else if IsOnePair(x.Hand) {
			x.Type = One_pair
		} else {
			x.Type = High_card
		}
		hands = append(hands, x)

		//fmt.Fprintln(os.Stderr, "Hand  = ", x.Hand)
		//fmt.Fprintln(os.Stderr, "Bid   = ", x.Bid)
		//fmt.Fprintln(os.Stderr, "Type  = ", x.Type)

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

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type != hands[j].Type {
			return hands[i].Type < hands[j].Type
		}
		return less(hands[i].Hand, hands[j].Hand)
	})

	// fmt.Fprintln(os.Stderr, "hands = ", hands)
	fmt.Println(totalWinnings(hands))

}
