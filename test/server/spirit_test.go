package test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateSpirit(t *testing.T) {
	clients := startServer(t)

	spirit := &api.Spirit{
		Name: "some-name",
		Stats: &api.SpiritStats{
			Health:  1,
			Agility: 2,
		},
	}
	rsp, err := clients.spirit.CreateSpirit(context.Background(), &api.CreateSpiritRequest{Spirit: spirit})
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

func TestUpdateSpirit(t *testing.T) {
	clients := startServer(t)

	spirit := &api.Spirit{
		Name: "some-name",
		Stats: &api.SpiritStats{
			Health:  1,
			Agility: 2,
		},
	}
	createRsp, err := clients.spirit.CreateSpirit(context.Background(), &api.CreateSpiritRequest{Spirit: spirit})
	if err != nil {
		t.Error("create spirit", err)
	}

	createRsp.GetSpirit().Name = "some-other-name"
	spirit.Name = createRsp.GetSpirit().Name
	createRsp.GetSpirit().GetStats().Agility += 1
	spirit.Stats.Agility = createRsp.GetSpirit().GetStats().Agility
	updateRsp, err := clients.spirit.UpdateSpirit(context.Background(), &api.UpdateSpiritRequest{Spirit: createRsp.GetSpirit()})
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
	clients := startServer(t)

	spirit := &api.Spirit{Meta: &api.Meta{Id: "tuna"}}
	_, err := clients.spirit.UpdateSpirit(context.Background(), &api.UpdateSpiritRequest{Spirit: spirit})
	if want, got := status.New(codes.NotFound, "not found"), status.Convert(err); !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetSpiritNotFound(t *testing.T) {
	clients := startServer(t)

	_, err := clients.spirit.GetSpirit(context.Background(), &api.GetSpiritRequest{Id: "foo"})
	if want, got := status.New(codes.NotFound, "not found"), status.Convert(err); !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestBuiltin(t *testing.T) {
	clients := startServer(t)

	_, err := clients.spirit.GetSpirit(context.Background(), &api.GetSpiritRequest{Id: "i"})
	if err != nil {
		t.Fatal("get spirit", err)
	}
}

func noMeta(spirit *api.Spirit) *api.Spirit {
	spirit = proto.Clone(spirit).(*api.Spirit)
	spirit.Meta = nil
	return spirit
}
