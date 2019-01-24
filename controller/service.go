package controller

import (
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
)

func GetAllServices() (resp chan *model.Service) {
	resp = make(chan *model.Service)
	go func() {
		ite, err := mem.GetAll(model.ServiceCollectionName)
		iteratorToManyServices(ite, err, resp)
		close(resp)
	}()
	return resp

}

func iteratorToManyServices(iterator memdb.ResultIterator, err error, out chan *model.Service) {
	c := mem.IteratorToChannel(iterator, err)
	for i := range c {
		item := i.(*model.Service)
		out <- item
	}

}

func InsertManyServices(l *[]model.Service) (err error) {
	for _, item := range *l {
		err = Upsert(&item)
		if err != nil {
			return err
		}
	}
	return
}
