package util

import (
	"reflect"
)

//GetType returns the name of the struct passed to it as a string
func GetType(obj interface{}) (res string) {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	res = t.Name()
	return
}
