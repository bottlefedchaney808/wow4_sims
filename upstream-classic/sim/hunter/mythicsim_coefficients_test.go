package hunter

import (
	"reflect"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestMythicSimHunterCoefficientCorrections(t *testing.T) {
	hunter := &Hunter{Talents: &proto.HunterTalents{}}
	for rank := 1; rank <= 8; rank++ {
		if got := hunter.getArcaneShotConfig(rank, hunter.NewTimer()).BonusCoefficient; got != 0 {
			t.Errorf("Arcane Shot rank %d direct coefficient = %v, want 0", rank, got)
		}
	}
	for rank := 1; rank <= 9; rank++ {
		if got := hunter.getSerpentStingConfig(rank).Dot.BonusCoefficient; got != 0 {
			t.Errorf("Serpent Sting rank %d periodic coefficient = %v, want 0", rank, got)
		}
	}
	for rank := 1; rank <= 3; rank++ {
		if got := hunter.getVolleyConfig(rank).Dot.BonusCoefficient; got != 0 {
			t.Errorf("Volley rank %d periodic coefficient = %v, want 0", rank, got)
		}
	}
}

func TestMythicSimHunterForeverSpellRows(t *testing.T) {
	if !reflect.DeepEqual(ForeverSniperShotSpellIDs, []int32{1310687, 1310785, 1310786}) {
		t.Fatalf("Sniper Shot spell IDs = %v", ForeverSniperShotSpellIDs)
	}
	if !reflect.DeepEqual(ForeverSniperShotCoefficients, []float64{1, 1, 1}) {
		t.Fatalf("Sniper Shot coefficients = %v", ForeverSniperShotCoefficients)
	}
}
