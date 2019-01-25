package controller

import (
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"testing"
)

func TestUpsert(t *testing.T) {
	mem.Init()
	r := model.Route{ID: "1", DomainName: "test.com", Address: "1"}
	err := Upsert(r)
	if err != nil {
		t.Error(err)
	}

}
