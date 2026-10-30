package mage

import "github.com/wowsims/classic/sim/core"

// Forever spell rows extracted from MythicSim. Cast times and mana costs were
// not present in the source rows, so these registrations intentionally only
// port the table's base values and coefficients; default APLs do not use them.
var ForeverArcaneBlastSpellIDs = []int32{1239696, 1239697, 1239699, 1239700}
var ForeverArcaneBlastCoefficients = []float64{.714, .714, .714, .714}
var foreverArcaneBlastBaseValues = []float64{141, 181, 290, 394}

var ForeverFrostfireBoltSpellIDs = []int32{1237312, 1237313}
var ForeverFrostfireBoltCoefficients = []float64{.814, .814}
var ForeverFrostfireBoltPeriodicCoefficients = []float64{0, 0}
var foreverFrostfireBoltBaseValues = []float64{195, 292}

func (mage *Mage) registerForeverArcaneBlast() {
	mage.ArcaneBlast = make([]*core.Spell, len(ForeverArcaneBlastSpellIDs))
	for rank, spellID := range ForeverArcaneBlastSpellIDs {
		baseValue := foreverArcaneBlastBaseValues[rank]
		coefficient := ForeverArcaneBlastCoefficients[rank]
		mage.ArcaneBlast[rank] = mage.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: spellID},
			SpellCode:        SpellCode_MageArcaneBlast,
			SpellSchool:      core.SpellSchoolArcane,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellDamage,
			Flags:            SpellFlagMage | core.SpellFlagAPL,
			Rank:             rank + 1,
			BonusCoefficient: coefficient,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, baseValue, spell.OutcomeMagicHitAndCrit)
			},
		})
	}
}

func (mage *Mage) registerForeverFrostfireBolt() {
	mage.FrostfireBolt = make([]*core.Spell, len(ForeverFrostfireBoltSpellIDs))
	for rank, spellID := range ForeverFrostfireBoltSpellIDs {
		baseValue := foreverFrostfireBoltBaseValues[rank]
		coefficient := ForeverFrostfireBoltCoefficients[rank]
		mage.FrostfireBolt[rank] = mage.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: spellID},
			SpellCode:        SpellCode_MageFrostfireBolt,
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellDamage,
			Flags:            SpellFlagMage | core.SpellFlagAPL,
			Rank:             rank + 1,
			BonusCoefficient: coefficient,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, baseValue, spell.OutcomeMagicHitAndCrit)
			},
		})
	}
}
