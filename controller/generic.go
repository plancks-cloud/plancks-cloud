package controller

import (
	"github.com/plancks-cloud/plancks-cloud/io/mem"
)

func Upsert(obj interface{}) error {
	return mem.Push(obj)
}
