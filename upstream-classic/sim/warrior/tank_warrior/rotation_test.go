package tankwarrior

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultTankWarriorRotationRetainsTankAPL(t *testing.T) {
	rotation := defaultTankWarriorRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"11597", "1680"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default tank warrior APL missing spell ID %s: %s", id, encoded)
		}
	}
	for _, id := range []string{"1240193", "1310222"} {
		if strings.Contains(string(encoded), id) {
			t.Errorf("default tank warrior APL should not force weapon attacks with its UI default gear: %s", id)
		}
	}
}

func TestApplyTankWarriorDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyTankWarriorDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
