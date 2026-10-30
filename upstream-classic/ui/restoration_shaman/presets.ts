import * as PresetUtils from '../core/preset_utils.js';
import { Consumes, Flask, Food, WeaponImbue } from '../core/proto/common.js';
import { RestorationShaman_Options as RestorationShamanOptions } from '../core/proto/shaman.js';
import { SavedTalents } from '../core/proto/ui.js';
import BlankGear from './gear_sets/blank.gear.json';
import DefaultApl from './apls/default.apl.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const DefaultGear = PresetUtils.makePresetGear('No Gear (Healing Baseline)', BlankGear, {
	tooltip: 'An empty gear baseline. Equip healing gear before comparing realistic raid-healing output.',
});
export const ROTATION_PRESET_RESTO = PresetUtils.makePresetAPLRotation('Restoration', DefaultApl);

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.
export const TankHealingTalents = {
	name: 'Tank Healing',
	data: SavedTalents.create({
		talentsString: '--502053313353101',
	}),
};
export const RaidHealingTalents = {
	name: 'Raid Healing',
	data: SavedTalents.create({
		talentsString: '--050053313353101',
	}),
};

export const DefaultOptions = RestorationShamanOptions.create({
	earthShieldPPM: 0,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskUnknown,
	food: Food.FoodUnknown,
	mainHandImbue: WeaponImbue.RockbiterWeapon,
	offHandImbue: WeaponImbue.RockbiterWeapon,
});
