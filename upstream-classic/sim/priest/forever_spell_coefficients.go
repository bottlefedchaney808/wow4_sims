package priest

import (
	"github.com/wowsims/classic/sim/core"
)

// These data-only spell registrations preserve Forever-era coefficients until
// the corresponding healing/damage mechanics and APL actions are implemented.
func (priest *Priest) registerForeverCoefficientSpells() {
	for _, entry := range []struct {
		spellID int32
		school  core.SpellSchool
		coeff   float64
	}{
		{spellID: 1240720, school: core.SpellSchoolHoly, coeff: 0.285},   // Penance (r2; base 113)
		{spellID: 1309595, school: core.SpellSchoolShadow, coeff: 0.429}, // Shadow Word: Death (base 302)
		{spellID: 1277331, school: core.SpellSchoolHoly, coeff: 0.143},   // Chastise (base 46)
	} {
		priest.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: entry.spellID},
			SpellSchool:      entry.school,
			BonusCoefficient: entry.coeff,
		})
	}
}
