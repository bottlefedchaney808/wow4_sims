package shadow

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestShadowPriestDefaultRotationUsesAvailableDamageSpells(t *testing.T) {
	player := &proto.Player{}
	applyShadowPriestDefaultRotation(player)
	if player.Rotation == nil {
		t.Fatal("default rotation was not applied")
	}
	if got := len(player.Rotation.PriorityList); got < 3 {
		t.Fatalf("default rotation has %d actions; want at least 3", got)
	}
	want := map[int32]bool{10894: false, 10947: false, 18807: false}
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

func TestShadowPriestDefaultRotationPreservesExplicitRotation(t *testing.T) {
	explicit := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	player := &proto.Player{Rotation: explicit}
	applyShadowPriestDefaultRotation(player)
	if player.Rotation != explicit {
		t.Fatal("explicit rotation was replaced")
	}
}
