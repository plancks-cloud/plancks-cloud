package controller

import (
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"reflect"
	"testing"
)

func TestGetAllRoutesCopy(t *testing.T) {
	mem.Init()
	r := model.Route{ID: "1", DomainName: "test.com", Address: "1"}
	err := Upsert(&r)
	if err != nil {
		t.Error(err)
	}

	arr := GetAllRoutesCopy()
	if len(arr) != 1 {
		t.Error("Could not find inserted route")

	}

	ok := reflect.DeepEqual(arr[0], r)
	if !ok {
		t.Error("Stored route not the same as found route")
	}

}

func TestGetAllRoutesCopyTwoItems(t *testing.T) {
	mem.Init()
	r := model.Route{ID: "1", DomainName: "test.com", Address: "1"}
	err := Upsert(&r)
	if err != nil {
		t.Error(err)
	}

	r = model.Route{ID: "2", DomainName: "test2.com", Address: "2"}
	err = Upsert(&r)
	if err != nil {
		t.Error(err)
	}

	arr := GetAllRoutesCopy()
	if len(arr) != 2 {
		t.Error("Could not find inserted route")

	}

	ok := reflect.DeepEqual(arr[0], r)
	if !ok {
		t.Error("Stored route not the same as found route")
	}

}

func TestGetAllRoutesCopyTwoItemsBatch(t *testing.T) {
	mem.Init()
	var arr []model.Route
	r1 := model.Route{ID: "1", DomainName: "test.com", Address: "1"}
	arr = append(arr, r1)

	r2 := model.Route{ID: "2", DomainName: "test2.com", Address: "2"}
	arr = append(arr, r2)

	err := InsertManyRoutes(&arr)
	if err != nil {
		t.Error(err)
	}

	res := GetAllRoutesCopy()
	if len(res) != 2 {
		t.Error("Could not find inserted route")

	}

	ok := reflect.DeepEqual(arr[0], r1)
	if !ok {
		t.Error("Stored route not the same as found route")
	}
	ok = reflect.DeepEqual(arr[1], r2)
	if !ok {
		t.Error("Stored route not the same as found route")
	}

}
