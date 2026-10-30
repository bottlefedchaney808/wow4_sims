package dps

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestWarlockDefaultRotationUsesHighestAvailableDamageRanks(t *testing.T) {
	player := &proto.Player{}
	applyWarlockDefaultRotation(player)
	if player.Rotation == nil {
		t.Fatal("default rotation was not applied")
	}
	want := map[int32]bool{11713: false, 25311: false, 25309: false, 25307: false, 11689: false}
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

func TestWarlockDefaultRotationPreservesExplicitRotation(t *testing.T) {
	explicit := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	player := &proto.Player{Rotation: explicit}
	applyWarlockDefaultRotation(player)
	if player.Rotation != explicit {
		t.Fatal("explicit rotation was replaced")
	}
}
