package balance

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const defaultRotationJSON = `{
	"type": "TypeAPL",
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":9835,"rank":10}}}}},"castSpell":{"spellId":{"spellId":9835,"rank":10}}}},
		{"action":{"castSpell":{"spellId":{"spellId":9912,"rank":8}}}}
	]
}`

func applyDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = core.APLRotationFromJsonString(defaultRotationJSON)
	}
}
