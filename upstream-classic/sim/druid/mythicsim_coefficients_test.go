package druid

import (
	"reflect"
	"testing"
)

func TestMythicSimDruidCoefficients(t *testing.T) {
	for rank, want := range []float64{0, .15, .15, .15} {
		if got := MoonfiresSpellCoeff[rank]; got != want {
			t.Errorf("Moonfire rank %d direct coefficient = %v, want %v", rank, got, want)
		}
	}
	for rank, want := range []float64{0, .13, .13, .13} {
		if got := MoonfiresSellDotCoeff[rank]; got != want {
			t.Errorf("Moonfire rank %d periodic coefficient = %v, want %v", rank, got, want)
		}
	}
	for rank, want := range []float64{0, .429, .486, .571} {
		if got := WrathSpellCoeff[rank]; got != want {
			t.Errorf("Wrath rank %d coefficient = %v, want %v", rank, got, want)
		}
	}
}

func TestMythicSimDruidForeverSpellRows(t *testing.T) {
	if !reflect.DeepEqual(ForeverMangleSpellIDs, []int32{1238069, 1238070, 1238073}) {
		t.Fatalf("Mangle spell IDs = %v", ForeverMangleSpellIDs)
	}
	if !reflect.DeepEqual(ForeverMangleCoefficients, []float64{1, 1, 1}) {
		t.Fatalf("Mangle coefficients = %v", ForeverMangleCoefficients)
	}
}
