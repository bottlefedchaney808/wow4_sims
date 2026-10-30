package elemental

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyElementalDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultElementalRotation()
	}
}

func defaultElementalRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"prepullActions":[
			{"action":{"castSpell":{"spellId":{"spellId":15208,"rank":10}}},"doAtValue":{"const":{"val":"-2s"}}}
		],
		"priorityList":[
			{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpEq","lhs":{"numberTargets":{}},"rhs":{"const":{"val":"1"}}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10438,"rank":6}}}}}]}},"castSpell":{"spellId":{"spellId":10438,"rank":6}}}},
			{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpGe","lhs":{"numberTargets":{}},"rhs":{"const":{"val":"2"}}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10587,"rank":4}}}}}]}},"castSpell":{"spellId":{"spellId":10587,"rank":4}}}},
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"castSpell":{"spellId":{"spellId":10605,"rank":4}}}},
			{"action":{"condition":{"auraIsActive":{"auraId":{"spellId":16166}}},"castSpell":{"spellId":{"spellId":10414,"rank":7}}}},
			{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"50%"}}}},"castSpell":{"spellId":{"spellId":15208,"rank":10}}}},
			{"action":{"castSpell":{"spellId":{"spellId":915,"rank":4}}}},
			{"action":{"condition":{"const":{"val":"false"}},"castSpell":{"spellId":{"spellId":20572}}}}
		]
	}`)
}
