package test

import (
	"reflect"
	"testing"

	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateSpirit(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	spirit := &spiritsv1.Spirit{
		Name: stringPtr("some-name"),
		Stats: &spiritsv1.SpiritStats{
			Health:               int64Ptr(1),
			PhysicalPower:        int64Ptr(2),
			PhysicalConstitution: int64Ptr(3),
			MentalPower:          int64Ptr(4),
			MentalConstitution:   int64Ptr(5),
			Agility:              int64Ptr(6),
		},
	}
	rsp, err := clients.spirit.CreateSpirit(state.ctx, &spiritsv1.CreateSpiritRequest{Spirit: spirit})
	if err != nil {
		t.Fatal("create spirit", err)
	}
	if len(rsp.GetSpirit().GetMeta().GetId()) == 0 {
		t.Error("got empty id")
	}
	if rsp.GetSpirit().GetMeta().GetCreatedTime().AsTime().IsZero() {
		t.Error("got empty created time")
	}
	if rsp.GetSpirit().GetMeta().GetUpdatedTime().AsTime().IsZero() {
		t.Error("got empty updated time")
	}
	if created, updated := rsp.GetSpirit().GetMeta().GetCreatedTime().AsTime(), rsp.GetSpirit().GetMeta().GetUpdatedTime().AsTime(); created != updated {
		t.Errorf("created/updated time mismatch: %s/%s", created, updated)
	}
	if diff := cmp.Diff(spirit, noMeta(rsp.GetSpirit()), protocmp.Transform()); len(diff) > 0 {
		t.Errorf("unexpected spirit, -want, +got:\n%s", diff)
	}
}

func TestCreateSpiritInvalidArgument(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	spirit := &spiritsv1.Spirit{
		Name: stringPtr("some-name"),
		Stats: &spiritsv1.SpiritStats{
			Health:  int64Ptr(1),
			Agility: int64Ptr(1),
		},
		Actions: []*spiritsv1.SpiritAction{
			{
				Name: stringPtr("tuna"),
				Definition: &spiritsv1.SpiritAction_Inline{
					Inline: &spiritsv1.Action{Definition: &spiritsv1.Action_Script{Script: ""}},
				},
			},
			{
				Name: stringPtr("tuna"),
				Definition: &spiritsv1.SpiritAction_Inline{
					Inline: &spiritsv1.Action{Definition: &spiritsv1.Action_Script{Script: ""}},
				},
			},
		},
	}
	_, err := clients.spirit.CreateSpirit(state.ctx, &spiritsv1.CreateSpiritRequest{Spirit: spirit})
	if want, got := status.New(codes.InvalidArgument, "duplicate action name: tuna"), status.Convert(err); !reflect.DeepEqual(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestUpdateSpirit(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	spirit := &spiritsv1.Spirit{
		Name: stringPtr("some-name"),
		Stats: &spiritsv1.SpiritStats{
			Health:               int64Ptr(1),
			PhysicalPower:        int64Ptr(2),
			PhysicalConstitution: int64Ptr(3),
			MentalPower:          int64Ptr(4),
			MentalConstitution:   int64Ptr(5),
			Agility:              int64Ptr(6),
		},
	}
	createRsp, err := clients.spirit.CreateSpirit(state.ctx, &spiritsv1.CreateSpiritRequest{Spirit: spirit})
	if err != nil {
		t.Fatal("update spirit", err)
	}

	createRsp.GetSpirit().Name = stringPtr("some-other-name")
	spirit.Name = createRsp.GetSpirit().Name
	*createRsp.GetSpirit().GetStats().Agility += 1
	spirit.Stats.Agility = createRsp.GetSpirit().GetStats().Agility
	updateRsp, err := clients.spirit.UpdateSpirit(state.ctx, &spiritsv1.UpdateSpiritRequest{Spirit: createRsp.GetSpirit()})
	if err != nil {
		t.Fatal("update spirit", err)
	}
	if created, updated := updateRsp.GetSpirit().GetMeta().GetCreatedTime().AsTime(), updateRsp.GetSpirit().GetMeta().GetUpdatedTime().AsTime(); created.After(updated) {
		t.Errorf("created time is after updated time: %s/%s", created, updated)
	}
	if diff := cmp.Diff(spirit, noMeta(updateRsp.GetSpirit()), protocmp.Transform()); len(diff) > 0 {
		t.Errorf("unexpected spirit, -want, +got:\n%s", diff)
	}
}

func TestUpdateSpiritNotFound(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	spirit := &spiritsv1.Spirit{Meta: &spiritsv1.Meta{Id: stringPtr("tuna")}, Name: stringPtr("aaa")}
	_, err := clients.spirit.UpdateSpirit(state.ctx, &spiritsv1.UpdateSpiritRequest{Spirit: spirit})
	if want, got := status.New(codes.NotFound, "not found"), status.Convert(err); !reflect.DeepEqual(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetSpiritNotFound(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	_, err := clients.spirit.GetSpirit(state.ctx, &spiritsv1.GetSpiritRequest{Id: stringPtr("foo")})
	if want, got := status.New(codes.NotFound, "not found"), status.Convert(err); !reflect.DeepEqual(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestInvalidSpiritName(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	_, err := clients.spirit.CreateSpirit(state.ctx, &spiritsv1.CreateSpiritRequest{Spirit: &spiritsv1.Spirit{}})
	s, ok := status.FromError(err)
	if !ok {
		t.Fatal("wanted error")
	}
	if want, got := codes.InvalidArgument, s.Code(); want != got {
		t.Errorf("wanted code %d, got %d (%+v)", want, got, err)
	}
}

func TestBuiltin(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	rsp, err := clients.spirit.ListSpirits(state.ctx, &spiritsv1.ListSpiritsRequest{
		Name: stringPtr("i"),
	})
	if err != nil {
		t.Fatal("get spirit", err)
	}

	if want, got := 1, len(rsp.Spirits); want != got {
		t.Errorf("wanted %d spirit, got %d", want, got)
	}
}

func noMeta(spirit *spiritsv1.Spirit) *spiritsv1.Spirit {
	spirit = proto.Clone(spirit).(*spiritsv1.Spirit)
	spirit.Meta = nil
	return spirit
}

func stringPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64    { return &i }
func intelligencePtr(intelligence spiritsv1.BattleTeamSpiritIntelligence) *spiritsv1.BattleTeamSpiritIntelligence {
	return &intelligence
}
