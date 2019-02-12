package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"testing"
)

var routeData2 = []byte(`
	[
		{
			"id": "1",
			"domainName": "bbc.com",
			"address": "bbc.com"
		},		
		{
			"id": "2",
			"domainName": "bbc2.com",
			"address": "bbc.com"
		}		
	]`)

func TestHandleApply(t *testing.T) {
	o := model.Object{
		Type: "bad-type",
	}
	err := HandleApply(&o)
	if err == nil {
		t.Error("A non-known type should throw an error")
	}

}

func TestRawToRoutes(t *testing.T) {
	raw := (*json.RawMessage)(&routeData2)
	routes, err := RawToRoutes(*raw)
	if err != nil {
		t.Error(err)
	}
	if len(*routes) != 2 {
		t.Error("Should have two routes!")
	}
	r := *routes
	if r[0].ID != "1" {
		t.Error("ID of route 1 should be 1")
	}

	if r[1].ID != "2" {
		t.Error("ID of route 2 should be 2")
	}
}

func TestRawToRoutesMemStore(t *testing.T) {
	mem.Init()

	raw := (*json.RawMessage)(&routeData2)
	r, err := RawToRoutes(*raw)
	if err != nil {
		t.Error(err)
	}
	err = InsertManyRoutes(r)
	if err != nil {
		t.Error(err)
	}
	routes := GetAllRoutesCopy()
	if len(routes) != 2 {
		t.Error("Should have two routes!")
	}

	found1 := false
	found2 := false

	for _, item := range routes {
		if item.ID == "1" {
			found1 = true
		} else if item.ID == "2" {
			found2 = true
		}
	}

	if !found1 {
		t.Error("Did not find route 1")
	}
	if !found2 {
		t.Error("Did not find route 2")
	}

}
