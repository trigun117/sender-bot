package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	link        = os.Getenv("URL")
	password    = os.Getenv("PASS")
	field       = os.Getenv("FI")
	t           = os.Getenv("TO")
	token       = os.Getenv("TN")
	channelName = os.Getenv("CN")
)

var p proxy

type proxy struct {
	Proxies []string
}

func fetchFreshProxy() (err error) {
	re, _ := http.PostForm(link, url.Values{field: {password}})
	b, _ := ioutil.ReadAll(re.Body)
	json.Unmarshal(b, &p)
	re.Body.Close()
	return
}

func main() {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("Invalid Token")
	}

	for {
		defer runtime.GC()
		fetchFreshProxy()
		pp := strings.Join(p.Proxies, "\n")
		msg := tgbotapi.NewMessageToChannel(channelName, fmt.Sprintf("Total socks5 proxies: %s\nProxies:\n%s", strconv.Itoa(len(p.Proxies)), pp))
		bot.Send(msg)
		to, _ := strconv.Atoi(t)
		d := time.Duration(to)
		time.Sleep(d * time.Second)
	}

}
