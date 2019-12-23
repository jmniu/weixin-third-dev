package wxopenapi

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const TIME_OUT = 360

func InitNet() {
	if http.DefaultClient.Transport == nil {
		t := &http.Transport{}
		http.DefaultClient.Transport = t
		t.Dial = func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(TIME_OUT * time.Second)
			d, err := net.DialTimeout(netw, addr, TIME_OUT*time.Second)
			if err != nil {
				return nil, err
			}
			d.SetDeadline(deadline)
			return d, nil
		}
	}
	if t2, ok := http.DefaultClient.Transport.(*http.Transport); ok {
		t2.DisableKeepAlives = true
	}
}

func PostJsonString(uri string, param string) ([]byte, error) {
	InitNet()
	rsp, err := http.Post(uri, "application/json", strings.NewReader(param))
	if err != nil {
		return nil, err
	}
	defer func() {
		if rsp != nil && rsp.Body != nil {
			rsp.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(rsp.Body)
	return body, err
}

func PostJsonByte(uri string, param []byte) ([]byte, error) {
	InitNet()
	rsp, err := http.Post(uri, "application/json", bytes.NewReader(param))
	if err != nil {
		return nil, err
	}
	defer func() {
		if rsp != nil && rsp.Body != nil {
			rsp.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(rsp.Body)
	return body, err
}
