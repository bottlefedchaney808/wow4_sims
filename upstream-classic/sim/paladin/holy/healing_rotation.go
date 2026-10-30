package holy

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyHolyPaladinDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultHolyPaladinHealingRotation()
	}
}

func defaultHolyPaladinHealingRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"castSpell":{"spellId":{"spellId":20473}}}},
			{"action":{"castSpell":{"spellId":{"spellId":25292}}}},
			{"action":{"castSpell":{"spellId":{"spellId":19943}}}}
		]
	}`)
}
