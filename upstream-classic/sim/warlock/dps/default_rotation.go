package dps

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyWarlockDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultWarlockRotation()
	}
}

func defaultWarlockRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"condition":{"cmp":{"op":"OpLt","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"10%"}}}},"castSpell":{"spellId":{"spellId":11689,"rank":6}}}},
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":11713,"rank":6}}}}},"castSpell":{"spellId":{"spellId":11713,"rank":6}}}},
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":25311,"rank":7}}}}},"castSpell":{"spellId":{"spellId":25311,"rank":7}}}},
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":25309,"rank":8}}}}},"castSpell":{"spellId":{"spellId":25309,"rank":8}}}},
			{"action":{"castSpell":{"spellId":{"spellId":25307,"rank":10}}}}
		]
	}`)
}
