package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Lava Burst ranks from the Forever spell table. These are registered so future
// APLs can target them by spell ID; the current default rotations do not use them.
var lavaBurstSpellIDs = [3]int32{408490, 1238299, 1238300}
var lavaBurstBaseDamage = [3]float64{120, 187, 220}
var lavaBurstManaCost = [3]float64{165, 230, 265}
var lavaBurstRequiredLevel = [3]int{40, 50, 60}
var lavaBurstSpellCoefficient = [3]float64{0.714, 0.714, 0.714}

func (shaman *Shaman) registerLavaBurstSpells() {
	cdTimer := shaman.NewTimer()
	for i := range lavaBurstSpellIDs {
		rank := i + 1
		baseDamage := lavaBurstBaseDamage[i]
		shaman.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: lavaBurstSpellIDs[i]},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       SpellFlagShaman | core.SpellFlagAPL,

			Rank:          rank,
			RequiredLevel: lavaBurstRequiredLevel[i],
			ManaCost: core.ManaCostOptions{
				FlatCost:   lavaBurstManaCost[i],
				Multiplier: 100 - 2*shaman.Talents.Convection,
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					CastTime: 2500 * time.Millisecond,
					GCD:      core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    cdTimer,
					Duration: 10 * time.Second,
				},
				ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
					castTime := shaman.ApplyCastSpeedForSpell(cast.CastTime, spell)
					shaman.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+castTime, false)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: lavaBurstSpellCoefficient[i],
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			},
		})
	}
}
