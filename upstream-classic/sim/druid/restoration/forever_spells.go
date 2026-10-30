package restoration

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/druid"
)

// Forever Wild Growth rows from the extracted MythicSim table.
var ForeverWildGrowthSpellIDs = []int32{1238214, 1238215}
var ForeverWildGrowthCoefficients = []float64{.033, .033}
var foreverWildGrowthBaseValues = []float64{70, 97}

func (resto *RestorationDruid) registerForeverWildGrowth() {
	resto.ForeverWildGrowth = make([]*druid.DruidSpell, len(ForeverWildGrowthSpellIDs))
	for rank, spellID := range ForeverWildGrowthSpellIDs {
		baseValue := foreverWildGrowthBaseValues[rank]
		coefficient := ForeverWildGrowthCoefficients[rank]
		resto.ForeverWildGrowth[rank] = resto.RegisterSpell(druid.Humanoid, core.SpellConfig{
			ActionID:         core.ActionID{SpellID: spellID},
			SpellSchool:      core.SpellSchoolHoly,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellHealing,
			Flags:            core.SpellFlagHelpful | core.SpellFlagAPL,
			Rank:             rank + 1,
			Cast:             core.CastConfig{DefaultCast: core.Cast{GCD: core.GCDDefault}},
			DamageMultiplier: resto.healingPowerMultiplier(),
			ThreatMultiplier: 1,
			Hot: core.DotConfig{
				Aura: core.Aura{
					Label:    "Forever Wild Growth",
					ActionID: core.ActionID{SpellID: spellID},
				},
				NumberOfTicks:    7,
				TickLength:       time.Second,
				BonusCoefficient: coefficient,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, rollover bool) {
					dot.Snapshot(target, baseValue, rollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				if resto.Party == nil {
					spell.Hot(target).Apply(sim)
					return
				}
				for _, agent := range resto.Party.Players {
					spell.Hot(&agent.GetCharacter().Unit).Apply(sim)
				}
			},
		})
	}
}
