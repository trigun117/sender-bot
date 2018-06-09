package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var proxy Proxy

var (
	link     = os.Getenv("URL")
	field    = os.Getenv("FI")
	password = os.Getenv("PASS")
)

type Proxy struct {
	Proxies []string
}

func fetchFreshProxy() (err error) {
	response, _ := http.PostForm(link, url.Values{field: {password}})
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(body, &proxy)
	return
}
