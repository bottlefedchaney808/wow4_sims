package hunter

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Forever Sniper Shot rows from the extracted MythicSim table.
var ForeverSniperShotSpellIDs = []int32{1310687, 1310785, 1310786}
var ForeverSniperShotCoefficients = []float64{1, 1, 1}
var foreverSniperShotBaseValues = []float64{160, 225, 295}

func (hunter *Hunter) registerForeverSniperShotSpell() {
	hunter.SniperShot = make([]*core.Spell, len(ForeverSniperShotSpellIDs))
	for rank, spellID := range ForeverSniperShotSpellIDs {
		baseValue := foreverSniperShotBaseValues[rank]
		coefficient := ForeverSniperShotCoefficients[rank]
		hunter.SniperShot[rank] = hunter.GetOrRegisterSpell(core.SpellConfig{
			SpellCode:        SpellCode_HunterSniperShot,
			ActionID:         core.ActionID{SpellID: spellID},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeRanged,
			ProcMask:         core.ProcMaskRangedSpecial,
			Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagShot,
			CastType:         proto.CastType_CastTypeRanged,
			Rank:             rank + 1,
			MissileSpeed:     24,
			CritDamageBonus:  hunter.mortalShots(),
			BonusCoefficient: coefficient,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
				return hunter.DistanceFromTarget >= core.MinRangedAttackDistance
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcDamage(sim, target, baseValue, spell.OutcomeRangedHitAndCrit)
				spell.WaitTravelTime(sim, func(sim *core.Simulation) {
					spell.DealDamage(sim, result)
				})
			},
		})
	}
}
