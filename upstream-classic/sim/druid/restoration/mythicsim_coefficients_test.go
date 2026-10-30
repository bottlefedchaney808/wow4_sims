package restoration

import (
	"reflect"
	"testing"
)

func TestMythicSimWildGrowthRows(t *testing.T) {
	if !reflect.DeepEqual(ForeverWildGrowthSpellIDs, []int32{1238214, 1238215}) {
		t.Fatalf("Wild Growth spell IDs = %v", ForeverWildGrowthSpellIDs)
	}
	if !reflect.DeepEqual(ForeverWildGrowthCoefficients, []float64{.033, .033}) {
		t.Fatalf("Wild Growth coefficients = %v", ForeverWildGrowthCoefficients)
	}
}
