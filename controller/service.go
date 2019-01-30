package controller

import (
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
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
			logrus.Error(err)
			return err
		}
	}
	return
}

func DeleteManyServices(l *[]model.Service) (err error) {
	for _, item := range *l {
		_, err = mem.Delete(model.ServiceCollectionName, model.ServiceCollectionID, item.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return
}
