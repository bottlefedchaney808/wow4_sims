package enhancement

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultEnhancementRotationUsesMaxRankShockAndLightningShield(t *testing.T) {
	rotation := defaultEnhancementRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"10414", "10438", "10432"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default enhancement APL missing spell ID %s: %s", id, encoded)
		}
	}
}

func TestApplyEnhancementDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyEnhancementDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
