package controller

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

func SaveConfig(item *model.Config) error {
	return mem.Push(item)
}

func GetPersistPath() string {
	c := GetConfig(model.PersistPath)
	v := c.Val
	fmt.Println(v)
	return fmt.Sprint(v, "\\config")
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
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	if iterator == nil {
		return
	}
	more := true
	count := 0
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(*model.Config)
		out <- *item
		count++
	}
	logrus.Debugln(fmt.Sprintf("Route iterator counts: %d", count))
}
