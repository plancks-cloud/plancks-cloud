package util

import (
	"strings"
)

func HostOfURL(url string) string {
	//TODO: proper handling with errors
	sl := strings.Split(url, ":")
	return sl[0]

}
