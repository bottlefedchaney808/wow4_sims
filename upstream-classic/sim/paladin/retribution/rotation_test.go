package retribution

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultRetributionRotationUsesMaxRankSealOfFury(t *testing.T) {
	rotation := defaultRetributionPaladinRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), "20423") {
		t.Fatalf("default retribution APL missing max-rank Seal of Fury: %s", encoded)
	}
}

func TestApplyRetributionDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyRetributionPaladinDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
