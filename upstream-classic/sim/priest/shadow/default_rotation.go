package shadow

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyShadowPriestDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultShadowPriestRotation()
	}
}

func defaultShadowPriestRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"prepullActions":[
			{"action":{"castSpell":{"spellId":{"spellId":15473}}},"doAtValue":{"const":{"val":"-1s"}}}
		],
		"priorityList":[
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10894,"rank":8}}}}},"castSpell":{"spellId":{"spellId":10894,"rank":8}}}},
			{"action":{"castSpell":{"spellId":{"spellId":10947,"rank":9}}}},
			{"action":{"castSpell":{"spellId":{"spellId":18807,"rank":6}}}}
		]
	}`)
}
