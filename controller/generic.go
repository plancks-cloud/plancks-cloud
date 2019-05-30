package controller

import (
	"github.com/hashicorp/go-memdb"
)

func iteratorToHandler(iterator memdb.ResultIterator, handler func(next interface{})) {
	if iterator == nil {
		return
	}
	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue // SHOULD THIS BE RETURN?
		}
		handler(next)
	}
}
