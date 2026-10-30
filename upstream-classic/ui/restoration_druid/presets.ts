import * as PresetUtils from '../core/preset_utils.js';
import { Consumes, Debuffs, Flask, Food, IndividualBuffs, PartyBuffs, RaidBuffs, TristateEffect, UnitReference } from '../core/proto/common.js';
import { RestorationDruid_Options as RestorationDruidOptions } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import BlankGear from './gear_sets/blank.gear.json';
import DefaultApl from './apls/default.apl.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const DefaultGear = PresetUtils.makePresetGear('Custom Healing Gear', BlankGear, {
	tooltip: 'Starts empty. Equip your own healing gear before comparing HPS or spell stat weights.',
});
export const ROTATION_PRESET_RESTO = {
	...PresetUtils.makePresetAPLRotation('Restoration', DefaultApl),
	tooltip: 'Prioritizes party-targeted healing over time effects and Swiftmend when its HoT requirement is met.',
};

// Valid 15-talent Restoration tree strings; all ranks are within the tree's limits.
export const CelestialFocusTalents = {
	name: 'Swiftmend Restoration',
	data: SavedTalents.create({
		talentsString: '--555523025515251',
	}),
};
export const ThiccRestoTalents = {
	name: 'Thicc Resto',
	data: SavedTalents.create({
		talentsString: '--555523045515051',
	}),
};

export const DefaultOptions = RestorationDruidOptions.create({
	innervateTarget: UnitReference.create(),
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskUnknown,
	food: Food.FoodUnknown,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	moonkinAura: true,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectRegular,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultPartyBuffs = PartyBuffs.create({});

export const DefaultDebuffs = Debuffs.create({
	faerieFire: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	distanceFromTarget: 18,
};
