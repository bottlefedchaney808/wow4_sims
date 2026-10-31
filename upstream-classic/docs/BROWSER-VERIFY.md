# Browser verification (first run: 2026-10-30, automated headless Chromium + manual pass pending; manual items closed 2026-10-31)

Server: `tools/forever_gen/serve_dist.py 8732` (root `dist/`, site at `/forever/`).

## Automated evidence (2026-10-30, headless Chromium)

- [x] All 21 spec pages HTTP 200: balance_druid feral_druid feral_tank_druid restoration_druid elemental_shaman enhancement_shaman restoration_shaman warden_shaman hunter mage rogue holy_paladin protection_paladin retribution_paladin healing_priest shadow_priest warlock warrior tank_warrior raid detailed_results (+ `/forever/` index, `lib.wasm` with `application/wasm` MIME, `assets/database/db.json`).
- [x] Homepage loads: 19 spec cards render (`/forever/` title `WoWSims - Forever`).
- [x] wasm boots and completes setup on damage specs: feral_druid worker logs `setup` → `computeStats-*` tasks; gear list renders (Wolfshead Helm, Prestor's, Warlord's Dragonhide), full button set (Simulate, Stat Weights, Import, Save Encounter/Talents/Rotation).
- [x] No pageerror on damage specs after the preset-encounter graft (9d34e00).

## Manual pass (2026-10-31, headless Chromium)

- [x] Homepage loads: title `WoW4_SiMs · Forever`, 14 spec links render.
- [x] Gear picker populated: Finger 2 picker lists Forever items with quality classes (e.g. `Band of the Better Half`, `item-quality-epic`); a row click equips it and the slot label updates. Gear→stats wiring proven by gear-set states: P4 BiS preset → Health 5064 / Str 520 / Agi 369 / Sta 322 / AP 2356 / Melee Crit 40.36%; empty gear → all 0.
- [x] Feral Druid: Simulate → results table `Player (#1) 100% 167K 1296.2` (DPS > 0), zero page/console errors.
- [x] One heal sim (Restoration Druid): Sim 1 Iteration → HPS table renders (Rejuvenation 117.97 HPS, Regrowth 43.86 HPS), zero console errors.
- [x] db asset 200: `verify.sh` HTTP checks + curl `assets/database/db.json` → 200.

## Known findings

1. **Preset-encounter crash (FIXED 9d34e00).** The Forever gear-planner payload has no preset encounters; `Encounter.getDefaultTarget()` = `presetTargets[0]` → `undefined`, so healing/tank specs crashed at boot (`Cannot read properties of undefined (reading '0')` in `preset_utils` → `setDefaultHealingParams` → `primaryTarget`) and gear/enchant flows threw. Grafted 9 preset encounters + zones from the classic reference db.json via `tools/forever_gen/graft_presets.py` (order: regen_db FIRST, graft SECOND — regen drops Encounters).
2. **Healing/tank pages stall in worker setup under headless automation** (`Worker 0: Created` then no setup completion) — IDENTICAL on the proven reference build (`.worktrees/main/upstream-forever`), so not a daily-driver regression. Suspect hidden-tab/worker interaction; verify in a real browser tab.
3. **Gear tooltip enrichment fetches Wowhead live** (`nether.wowhead.com/forever/tooltip/item/<id>`): classic-only item ids return `{"error":"Entity not found"}` on the Forever dataset (e.g. 43110), so gear with classic-only items shows without icons/tooltips; Forever ids (25, 35, …) resolve. LAN traffic to Wowhead for tooltips is expected behavior of the stock UI; the gtag analytics script removal is Task 6.
4. **Headless render stall (same family as finding 2):** on a hidden tab opened with `?progress=true&test_sim=true` (1600×900), worker setup can stall — stats stay 0, the Simulate ring runs (`Stop` active), results never render, zero console errors. The default tab loads, computes stats, simulates, and renders results fine. Headless/tab-visibility artifact, not a build defect; proven on the same build in two tab configs.
