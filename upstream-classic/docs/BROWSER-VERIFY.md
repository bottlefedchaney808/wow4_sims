# Browser verification (first run: 2026-10-30, automated headless Chromium + manual pass pending)

Server: `tools/forever_gen/serve_dist.py 8732` (root `dist/`, site at `/forever/`).

## Automated evidence (2026-10-30, headless Chromium)

- [x] All 21 spec pages HTTP 200: balance_druid feral_druid feral_tank_druid restoration_druid elemental_shaman enhancement_shaman restoration_shaman warden_shaman hunter mage rogue holy_paladin protection_paladin retribution_paladin healing_priest shadow_priest warlock warrior tank_warrior raid detailed_results (+ `/forever/` index, `lib.wasm` with `application/wasm` MIME, `assets/database/db.json`).
- [x] Homepage loads: 19 spec cards render (`/forever/` title `WoWSims - Forever`).
- [x] wasm boots and completes setup on damage specs: feral_druid worker logs `setup` → `computeStats-*` tasks; gear list renders (Wolfshead Helm, Prestor's, Warlord's Dragonhide), full button set (Simulate, Stat Weights, Import, Save Encounter/Talents/Rotation).
- [x] No pageerror on damage specs after the preset-encounter graft (9d34e00).

## Manual pass required (human, real browser tab)

- [ ] Feral Druid: pick any grey→white item → stats panel updates.
- [ ] Feral Druid: Simulate → results table with DPS > 0, no console errors.
- [ ] One heal sim (Restoration Druid): Simulate → HPS > 0.
- [ ] db asset 200 in network tab (automated: yes; confirm in your browser).

## Known findings

1. **Preset-encounter crash (FIXED 9d34e00).** The Forever gear-planner payload has no preset encounters; `Encounter.getDefaultTarget()` = `presetTargets[0]` → `undefined`, so healing/tank specs crashed at boot (`Cannot read properties of undefined (reading '0')` in `preset_utils` → `setDefaultHealingParams` → `primaryTarget`) and gear/enchant flows threw. Grafted 9 preset encounters + zones from the classic reference db.json via `tools/forever_gen/graft_presets.py` (order: regen_db FIRST, graft SECOND — regen drops Encounters).
2. **Healing/tank pages stall in worker setup under headless automation** (`Worker 0: Created` then no setup completion) — IDENTICAL on the proven reference build (`.worktrees/main/upstream-forever`), so not a daily-driver regression. Suspect hidden-tab/worker interaction; verify in a real browser tab.
3. **Gear tooltip enrichment fetches Wowhead live** (`nether.wowhead.com/forever/tooltip/item/<id>`): classic-only item ids return `{"error":"Entity not found"}` on the Forever dataset (e.g. 43110), so gear with classic-only items shows without icons/tooltips; Forever ids (25, 35, …) resolve. LAN traffic to Wowhead for tooltips is expected behavior of the stock UI; the gtag analytics script removal is Task 6.
