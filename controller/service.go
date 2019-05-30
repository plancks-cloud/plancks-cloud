package controller

import (
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	pc_model "github.com/plancks-cloud/plancks-docker/model"
	"github.com/sirupsen/logrus"
)

func GetAllServices() (resp chan *model.Service) {
	resp = make(chan *model.Service)
	go func() {
		ite, err := mem.GetAll(model.ServiceCollectionName)
		if err != nil {
			logrus.Errorln(err)
			close(resp)
			return
		}
		iteratorToManyServices(ite, resp)
		close(resp)
	}()
	return resp

}

func GetAllServicesCopy() []model.Service {
	var arr []model.Service
	for item := range GetAllServices() {
		arr = append(arr, *item)
	}
	return arr
}

func iteratorToManyServices(iterator memdb.ResultIterator, out chan *model.Service) {
	iteratorToHandler(iterator, func(next interface{}) {
		item := next.(*model.Service)
		out <- item
	})
}

func InsertManyServices(l *[]model.Service) (err error) {
	for _, item := range *l {
		itemN := item
		err = mem.Push(&itemN)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	syncRoutesToDisk()
	return
}

func DeleteManyServices(l *[]model.Service) (err error) {
	for _, item := range *l {
		_, err = mem.Delete(model.ServiceCollectionName, model.CollectionID, item.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	syncRoutesToDisk()
	return
}

func convertServices(services *[]model.Service) (res []pc_model.ServiceState) {
	res = make([]pc_model.ServiceState, len(*services))
	for _, service := range *services {
		item := pc_model.ServiceState{
			ID:          service.ID,
			Name:        service.Name,
			MemoryLimit: service.MemoryLimit,
			Image:       service.Image,
		}
		res = append(res, item)
	}
	return
}
