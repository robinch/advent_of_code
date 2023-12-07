package day7

import (
	"bufio"
	"cmp"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	highCard      = 0
	onePair       = 1
	twoPairs      = 2
	fullHouse     = 4
	threeOfAKind  = 3
	fourOfAKind   = 5
	fiveOfAKind   = 6
	cardValueBase = 15
)

type Hand struct {
	Type int
	Bid  int
	Val  int
}

func Part1(filePath string) int {
	return getTotalWinnings(filePath, false)
}

func Part2(filePath string) int {
	return getTotalWinnings(filePath, true)
}

func getTotalWinnings(filePath string, isPart2 bool) int {
	// Group by hand type (lowest to highest)
	// Sort each group by hand value (lowest to highest)
	// loop from start to end and increase rank per hand

	totalWinnings := 0

	groupedAndSortedHands := groupAndSortHands(filePath, isPart2)

	rank := 1
	for i := 0; i < len(groupedAndSortedHands); i++ {
		for j := 0; j < len(groupedAndSortedHands[i]); j++ {
			totalWinnings += rank * groupedAndSortedHands[i][j].Bid
			rank++
		}
	}

	return totalWinnings
}

func groupAndSortHands(filePath string, isPart2 bool) [7][]Hand {
	groupedAndSortedHands := [7][]Hand{}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		hand := getHand(scanner.Text(), isPart2)
		groupedAndSortedHands[hand.Type] = append(groupedAndSortedHands[hand.Type], hand)
	}

	for i := 0; i < len(groupedAndSortedHands); i++ {
		slices.SortFunc(groupedAndSortedHands[i], func(a, b Hand) int { return cmp.Compare(a.Val, b.Val) })
	}

	return groupedAndSortedHands
}

func getHand(line string, isPart2 bool) Hand {
	split := strings.Split(line, " ")
	hand := []rune(split[0])
	bid := split[1]

	return Hand{
		Type: getType(hand, isPart2),
		Val:  getHandVal(hand, isPart2),
		Bid:  getBid(bid),
	}
}

func getHandVal(hand []rune, isPart2 bool) int {
	val := 0
	for i := 0; i < len(hand); i++ {
		val += getCardValue(hand[i], isPart2) * int(math.Pow(float64(cardValueBase), float64(len(hand)-1-i)))
	}
	return val
}

func getType(hand []rune, isPart2 bool) int {
	sortedCardCount, jokers := getSortedCardCountAndJokers(hand, isPart2)

	if jokers == 5 || sortedCardCount[0]+jokers == 5 {
		return fiveOfAKind
	} else if sortedCardCount[0]+jokers == 4 {
		return fourOfAKind
	} else if sortedCardCount[0]+jokers == 3 && sortedCardCount[1] == 2 {
		return fullHouse
	} else if sortedCardCount[0]+jokers == 3 {
		return threeOfAKind
	} else if sortedCardCount[0]+jokers == 2 && sortedCardCount[1] == 2 {
		return twoPairs
	} else if sortedCardCount[0]+jokers == 2 {
		return onePair
	} else {
		return highCard
	}
}

func getSortedCardCountAndJokers(hand []rune, isPart2 bool) ([]int, int) {
	cardCountByCardType := map[rune]int{}
	sortedCardCount := []int{}

	for _, card := range hand {
		cardCountByCardType[card]++
	}

	for card, count := range cardCountByCardType {
		if isPart2 && card == 'J' {
			continue
		}

		sortedCardCount = append(sortedCardCount, count)
	}

	slices.Sort(sortedCardCount)
	slices.Reverse(sortedCardCount)

	jokers := 0
	if isPart2 {
		jokers = cardCountByCardType['J']
	}

	return sortedCardCount, jokers
}

func getBid(s string) int {
	bid, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return bid
}

func getCardValue(card rune, isPart2 bool) int {
	switch card {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		if isPart2 {
			return 1
		} else {
			return 11
		}
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		panic("Invalid card!")
	}
}
