package holy_test

import (
	"testing"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestHolyPaladinDefaultHealingActionProducesHPS(t *testing.T) {
	sim.RegisterAll()

	player := &proto.Player{
		Name:          "Holy Paladin",
		Race:          proto.Race_RaceHuman,
		Class:         proto.Class_ClassPaladin,
		Equipment:     &proto.EquipmentSpec{},
		TalentsString: "00303000311011-000000000000000",
		Spec:          &proto.Player_HolyPaladin{HolyPaladin: &proto.HolyPaladin{Options: &proto.PaladinOptions{}}},
	}

	result := core.RunRaidSim(&proto.RaidSimRequest{
		Raid: &proto.Raid{
			Parties: []*proto.Party{{Players: []*proto.Player{player}}},
			Tanks:   []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: 0}},
		},
		Encounter: &proto.Encounter{
			Duration: 20,
			Targets: []*proto.Target{{
				Name:          "Training Dummy",
				Level:         60,
				MobType:       proto.MobType_MobTypeHumanoid,
				MinBaseDamage: 1800,
				SwingSpeed:    2,
				TankIndex:     0,
			}},
		},
		SimOptions: &proto.SimOptions{Iterations: 1, RandomSeed: 1},
	})

	if result.Error != nil {
		t.Fatalf("raid sim failed: %s", result.Error.Message)
	}
	if got := result.RaidMetrics.Parties[0].Players[0].Hps.Avg; got <= 0 {
		t.Fatalf("Holy Light did no effective healing: HPS = %v", got)
	}

	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() != 25292 {
			continue
		}
		for _, target := range action.Targets {
			if target.Casts > 0 && target.Healing > 0 {
				return
			}
		}
	}
	t.Fatal("Holy Light casts with healing were not present in the action metrics")
}
