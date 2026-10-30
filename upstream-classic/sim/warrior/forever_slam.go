package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// registerForeverSlamSpell adds the Forever rank while retaining the legacy
// Slam spell ID used by existing rotations.
func (warrior *Warrior) registerForeverSlamSpell() {
	warrior.RegisterSpell(AnyStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorSlam,
		ActionID:    core.ActionID{SpellID: 1240193},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,

		RequiredLevel: 20,
		Rank:          1,
		RageCost: core.RageCostOptions{
			Cost:   15,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 1500*time.Millisecond - time.Millisecond*100*time.Duration(warrior.Talents.ImprovedSlam),
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				if spell.CastTime() > 0 {
					warrior.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+cast.CastTime, true)
				}
			},
		},

		CritDamageBonus:  warrior.impale(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		FlatThreatBonus:  140,
		BonusCoefficient: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 16 + spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
}

// Spearing Strike (1310222) is a Forever instant attack: 40% weapon damage,
// 20-second cooldown, and an additional 80% weapon damage against giants and
// dragonkin. Its exported coefficient is 1.0.
func (warrior *Warrior) registerSpearingStrikeSpell() {
	cdTimer := warrior.NewTimer()
	warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:      core.ActionID{SpellID: 1310222},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,
		RequiredLevel: 1,
		Rank:          1,
		RageCost:      core.RageCostOptions{Cost: 15, Refund: 0.8},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			CD:          core.Cooldown{Timer: cdTimer, Duration: 20 * time.Second},
		},
		CritDamageBonus:  warrior.impale(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target)) * 0.4
			if target.MobType == proto.MobType_MobTypeGiant || target.MobType == proto.MobType_MobTypeDragonkin {
				baseDamage *= 3
			}
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
}
