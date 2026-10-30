package druid

import "github.com/wowsims/classic/sim/core"

// Forever Mangle rows from the extracted MythicSim table.
var ForeverMangleSpellIDs = []int32{1238069, 1238070, 1238073}
var ForeverMangleCoefficients = []float64{1, 1, 1}
var foreverMangleBaseValues = []float64{38, 59, 77}

func (druid *Druid) registerForeverMangle(form DruidForm) {
	for rank, spellID := range ForeverMangleSpellIDs {
		baseValue := foreverMangleBaseValues[rank]
		coefficient := ForeverMangleCoefficients[rank]
		druid.RegisterSpell(form, core.SpellConfig{
			ActionID:         core.ActionID{SpellID: spellID},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			Flags:            core.SpellFlagAPL,
			Rank:             rank + 1,
			BonusCoefficient: coefficient,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, baseValue, spell.OutcomeMeleeSpecialHitAndCrit)
			},
		})
	}
}
