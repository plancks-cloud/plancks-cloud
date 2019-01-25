package controller

import (
	"github.com/plancks-cloud/plancks-cloud/model"
	"testing"
)

func TestHandleApply(t *testing.T) {
	o := model.Object{
		Type: "bad-type",
	}
	err := HandleApply(&o)
	if err == nil {
		t.Error("A non-known type should throw an error")
	}

}
