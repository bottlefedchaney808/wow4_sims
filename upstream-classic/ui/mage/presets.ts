import { Phase } from '../core/constants/other';
import * as PresetUtils from '../core/preset_utils';
import {
	Conjured,
	Consumes,
	Debuffs,
	FirePowerBuff,
	Flask,
	Food,
	FrostPowerBuff,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common';
import { Mage_Options as MageOptions, Mage_Options_ArmorType as ArmorType } from '../core/proto/mage';
import { SavedTalents } from '../core/proto/ui';
import ArcaneAPL from './apls/arcane.apl.json';
import FireAPL from './apls/fire.apl.json';
import FrostAPL from './apls/frost.apl.json';
import P1APL from './apls/p1.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBaseline = PresetUtils.makePresetGear('Compatible baseline', BlankGear, {
	tooltip: 'Intentionally unequipped: the imported Pre-BiS/P1 item IDs are not in the active sim database.',
});

export const GearPresets = {
	[Phase.Phase1]: [GearBaseline],
};

export const DefaultGear = GearBaseline;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP1DPS = {
	...PresetUtils.makePresetAPLRotation('Frost DPS', P1APL),
	tooltip: 'Frost single-target priority with mana recovery and cooldown support.',
};
export const APLArcane = {
	...PresetUtils.makePresetAPLRotation('Arcane', ArcaneAPL),
	tooltip: 'Arcane-focused single-target rotation.',
};
export const APLFire = {
	...PresetUtils.makePresetAPLRotation('Fire', FireAPL),
	tooltip: 'Fire single-target rotation with Scorch support.',
};
export const APLFrost = {
	...PresetUtils.makePresetAPLRotation('Frost', FrostAPL),
	tooltip: 'Frostbolt-focused single-target rotation.',
};

export const APLPresets = {
	[Phase.Phase1]: [APLP1DPS, APLArcane, APLFire, APLFrost],
};

export const DefaultAPL = APLPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1DPS = PresetUtils.makePresetTalents('Frost DPS', SavedTalents.create({ talentsString: '230205021002--05353203102351001' }));

export const TalentPresets = {
	[Phase.Phase1]: [TalentsP1DPS],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = MageOptions.create({
	armor: ArmorType.MoltenArmor,
});

export const DefaultConsumes = Consumes.create({
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	frostPowerBuff: FrostPowerBuff.ElixirOfFrostPower,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	improvedScorch: true,
	judgementOfWisdom: true,
	wintersChill: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Alchemy,
	profession2: Profession.Tailoring,
};
