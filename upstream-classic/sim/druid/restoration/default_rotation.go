package restoration

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const defaultRotationJSON = `{
	"type": "TypeAPL",
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"not":{"val":{"dotIsActive":{"targetUnit":{"type":4},"spellId":{"spellId":1238215}}}}},"castSpell":{"spellId":{"spellId":1238215},"target":{"type":4}}}}
	]
}`

func applyDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = core.APLRotationFromJsonString(defaultRotationJSON)
	}
}
