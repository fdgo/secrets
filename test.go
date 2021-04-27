package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func main() {
	Init()
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go regist(&wg)
	}
	wg.Wait()
}
func regist(wg *sync.WaitGroup) {
	defer wg.Done()
	p :=make(map[string]string)
	p["mail"] = "wangminghui0622@126.com"
	p["pwd"] = "123456"
	PostJson("http://127.0.0.1:8182/regist", p)
}
func Get(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

func Init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 3 * time.Minute,
			//MaxConnsPerHost: 10000,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
}
func PostJson(requestUrl string, params map[string]string) (int, string, error) {
	req := bytes.NewBuffer([]byte(ToJson(params)))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		fmt.Println(err, "**********")
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

func ToJson(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
