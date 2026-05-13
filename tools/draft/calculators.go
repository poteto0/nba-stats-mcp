package draft

import (
	"math"

	"github.com/poteto0/go-nba-sdk/types"
)

func calculateSimilarityScore(record1, record2 types.DraftCombineStatsRecord) float64 {
	var sumSquares float64
	var count int

	// with shoesはほとんど変わらないので、飛ばす
	if record1.HeightWoShoes != nil && record2.HeightWoShoes != nil {
		diff := *record1.HeightWoShoes - *record2.HeightWoShoes
		sumSquares += HeightWoShoesWeight * diff * diff
		count++
	}
	if record1.Weight != nil && record2.Weight != nil {
		diff := *record1.Weight - *record2.Weight
		sumSquares += WeightWeight * diff * diff
		count++
	}
	if record1.Wingspan != nil && record2.Wingspan != nil {
		diff := *record1.Wingspan - *record2.Wingspan
		sumSquares += WingspanWeight * diff * diff
		count++
	}
	if record1.StandingVertical != nil && record2.StandingVertical != nil {
		diff := *record1.StandingVertical - *record2.StandingVertical
		sumSquares += StandingVerticalWeight * diff * diff
		count++
	}
	if record1.MaxVertical != nil && record2.MaxVertical != nil {
		diff := *record1.MaxVertical - *record2.MaxVertical
		sumSquares += MaxVerticalWeight * diff * diff
		count++
	}
	if record1.StandingReach != nil && record2.StandingReach != nil {
		diff := *record1.StandingReach - *record2.StandingReach
		sumSquares += StandingReachWeight * diff * diff
		count++
	}
	if record1.LaneAgility != nil && record2.LaneAgility != nil {
		diff := *record1.LaneAgility - *record2.LaneAgility
		sumSquares += LaneAgilityWeight * diff * diff
		count++
	}
	if record1.ThreeQuarterSprint != nil && record2.ThreeQuarterSprint != nil {
		diff := *record1.ThreeQuarterSprint - *record2.ThreeQuarterSprint
		sumSquares += ThreeQuarterSprintWeight * diff * diff
		count++
	}
	if count == 0 {
		return 0
	}

	return math.Sqrt(sumSquares) / float64(count)
}
