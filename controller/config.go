package controller

import (
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-cloud/util"
)

func SaveConfig(item *model.Config) error {
	return mem.Push(item)
}

func GetPersistPath() string {
	c := GetConfig(model.PersistPath)
	v := c.Val
	v = util.Append(v, "\\config")
	return v
}

func GetConfig(id string) (r model.Config) {
	resp := make(chan model.Config)
	go func() {
		ite, err := mem.GetAll(model.ConfigCollectionName)
		iteratorToManyConfigs(ite, err, resp)
		close(resp)
	}()
	for i := range resp {
		if i.ID == id {
			t := i
			r = t
			return
		}
	}
	return
}

func iteratorToManyConfigs(iterator memdb.ResultIterator, err error, out chan model.Config) {
	iteratorToHandler(iterator, err, func(next interface{}) {
		item := next.(*model.Config)
		out <- *item
	})
}
