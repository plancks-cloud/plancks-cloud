package controller

import (
	"github.com/hashicorp/go-memdb"
)

func iteratorToHandler(iterator memdb.ResultIterator, handler func(next interface{})) {
	if iterator == nil {
		return
	}
	for {
		next := iterator.Next()
		if next == nil {
			return
		}
		handler(next)
	}
}
