#!/usr/bin/env python3
"""Graft preset encounters + npcs/factions/spellIcons/zones from the classic
reference db.json into the Forever db.json (in place, tracked file).

Why: the Forever gear-planner payload carries no preset encounters. The UI's
Encounter.getDefaultTarget() takes presetTargets[0]; with an empty encounters
list every healing/tank spec crashes at boot (primaryTarget -> clone(undefined)).
ORDER MATTERS: run `go run ./tools/forever_gen/regen_db.go` FIRST, then this
script, then rebuild the wasm (`GOOS=js GOARCH=wasm go build -tags=with_db ...`)
and `make copydb`. regen_db's WowDatabase struct carries no Encounters field,
so WriteJson re-emits db.json WITHOUT the graft if the graft runs before regen.
The UI reads db.encounters straight from the JSON; db.bin (wasm/headless) does
not need presets.

Usage: graft_presets.py <reference-db.json> [target-db.json]
"""
import json
import sys

FIELDS = ("encounters", "npcs", "factions", "spellIcons", "zones")


def main():
    if len(sys.argv) < 2:
        sys.exit(__doc__)
    ref_path = sys.argv[1]
    tgt_path = sys.argv[2] if len(sys.argv) > 2 else "assets/database/db.json"

    ref = json.load(open(ref_path, encoding="utf-8"))
    tgt = json.load(open(tgt_path, encoding="utf-8"))

    for enc in ref.get("encounters", []):
        if not any(e.get("path") == enc.get("path") for e in tgt.get("encounters", [])):
            tgt.setdefault("encounters", []).append(enc)
    for f in ("npcs", "factions", "spellIcons", "zones"):
        # id-keyed maps in the UI loader; replace the (empty) Forever lists wholesale
        if not tgt.get(f):
            tgt[f] = ref.get(f, [])

    report = {f: len(tgt.get(f, [])) for f in FIELDS}
    report["items"] = len(tgt.get("items", []))
    with open(tgt_path, "w", encoding="utf-8") as fh:
        json.dump(tgt, fh)
    print("grafted:", report)
    if report["encounters"] == 0:
        sys.exit("no preset encounters available from reference — STOP")


if __name__ == "__main__":
    main()
