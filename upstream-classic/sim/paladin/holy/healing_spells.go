package holy

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/paladin"
)

const (
	spellIDHolyLightRank9  = 25292
	spellIDFlashLightRank6 = 19943
)

type paladinHealRank struct {
	level    int32
	spellID  int32
	manaCost float64
	minHeal  float64
	maxHeal  float64
}

var holyLightRanks = []paladinHealRank{
	{level: 1, spellID: 635, manaCost: 35, minHeal: 39, maxHeal: 53},
	{level: 6, spellID: 639, manaCost: 70, minHeal: 76, maxHeal: 96},
	{level: 14, spellID: 647, manaCost: 120, minHeal: 159, maxHeal: 185},
	{level: 22, spellID: 1026, manaCost: 180, minHeal: 292, maxHeal: 334},
	{level: 30, spellID: 1042, manaCost: 260, minHeal: 397, maxHeal: 447},
	{level: 38, spellID: 3472, manaCost: 340, minHeal: 582, maxHeal: 648},
	{level: 46, spellID: 10328, manaCost: 425, minHeal: 832, maxHeal: 928},
	{level: 54, spellID: 10329, manaCost: 525, minHeal: 1105, maxHeal: 1235},
	{level: 60, spellID: spellIDHolyLightRank9, manaCost: 660, minHeal: 1590, maxHeal: 1770},
}

var flashOfLightRanks = []paladinHealRank{
	{level: 20, spellID: 19750, manaCost: 35, minHeal: 67, maxHeal: 77},
	{level: 26, spellID: 19939, manaCost: 50, minHeal: 102, maxHeal: 120},
	{level: 34, spellID: 19940, manaCost: 70, minHeal: 153, maxHeal: 169},
	{level: 42, spellID: 19941, manaCost: 90, minHeal: 208, maxHeal: 232},
	{level: 50, spellID: 19942, manaCost: 115, minHeal: 282, maxHeal: 318},
	{level: 58, spellID: spellIDFlashLightRank6, manaCost: 140, minHeal: 448, maxHeal: 502},
}

func (hp *HolyPaladin) registerHolyLight() {
	for rankIndex, rank := range holyLightRanks {
		if hp.Level < rank.level {
			break
		}
		rank := rank
		hp.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: rank.spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagHelpful | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          rankIndex + 1,
			SpellCode:     paladin.SpellCode_PaladinHolyLight,

			ManaCost: core.ManaCostOptions{FlatCost: rank.manaCost},
			Cast: core.CastConfig{DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 2500 * time.Millisecond,
			}},

			DamageMultiplier: 1 + 0.02*float64(hp.Talents.HealingLight),
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealHealing(sim, target, sim.Roll(rank.minHeal, rank.maxHeal), spell.OutcomeHealingCrit)
			},
		})
	}
}

func (hp *HolyPaladin) registerFlashOfLight() {
	for rankIndex, rank := range flashOfLightRanks {
		if hp.Level < rank.level {
			break
		}
		rank := rank
		hp.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: rank.spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagHelpful | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          rankIndex + 1,
			SpellCode:     paladin.SpellCode_PaladinFlashOfLight,

			ManaCost: core.ManaCostOptions{FlatCost: rank.manaCost},
			Cast: core.CastConfig{DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 1500 * time.Millisecond,
			}},

			DamageMultiplier: 1 + 0.02*float64(hp.Talents.HealingLight),
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealHealing(sim, target, sim.Roll(rank.minHeal, rank.maxHeal), spell.OutcomeHealingCrit)
			},
		})
	}
}

func (hp *HolyPaladin) RegisterHealingSpells() {
	hp.registerHolyLight()
	hp.registerFlashOfLight()
}
