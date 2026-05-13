package draft

import (
	"math"
	"testing"

	"github.com/poteto0/go-nba-sdk/types"
)

func TestCalculateSimilarityScore(t *testing.T) {
	type args struct {
		playerStats1 DraftCombineSimilarityRecord
		playerStats2 DraftCombineSimilarityRecord
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "identical players should have similarity score of 0",
			args: args{
				playerStats1: DraftCombineSimilarityRecord{
					DraftCombineStatsRecord: types.DraftCombineStatsRecord{
						PlayerName:         "Player A",
						HeightWoShoes:      new(float64(80)),
						Weight:             new(float64(100)),
						Wingspan:           new(float64(82)),
						MaxVertical:        new(float64(40)),
						StandingVertical:   new(float64(30)),
						StandingReach:      new(float64(120)),
						LaneAgility:        new(float64(4.5)),
						ThreeQuarterSprint: new(float64(3.0)),
					},
				},
				playerStats2: DraftCombineSimilarityRecord{
					DraftCombineStatsRecord: types.DraftCombineStatsRecord{
						PlayerName:         "Player B",
						HeightWoShoes:      new(float64(80)),
						Weight:             new(float64(100)),
						Wingspan:           new(float64(82)),
						MaxVertical:        new(float64(40)),
						StandingVertical:   new(float64(30)),
						StandingReach:      new(float64(120)),
						LaneAgility:        new(float64(4.5)),
						ThreeQuarterSprint: new(float64(3.0)),
					},
				},
			},
			want: 0,
		},
		{
			name: "different players should have positive similarity score",
			args: args{
				playerStats1: DraftCombineSimilarityRecord{
					DraftCombineStatsRecord: types.DraftCombineStatsRecord{
						PlayerName:         "Player A",
						HeightWoShoes:      new(float64(80)),
						Weight:             new(float64(100)),
						Wingspan:           new(float64(82)),
						MaxVertical:        new(float64(40)),
						StandingVertical:   new(float64(30)),
						StandingReach:      new(float64(120)),
						LaneAgility:        new(float64(4.5)),
						ThreeQuarterSprint: new(float64(3.0)),
					},
				},
				playerStats2: DraftCombineSimilarityRecord{
					DraftCombineStatsRecord: types.DraftCombineStatsRecord{
						PlayerName:         "Player B",
						HeightWoShoes:      new(float64(78)),
						Weight:             new(float64(95)),
						Wingspan:           new(float64(80)),
						MaxVertical:        new(float64(35)),
						StandingVertical:   new(float64(25)),
						StandingReach:      new(float64(118)),
						LaneAgility:        new(float64(4.7)),
						ThreeQuarterSprint: new(float64(1.0)),
					},
				},
			},
			want: math.Sqrt(
				HeightWoShoesWeight*math.Pow(80-78, 2)+
					WeightWeight*math.Pow(100-95, 2)+
					WingspanWeight*math.Pow(82-80, 2)+
					MaxVerticalWeight*math.Pow(40-35, 2)+
					StandingVerticalWeight*math.Pow(30-25, 2)+
					StandingReachWeight*math.Pow(120-118, 2)+
					LaneAgilityWeight*math.Pow(4.5-4.7, 2)+
					ThreeQuarterSprintWeight*math.Pow(3.0-1.0, 2),
			) / 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateSimilarityScore(tt.args.playerStats1.DraftCombineStatsRecord, tt.args.playerStats2.DraftCombineStatsRecord)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("calculateSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
