package restoration

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultRestorationRotationCastsHealingRanks(t *testing.T) {
	rotation := defaultRestorationRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"331", "1064", "8004"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default restoration APL missing spell ID %s: %s", id, encoded)
		}
	}
}

func TestApplyRestorationDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyRestorationDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
