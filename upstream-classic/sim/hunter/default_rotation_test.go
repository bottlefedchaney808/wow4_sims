package hunter

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestDefaultRotationAppliedOnlyWhenMissing(t *testing.T) {
	player := &proto.Player{}
	applyDefaultRotation(player)
	if player.Rotation == nil || len(player.Rotation.PriorityList) == 0 {
		t.Fatal("expected a non-empty hunter default rotation")
	}

	explicit := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	player.Rotation = explicit
	applyDefaultRotation(player)
	if player.Rotation != explicit {
		t.Fatal("explicit rotation was replaced")
	}
}
