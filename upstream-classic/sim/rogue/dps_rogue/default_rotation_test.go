package dpsrogue

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestRogueDefaultRotationUsesBuilderAndFinishers(t *testing.T) {
	player := &proto.Player{}
	applyRogueDefaultRotation(player)
	if player.Rotation == nil {
		t.Fatal("default rotation was not applied")
	}
	want := map[int32]bool{6774: false, 11294: false, 31016: false}
	for _, item := range player.Rotation.PriorityList {
		if cast := item.Action.GetCastSpell(); cast != nil && cast.SpellId != nil {
			if _, ok := want[cast.SpellId.GetSpellId()]; ok {
				want[cast.SpellId.GetSpellId()] = true
			}
		}
	}
	for spellID, found := range want {
		if !found {
			t.Errorf("default rotation is missing spell ID %d", spellID)
		}
	}
}

func TestRogueDefaultRotationPreservesExplicitRotation(t *testing.T) {
	explicit := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	player := &proto.Player{Rotation: explicit}
	applyRogueDefaultRotation(player)
	if player.Rotation != explicit {
		t.Fatal("explicit rotation was replaced")
	}
}
