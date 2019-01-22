package mem

import (
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

//IteratorToChannel reads an iterator into a channel
func IteratorToChannel(iterator memdb.ResultIterator, err error) chan interface{} {
	c := make(chan interface{})
	go func(c chan interface{}) {
		if err != nil {
			log.Error(err.Error())
			close(c)
			return
		}
		if iterator == nil {
			close(c)
			return
		}
		more := true
		for more {
			next := iterator.Next()
			if next == nil {
				more = false
				continue
			}
			c <- next
		}
		close(c)
	}(c)
	return c
}
