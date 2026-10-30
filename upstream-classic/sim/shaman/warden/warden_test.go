package warden

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	RegisterWardenShaman()
}

func TestWardenShaman(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassShaman,
			Phase:      1,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     DefaultTalents,
			GearSet:     core.GetGearSet("../../../ui/warden_shaman/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/warden_shaman/apls", "default"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Default", SpecOptions: PlayerOptionsBasic},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var PlayerOptionsBasic = &proto.Player_WardenShaman{
	WardenShaman: &proto.WardenShaman{
		Options: &proto.WardenShaman_Options{},
	},
}

var DefaultTalents = "5203015-0505000145503151"

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		DefaultConjured:   proto.Conjured_ConjuredDemonicRune,
		DefaultPotion:     proto.Potions_MajorManaPotion,
		DragonBreathChili: true,
		FirePowerBuff:     proto.FirePowerBuff_ElixirOfGreaterFirepower,
		Flask:             proto.Flask_FlaskOfTheTitans,
		Food:              proto.Food_FoodBlessSunfruit,
		MainHandImbue:     proto.WeaponImbue_RockbiterWeapon,
		SpellPowerBuff:    proto.SpellPowerBuff_GreaterArcaneElixir,
		StrengthBuff:      proto.StrengthBuff_JujuPower,
	},
}

var ItemFilters = core.ItemFilter{
	ArmorType: proto.ArmorType_ArmorTypeMail,

	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
		proto.WeaponType_WeaponTypeShield,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStrength,
	proto.Stat_StatAttackPower,
	proto.Stat_StatSpellPower,
	proto.Stat_StatArmor,
	proto.Stat_StatDodge,
	proto.Stat_StatParry,
	proto.Stat_StatBlockValue,
	proto.Stat_StatDefense,
}

func TestWardenDefaultsRockbiterWhenNoMainHandImbueIsSelected(t *testing.T) {
	character := &core.Character{Consumes: &proto.Consumes{}}

	setWardenDefaultImbue(character)

	if got := character.Consumes.MainHandImbue; got != proto.WeaponImbue_RockbiterWeapon {
		t.Fatalf("default main-hand imbue = %v, want Rockbiter Weapon", got)
	}
}

func TestWardenPreservesExplicitMainHandImbue(t *testing.T) {
	character := &core.Character{Consumes: &proto.Consumes{
		MainHandImbue: proto.WeaponImbue_FlametongueWeapon,
	}}

	setWardenDefaultImbue(character)

	if got := character.Consumes.MainHandImbue; got != proto.WeaponImbue_FlametongueWeapon {
		t.Fatalf("explicit main-hand imbue changed to %v", got)
	}
}

func TestWardenDefaultRotationIsConfigured(t *testing.T) {
	rotation := defaultTankRotation()
	if rotation == nil || len(rotation.PriorityList) < 4 {
		t.Fatal("default tank rotation is missing maintenance, threat, or attack actions")
	}
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"10432", "10438", "10414", "17364"} {
		if !strings.Contains(string(encoded), id) {
			t.Errorf("default Warden APL missing spell ID %s: %s", id, encoded)
		}
	}
}

func TestWardenTauntUsesDatabaseSpellID(t *testing.T) {
	if wardenTauntSpellID != 20859 {
		t.Fatalf("taunt spell ID = %d, want 20859", wardenTauntSpellID)
	}
}

func TestWardenTauntThreatValueIsDatabaseValue(t *testing.T) {
	if wardenTauntThreat != 1 {
		t.Fatalf("taunt threat = %v, want DB2/Wowhead value 1", wardenTauntThreat)
	}
}

func TestWardenDefaultRotationPreservesSuppliedRotation(t *testing.T) {
	customRotation := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	player := &proto.Player{Rotation: customRotation}
	applyWardenDefaultRotation(player)
	if player.Rotation != customRotation {
		t.Fatal("default tank rotation replaced an explicitly supplied rotation")
	}
}

func TestWardenDefaultRotationFillsMissingRotation(t *testing.T) {
	player := &proto.Player{}
	applyWardenDefaultRotation(player)
	if player.Rotation == nil || len(player.Rotation.PriorityList) < 3 {
		t.Fatal("missing rotation was not populated with the tank APL")
	}
}
