package warden

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	// 1.60.1.69913 Wago SpellName.csv labels this as Rockbiter Weapon Taunt.
	// Wowhead Classic spell 20859 reports a one-point Threat effect:
	// https://www.wowhead.com/classic/spell=20859/rockbiter-weapon-taunt
	wardenTauntSpellID int32   = 20859
	wardenTauntThreat  float64 = 1
)

func setWardenDefaultImbue(character *core.Character) {
	if character == nil || character.Consumes == nil {
		return
	}
	if character.Consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueUnknown {
		// Rockbiter is the only sourced Shaman imbue here with an explicit
		// threat-per-second bonus; preserve any explicit user-selected imbue.
		character.Consumes.MainHandImbue = proto.WeaponImbue_RockbiterWeapon
	}
}

func applyWardenDefaultRotation(player *proto.Player) {
	if player != nil && player.Rotation == nil {
		player.Rotation = defaultTankRotation()
	}
}

// defaultTankRotation is only injected when the player supplied no APL at all.
// Earth Shock's existing Shaman registration applies the simulator's 2x threat
// multiplier; Stormstrike and Lightning Shield reuse their existing mechanics.
func defaultTankRotation() *proto.APLRotation {
	return core.APLRotationFromJsonString(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":10432,"rank":7}}}}},"castSpell":{"spellId":{"spellId":10432,"rank":7}}}},
			{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10438,"rank":6}}}}},"castSpell":{"spellId":{"spellId":10438,"rank":6}}}},
			{"action":{"castSpell":{"spellId":{"spellId":10414,"rank":7}}}},
			{"action":{"castSpell":{"spellId":{"spellId":17364}}}},
			{"action":{"autocastOtherCooldowns":{}}}
		]
	}`)
}

func (warden *WardenShaman) registerTauntSpell() {
	warden.Taunt = warden.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: wardenTauntSpellID},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ApplyEffects: func(_ *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Spell 20859's Classic tooltip reports Effect: Threat, Value: 1.
			// This simulator exposes per-spell threat metrics, not an enemy
			// threat table, so a threat-lead/forced-target swap cannot be modeled.
			spell.SpellMetrics[target.UnitIndex].TotalThreat += wardenTauntThreat * warden.PseudoStats.ThreatMultiplier
		},
	})
}
