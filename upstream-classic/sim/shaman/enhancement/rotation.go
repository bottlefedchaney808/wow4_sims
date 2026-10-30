package enhancement

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyEnhancementDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultEnhancementRotation()
	}
}

func defaultEnhancementRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"prepullActions":[
			{"action":{"castSpell":{"spellId":{"spellId":10432,"rank":7}}},"doAtValue":{"const":{"val":"-1.5s"}}}
		],
		"priorityList":[
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10438,"rank":6}}}}},"castSpell":{"spellId":{"spellId":10438,"rank":6}}}},
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"castSpell":{"spellId":{"spellId":17364,"rank":1}}}},
			{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"50%"}}}},"castSpell":{"spellId":{"spellId":10414,"rank":7}}}}
		]
	}`)
}
