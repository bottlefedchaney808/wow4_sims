package dpswarrior

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultDpsWarriorRotationUsesForeverAttacks(t *testing.T) {
	rotation := defaultDpsWarriorRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"1240193", "1310222"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default warrior APL missing spell ID %s: %s", id, encoded)
		}
	}
}

func TestApplyDpsWarriorDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyDpsWarriorDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
