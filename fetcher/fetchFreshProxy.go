package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Proxy contain unmarshaled proxies
var Proxy ProxyS

var (
	link     = os.Getenv("URL")
	field    = os.Getenv("FI")
	password = os.Getenv("PASS")
)

// ProxyS struct which contain proxies
type ProxyS struct {
	Proxies []string
}

// FetchFreshProxy fetching json with proxies and unmarshal it to struct
func FetchFreshProxy() (err error) {
	response, _ := http.PostForm(link, url.Values{field: {password}})
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(body, &Proxy)
	return
}
