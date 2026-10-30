package dpswarrior

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/warrior"
)

func applyDpsWarriorDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultDpsWarriorRotation()
	}
}

func defaultDpsWarriorRotation() *proto.APLRotation {
	return warrior.DefaultDpsAPLRotation()
}
