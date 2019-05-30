package controller

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

func iteratorToHandler(iterator memdb.ResultIterator, err error, handler func(next interface{})) {
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
		handler(next)
		count++
	}
	logrus.Debugln(fmt.Sprintf("iteratorToHandler counts: %d", count))
}
