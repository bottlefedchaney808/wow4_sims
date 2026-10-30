package tankwarrior

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/warrior"
)

func applyTankWarriorDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultTankWarriorRotation()
	}
}

func defaultTankWarriorRotation() *proto.APLRotation {
	return warrior.DefaultTankAPLRotation()
}
