import {
	Consumes,
	Flask,
	Food,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';

import {
	PaladinAura,
	PaladinOptions as HolyPaladinOptions,
} from '../core/proto/paladin.js';

import * as PresetUtils from '../core/preset_utils.js';

import APLDefaultHolyJson from './apls/holy_default.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const DefaultGear = PresetUtils.makePresetGear('Blank', BlankGear);

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const StandardTalents = {
	name: 'Healing Baseline',
	data: SavedTalents.create({
		// Match the PaladinTalents proto field layout (14 Holy, 15 Protection,
		// 15 Retribution fields). The legacy Wowhead string overflows these
		// segments and assigns rank 5 to Improved Lay on Hands, panicking the sim.
		talentsString: '00303000311011-000000000000000',
	}),
};

export const DefaultAPL = PresetUtils.makePresetAPLRotation('Holy Shock + Seal', APLDefaultHolyJson);

export const DefaultOptions = HolyPaladinOptions.create({
	aura: PaladinAura.DevotionAura,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskUnknown,
	food: Food.FoodUnknown,
});
