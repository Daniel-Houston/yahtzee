package main

import(
	"sort"
)

type RuleFulfillment struct {
	Fulfills bool
	Points int
	IncludedVals []int
}

type Rule interface{
	Fulfills(dice []int) *RuleFulfillment
}

type YahtzeeRule struct { }
type FullHouseRule struct { }
type StraightRule struct { }
type OfAKindRule struct { 
	//Number of the '# of a kind' to look for
	Number int
}
type TwoPairRule struct{ }

func (r *YahtzeeRule) Fulfills(dice []int) *RuleFulfillment {
	a := dice[0]
	for _, val := range dice {
		if a != val {
			return &RuleFulfillment{
				Fulfills: false,
			}
		}
	}

	return &RuleFulfillment{
		Fulfills: true,
		Points: 50,
		IncludedVals: dice,
	}
}

func (r *FullHouseRule) Fulfills(dice []int) *RuleFulfillment {
	counts := getCounts(dice)

	doublesVal := 0
	triplesVal := 0
	for val, count := range counts {
		if count == 2 {
			doublesVal = val
		}
		if count == 3 {
			triplesVal = val
		}
	}

	if doublesVal != 0 && triplesVal != 0 {
		return &RuleFulfillment{
			Fulfills: true,
			Points: (doublesVal * 2) + (triplesVal * 3),
			IncludedVals: []int{doublesVal, doublesVal, triplesVal, triplesVal, triplesVal},
		}
	}

	return &RuleFulfillment{
		Fulfills: false,
	}
}

func (r *StraightRule) Fulfills(dice []int) *RuleFulfillment {
	sort.Ints(dice)

	curr := dice[0]
	for i := 1; i < len(dice); i++{
		next := dice[i]
		if (curr+1) != next {
			return &RuleFulfillment{
				Fulfills: false,
			}
		}
		curr = next
	}

	points := 0
	for _, val := range dice{
		points += val
	}

	return &RuleFulfillment{
		Fulfills: true,
		Points :  points,
		IncludedVals: dice,
	}
}

func (r *OfAKindRule) Fulfills(dice []int) *RuleFulfillment {
	counts := getCounts(dice)

	for key, count := range counts {
		if count == r.Number {
			return &RuleFulfillment{
				Fulfills: true,
				Points: r.Number*key,
				IncludedVals: func() []int {
					vals := make([]int, 0, r.Number)
					for i := 0; i < r.Number; i++ {
						vals = append(vals, key)
					}
					return vals
				}(),
			}
		}
	}

	return &RuleFulfillment{
		Fulfills: false,
	}
}

func (r *TwoPairRule) Fulfills(dice []int) *RuleFulfillment {
	counts := getCounts(dice)

	firstPair := 0
	secondPair := 0

	for num, count := range counts {
		if count == 2 {
			if firstPair == 0 {
				firstPair = num
				continue
			}
			secondPair = num
		}
	}

	if firstPair == 0 || secondPair == 0 {
		return &RuleFulfillment{
			Fulfills: false,
		}
	}

	includedVals := []int{firstPair, firstPair, secondPair, secondPair}
	sort.Ints(includedVals)

	return &RuleFulfillment{
		Fulfills: true,
		Points: (2*firstPair) + (2*secondPair),
		IncludedVals: includedVals,
	}
}

func getCounts(dice []int) map[int]int {
	counts := map[int]int{}
	for _, val := range dice{
		counts[val]++
	}

	return counts
}
