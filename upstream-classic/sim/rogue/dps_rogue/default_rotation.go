package dpsrogue

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func applyRogueDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultRogueRotation()
	}
}

func defaultRogueRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"condition":{"or":{"vals":[
				{"and":{"vals":[
					{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"1"}}}},
					{"not":{"val":{"auraIsActive":{"auraId":{"spellId":6774,"rank":2}}}}},
					{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"6s"}}}}
				]}},
				{"and":{"vals":[
					{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"5"}}}},
					{"cmp":{"op":"OpLt","lhs":{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}},"rhs":{"const":{"val":"3s"}}}},
					{"cmp":{"op":"OpGt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"9s"}}}}
				]}}
			]}},"castSpell":{"spellId":{"spellId":6774,"rank":2}}}},
			{"action":{"condition":{"auraIsActive":{"auraId":{"spellId":6774,"rank":2}}},"autocastOtherCooldowns":{}}},
			{"action":{"condition":{"and":{"vals":[
				{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"5"}}}},
				{"or":{"vals":[
					{"auraIsActive":{"auraId":{"spellId":6774,"rank":2}}},
					{"cmp":{"op":"OpGe","lhs":{"currentEnergy":{}},"rhs":{"const":{"val":"79"}}}},
					{"cmp":{"op":"OpLt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"6s"}}}}
				]}}
			]}},"castSpell":{"spellId":{"spellId":31016}}}},
			{"action":{"castSpell":{"spellId":{"spellId":11294,"rank":8}}}}
		]
	}`)
}
