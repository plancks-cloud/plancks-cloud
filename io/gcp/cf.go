package gcp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetCollection(endpoint, id, key, collection string) (body []byte, err error) {
	url := fmt.Sprint(endpoint, "?action=get&collection=", collection)
	body, err = doHttpCall(http.MethodGet, url, getHeaders(id, key), nil)
	return
}

func SetCollection(endpoint, id, key, collection string, val []byte) (body []byte, err error) {
	url := fmt.Sprint(endpoint, "?action=set&collection=", collection)
	body, err = doHttpCall(http.MethodPost, url, getHeaders(id, key), val)
	return
}

func doHttpCall(method, url string, headers map[string]string, v []byte) (body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func getHeaders(id, key string) (m map[string]string) {
	m = make(map[string]string)
	m["persist-id"] = id
	m["persist-key"] = key
	return
}
