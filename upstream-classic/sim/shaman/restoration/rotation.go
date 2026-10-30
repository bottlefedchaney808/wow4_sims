package restoration

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyRestorationDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultRestorationRotation()
	}
}

func defaultRestorationRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"autocastOtherCooldowns":{}}},
			{"action":{"castSpell":{"spellId":{"spellId":331},"target":{"type":4,"index":0}}}},
			{"action":{"castSpell":{"spellId":{"spellId":1064},"target":{"type":4,"index":0}}}},
			{"action":{"castSpell":{"spellId":{"spellId":8004},"target":{"type":4,"index":0}}}}
		]
	}`)
}
