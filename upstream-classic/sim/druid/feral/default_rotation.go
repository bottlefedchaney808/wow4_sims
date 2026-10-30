package feral

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const defaultRotationJSON = `{
	"type": "TypeAPL",
	"priorityList": [
		{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":9904,"rank":4}}}}},"castSpell":{"spellId":{"spellId":9904,"rank":4}}}},
		{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpEq","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"5"}}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":9896,"rank":6}}}}}]}},"castSpell":{"spellId":{"spellId":9896,"rank":6}}}},
		{"action":{"castSpell":{"spellId":{"spellId":9830}}}}
	]
}`

func applyDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = core.APLRotationFromJsonString(defaultRotationJSON)
	}
}
