package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpGet(rawUrl string, kvp map[string]string, kvh map[string]string) ([]byte, error) {
	var extra string
	for k, v := range kvp {
		extra = extra + "&" + k + "=" + v
	}
	if extra != "" {
		rawUrl = rawUrl + "?" + extra
	}
	fmt.Println(rawUrl)
	req, err := http.NewRequest("GET", rawUrl, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for k, v := range kvh {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rsp.Body.Close()
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
