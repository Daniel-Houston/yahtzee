package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Points struct{
	RuleKey string
	Points int
	IncludedVals []int
}

var yahtzeeRules map[string]Rule = map[string]Rule{
	"yahtzee": &YahtzeeRule{},
	"fullHouse": &FullHouseRule{},
	"straight": &StraightRule{},
	"2 of a kind": &OfAKindRule{2},
	"3 of a kind": &OfAKindRule{3},
	"4 of a kind": &OfAKindRule{4},
	"two pair": &TwoPairRule{},
}

func main() {
	for {
		dice, err := getDice()
		if err != nil {
			log.Fatal(err)
		}

		p := FindYahtzeePoints(dice)
		fmt.Printf("Inputs: %v\n", dice)
		fmt.Printf("Points: %v, From: %v, IncludedVals: %v\n", p.Points, p.RuleKey, p.IncludedVals)
	}
}

func getDice() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the dice separated by commas: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	strs := strings.Split(strings.TrimSpace(text), ",")
	if len(strs) != 5 {
		return nil, fmt.Errorf("Wrong number of inputs. Want 5 got %v", len(strs))
	}

	dice := make([]int, 0)
	for _, s := range strs {
		die, err := strconv.Atoi(s)	
		if err != nil {
			return nil, err
		}

		if die < 1 || die > 6 {
			return nil, fmt.Errorf("Value %v is invalid, must be between 1 and 6", die)
		}

		dice = append(dice, die)
	}

	return dice, nil
}

func FindYahtzeePoints(dice []int) *Points{
	currentMax := Points{
		RuleKey: "no matching rule",
		Points: 0,
		IncludedVals: make([]int, 0),
	}

	for key, rule := range yahtzeeRules {
		f := rule.Fulfills(dice)
		if f.Fulfills && f.Points > currentMax.Points {
			currentMax = Points{
				RuleKey: key,
				Points: f.Points,
				IncludedVals: f.IncludedVals,
			}
		}
	}

	return &currentMax
}
