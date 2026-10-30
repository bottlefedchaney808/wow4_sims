package paladin

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Seal of Fury player-cast ranks and their melee proc/Judgement rows from the
// Forever coefficient export. Default APLs do not select this seal yet.
var sealOfFuryRanks = [...]struct {
	spellID          int32
	procSpellID      int32
	judgementID      int32
	level            int32
	manaCost         float64
	procBase         float64
	judgementBase    float64
	spellCoefficient float64
	procCoefficient  float64
	judgeCoefficient float64
}{
	{1311649, 1311647, 1311650, 10, 40, 6, 33, 0.1, 0.1, 0.45},
	{1311656, 1311654, 1311655, 18, 60, 9, 48, 0.1, 0.1, 0.45},
	{20163, 20231, 20183, 25, 90, 14, 69, 0.1, 0.1, 0.45},
	{20419, 20415, 20411, 34, 120, 19, 90, 0.1, 0.1, 0.45},
	{20421, 20416, 20412, 42, 140, 25, 116, 0.1, 0.1, 0.45},
	{20422, 20417, 20413, 50, 170, 32, 145, 0.1, 0.1, 0.45},
	{20423, 20418, 20414, 58, 200, 35, 160, 0.1, 0.1, 0.45},
}

func (paladin *Paladin) registerSealOfFurySpells() {
	for i, rank := range sealOfFuryRanks {
		rankNumber := i + 1
		procSpell := paladin.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: rank.procSpellID},
			SpellSchool:      core.SpellSchoolHoly,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagSuppressEquipProcs,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: rank.procCoefficient,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, rank.procBase, spell.OutcomeMeleeSpecialCritOnly)
			},
		})

		judgementSpell := paladin.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: rank.judgementID},
			SpellSchool:      core.SpellSchoolHoly,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellDamage,
			Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagSuppressWeaponProcs | core.SpellFlagSuppressEquipProcs | core.SpellFlagBinary,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: rank.judgeCoefficient,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, rank.judgementBase, spell.OutcomeMagicHitAndCrit)
			},
		})
		paladin.spellsJoR = append(paladin.spellsJoR, judgementSpell)

		aura := paladin.RegisterAura(core.Aura{
			Label:    "Seal of Fury" + paladin.Label + strconv.Itoa(rankNumber),
			ActionID: core.ActionID{SpellID: rank.spellID},
			Duration: 30 * time.Second,
			OnSpellHitDealt: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) {
					procSpell.Cast(sim, result.Target)
				}
			},
		})

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:         aura.ActionID,
			SpellSchool:      core.SpellSchoolHoly,
			Flags:            core.SpellFlagAPL,
			RequiredLevel:    int(rank.level),
			Rank:             rankNumber,
			BonusCoefficient: rank.spellCoefficient,
			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost,
				Multiplier: paladin.benediction(),
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{GCD: core.GCDDefault},
			},
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
				paladin.applySeal(aura, judgementSpell, sim)
			},
		})
	}
}
