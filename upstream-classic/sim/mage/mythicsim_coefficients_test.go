package mage

import (
	"reflect"
	"testing"
)

func TestMythicSimMageCoefficients(t *testing.T) {
	cases := []struct {
		name string
		got  float64
		want float64
	}{
		{"Arcane Explosion r1", ArcaneExplosionSpellCoeff[1], .143},
		{"Arcane Missiles r1", ArcaneMissilesSpellCoeff[1], .286},
		{"Arcane Missiles r2", ArcaneMissilesSpellCoeff[2], .286},
		{"Arcane Missiles r3", ArcaneMissilesSpellCoeff[3], .286},
		{"Arcane Missiles r8", ArcaneMissilesSpellCoeff[8], .286},
		{"Fire Blast r1", FireBlastSpellCoeff[1], .429},
		{"Fire Blast r2", FireBlastSpellCoeff[2], .429},
		{"Fireball r1", FireballSpellCoeff[1], .429},
		{"Fireball r2", FireballSpellCoeff[2], .571},
		{"Fireball r3", FireballSpellCoeff[3], .714},
		{"Fireball r4", FireballSpellCoeff[4], .857},
		{"Flamestrike r1 direct", FlamestrikeSpellCoeff[1], .157},
		{"Flamestrike r1 periodic", FlamestrikeDotCoeff[1], .032},
		{"Flamestrike r6 periodic", FlamestrikeDotCoeff[6], .032},
		{"Frostbolt r1", FrostboltSpellCoeff[1], .407},
		{"Frostbolt r2", FrostboltSpellCoeff[2], .489},
		{"Frostbolt r3", FrostboltSpellCoeff[3], .597},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("coefficient = %v, want %v", tc.got, tc.want)
			}
		})
	}
}

func TestMythicSimMageForeverSpellRows(t *testing.T) {
	if !reflect.DeepEqual(ForeverArcaneBlastSpellIDs, []int32{1239696, 1239697, 1239699, 1239700}) {
		t.Fatalf("Arcane Blast spell IDs = %v", ForeverArcaneBlastSpellIDs)
	}
	if !reflect.DeepEqual(ForeverArcaneBlastCoefficients, []float64{.714, .714, .714, .714}) {
		t.Fatalf("Arcane Blast coefficients = %v", ForeverArcaneBlastCoefficients)
	}
	if !reflect.DeepEqual(ForeverFrostfireBoltSpellIDs, []int32{1237312, 1237313}) {
		t.Fatalf("Frostfire Bolt spell IDs = %v", ForeverFrostfireBoltSpellIDs)
	}
	if !reflect.DeepEqual(ForeverFrostfireBoltCoefficients, []float64{.814, .814}) {
		t.Fatalf("Frostfire Bolt coefficients = %v", ForeverFrostfireBoltCoefficients)
	}
}
