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
