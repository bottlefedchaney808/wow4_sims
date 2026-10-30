package mage

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const defaultRotationJSON = `{
	"type": "TypeAPL",
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"castSpell":{"spellId":{"spellId":25306,"rank":12}}}}
	]
}`

func applyDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = core.APLRotationFromJsonString(defaultRotationJSON)
	}
}
