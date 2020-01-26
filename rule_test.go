package main

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule_Fulfills(t *testing.T) {
	type inputs struct{
		nums []int
	}
	type outputs struct{
		Fulfillment RuleFulfillment
	}
	tests := []struct{
		name string
		toTest Rule
		inputs inputs
		outputs outputs
	} {
		{
			name: "yahtzee",
			toTest: &YahtzeeRule{},
			inputs: inputs{
				nums: []int{1, 1, 1, 1, 1},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 50,
					IncludedVals: []int{1, 1, 1, 1, 1},
				},
			},
		},
		{
			name: "fullHouse",
			toTest: &FullHouseRule{},
			inputs: inputs{
				nums: []int{1, 1, 2, 2, 2},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 8,
					IncludedVals: []int{1, 1, 2, 2, 2},
				},
			},
		},
		{
			name: "small straight",
			toTest: &StraightRule{},
			inputs: inputs{
				nums: []int{3, 4, 1, 2, 5},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 15,
					IncludedVals: []int{1, 2, 3, 4, 5},
				},
			},
		},
		{
			name: "large straight",
			toTest: &StraightRule{},
			inputs: inputs{
				nums: []int{4, 6, 3, 2, 5},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 20,
					IncludedVals: []int{2, 3, 4, 5, 6},
				},
			},
		},
		{
			name: "4 of a kind",
			toTest: &OfAKindRule{4},
			inputs: inputs{
				nums: []int{4, 4, 4, 4, 5},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 16,
					IncludedVals: []int{4, 4, 4, 4},
				},
			},
		},
		{
			name: "3 of a kind",
			toTest: &OfAKindRule{3},
			inputs: inputs{
				nums: []int{4, 4, 4, 3, 5},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 12,
					IncludedVals: []int{4, 4, 4},
				},
			},
		},
		{
			// TODO there may be more than one pair here, make sure that is covered
			name: "2 of a kind",
			toTest: &OfAKindRule{2},
			inputs: inputs{
				nums: []int{4, 4, 1, 3, 5},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 8,
					IncludedVals: []int{4, 4},
				},
			},
		},
		{
			name: "Two Pair",
			toTest: &TwoPairRule{},
			inputs: inputs{
				nums: []int{2, 4, 6, 2, 4},
			},
			outputs: outputs{
				Fulfillment: RuleFulfillment{
					Fulfills: true,
					Points: 12,
					IncludedVals: []int{2,2, 4, 4},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.toTest
			f := r.Fulfills(test.inputs.nums)
			assert.Equal(t, test.outputs.Fulfillment.Fulfills, f.Fulfills) 
			assert.Equal(t, test.outputs.Fulfillment.Points, f.Points) 
			assert.Equal(t, test.outputs.Fulfillment.IncludedVals, f.IncludedVals) 
		})
	}

}

func TestRule_DoesNotFulfill(t *testing.T) {
	type inputs struct{
		nums []int
	}
	tests := []struct{
		name string
		toTest Rule
		inputs inputs
	} {
		{
			name: "yahtzee",
			toTest: &YahtzeeRule{},
			inputs: inputs{
				nums: []int{1,1,1,1,2},
			},
		},
		{
			name: "fullHouse",
			toTest: &FullHouseRule{},
			inputs: inputs{
				nums: []int{1,1,1,1,2},
			},
		},
		{
			name: "straight",
			toTest: &StraightRule{},
			inputs: inputs{
				nums: []int{3,4,2,6,1},
			},
		},
		{
			name: "straight",
			toTest: &StraightRule{},
			inputs: inputs{
				nums: []int{1,4,3,6,5},
			},
		},
		{
			name: "3 of a kind",
			toTest: &OfAKindRule{3},
			inputs: inputs{
				nums: []int{4,4,4,4,5},
			},
		},
		{
			name: "4 of a kind",
			toTest: &OfAKindRule{4},
			inputs: inputs{
				nums: []int{4,4,4,3,5},
			},
		},
		{
			name: "4 of a kind",
			toTest: &OfAKindRule{4},
			inputs: inputs{
				nums: []int{4,4,4,4,4},
			},
		},
		{
			name: "2 of a kind",
			toTest: &OfAKindRule{2},
			inputs: inputs{
				nums: []int{4,4,4,4,4},
			},
		},
		{
			name: "2 of a kind",
			toTest: &OfAKindRule{2},
			inputs: inputs{
				nums: []int{1,2,3,4,5},
			},
		},
		{
			name: "two pair",
			toTest: &TwoPairRule{},
			inputs: inputs{
				nums: []int{4,4,4,4,4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.toTest
			f := r.Fulfills(test.inputs.nums)
			assert.Equal(t, false, f.Fulfills) 
		})
	}

}
