package rogue

import "github.com/wowsims/classic/sim/core"

// Data-only registrations preserve Forever coefficients pending ability/APL
// integration; these entries are intentionally not exposed as APL actions.
func (rogue *Rogue) registerForeverCoefficientSpells() {
	rogue.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1241585}, // Mutilate triggered hit (base 48)
		SpellSchool:      core.SpellSchoolPhysical,
		BonusCoefficient: 1.0,
	})
	rogue.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1310703}, // Venom (base 30)
		SpellSchool:      core.SpellSchoolNature,
		BonusCoefficient: 1.0,
	})
	rogue.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1310728}, // Quietus rank data
		SpellSchool:      core.SpellSchoolPhysical,
		BonusCoefficient: 1.0,
	})
}
