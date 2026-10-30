package elemental

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultElementalRotationUsesForeverLavaBurstAndCorrectedTopRanks(t *testing.T) {
	rotation := defaultElementalRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"10605", "10414", "15208", "915"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default elemental APL missing spell ID %s: %s", id, encoded)
		}
	}
	if !strings.Contains(string(encoded), "50%") {
		t.Fatalf("default elemental APL lost its mana-efficient max-rank threshold: %s", encoded)
	}
}

func TestApplyElementalDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyElementalDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
