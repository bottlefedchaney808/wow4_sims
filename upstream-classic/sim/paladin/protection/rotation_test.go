package protection

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDefaultProtectionRotationUsesMaxRankSealOfFury(t *testing.T) {
	rotation := defaultProtectionPaladinRotation()
	encoded, err := protojson.Marshal(rotation)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), "20423") {
		t.Fatalf("default protection APL missing max-rank Seal of Fury: %s", encoded)
	}
}

func TestApplyProtectionDefaultRotationPreservesCustomRotation(t *testing.T) {
	custom := &proto.APLRotation{}
	player := &proto.Player{Rotation: custom}
	applyProtectionPaladinDefaultRotation(player)
	if player.Rotation != custom {
		t.Fatal("custom rotation was replaced")
	}
}
