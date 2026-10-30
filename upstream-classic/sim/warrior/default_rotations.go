package warrior

import (
	_ "embed"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

//go:embed dps_warrior/default.apl.json
var dpsDefaultAPL string

//go:embed tank_warrior/default.apl.json
var tankDefaultAPL string

func DefaultDpsAPLRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(dpsDefaultAPL)
}

func DefaultTankAPLRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(tankDefaultAPL)
}
