package controller

import (
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-docker/controller/pc-docker"
	pcmodel "github.com/plancks-cloud/plancks-docker/model"
	"github.com/sirupsen/logrus"
	"time"
)

var healthDoorbell = make(chan bool)

func StartHealthServer() {

	go func() {
		for {
			select {
			case <-healthDoorbell:
				getDesiredServiceStateAndFix()
			case <-time.After(60 * time.Second):
				getDesiredServiceStateAndFix()
			}
		}
	}()

}

func getDesiredServiceStateAndFix() {
	logrus.Println("⏰ Running health check")
	svcChan := GetAllServices()
	dockerServices, err := pc_docker.GetAllServiceStates()
	if err != nil {
		logrus.Println(err)
		return
	}
	checkExistingServices(svcChan, dockerServices)

}

func checkExistingServices(desiredServices chan *model.Service, foundServices []pcmodel.ServiceState) {
	count := 0
	for d := range desiredServices {
		count++
		found := false
		for _, s := range foundServices {
			if d.Name == s.Name {
				found = true
			}
		}
		if !found {
			err := pc_docker.CreateService(d)
			if err != nil {
				logrus.Println("Could not start docker service ", err)
			}
		}
	}
	logrus.Println("... health check tried to align ", count, " services")
}
