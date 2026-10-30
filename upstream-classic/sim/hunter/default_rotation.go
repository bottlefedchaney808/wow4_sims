package hunter

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const defaultRotationJSON = `{
	"type": "TypeAPL",
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":25295,"rank":9}}}}},"castSpell":{"spellId":{"spellId":25295,"rank":9}}}},
		{"action":{"castSpell":{"spellId":{"spellId":14287,"rank":8}}}}
	]
}`

func applyDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = core.APLRotationFromJsonString(defaultRotationJSON)
	}
}
