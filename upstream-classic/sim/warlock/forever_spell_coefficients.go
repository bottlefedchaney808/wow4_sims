package warlock

import "github.com/wowsims/classic/sim/core"

// These data-only registrations retain Forever coefficients pending cast/APL
// integration; none is flagged as an APL-castable spell.
func (warlock *Warlock) registerForeverCoefficientSpells() {
	warlock.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1316697}, // Wrack (base 36, 6 ticks)
		SpellSchool: core.SpellSchoolShadow,
		Dot:         core.DotConfig{BonusCoefficient: 0.143},
	})
	warlock.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1293812}, // Incinerate r2 (base 156)
		SpellSchool:      core.SpellSchoolFire,
		BonusCoefficient: 0.714,
	})
	warlock.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1293817}, // Conflagrate r1 (base 99)
		SpellSchool:      core.SpellSchoolFire,
		BonusCoefficient: 0.429,
	})
}
