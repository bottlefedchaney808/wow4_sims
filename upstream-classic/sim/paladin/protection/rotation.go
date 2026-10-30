package protection

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyProtectionPaladinDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultProtectionPaladinRotation()
	}
}

func defaultProtectionPaladinRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"prepullActions":[
			{"action":{"castSpell":{"spellId":{"spellId":20423,"rank":7}}},"doAtValue":{"const":{"val":"-3s"}}},
			{"action":{"castSpell":{"spellId":{"spellId":20928,"rank":3}}},"doAtValue":{"const":{"val":"-1.5s"}}}
		],
		"priorityList":[
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"currentSealRemainingTime":{}},"rhs":{"const":{"val":"1.5s"}}}},"castSpell":{"spellId":{"spellId":20423,"rank":7}}}},
			{"action":{"condition":{"spellCanCast":{"spellId":{"spellId":20271}}},"castSpell":{"spellId":{"spellId":20271}}}},
			{"action":{"castSpell":{"spellId":{"spellId":20928,"rank":3}}}},
			{"action":{"castSpell":{"spellId":{"spellId":20924,"rank":5}}}}
		]
	}`)
}
