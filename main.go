package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	// "runtime"
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
	mtserver    = os.Getenv("MTSV")
	mtport      = os.Getenv("MTP")
	mtsecret    = os.Getenv("MTSE")
	mtlink      = os.Getenv("MTL")
	site        = os.Getenv("ST")
)

var p proxy

type proxy struct {
	Proxies []string
}

func fetchFreshProxy() (err error) {
	re, _ := http.PostForm(link, url.Values{field: {password}})
	b, _ := ioutil.ReadAll(re.Body)
	defer re.Body.Close()
	json.Unmarshal(b, &p)
	return
}

func main() {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("Invalid Token")
	}

	for {
		// defer runtime.GC()
		fetchFreshProxy()
		pp := strings.Join(p.Proxies, "\n")
		msg := tgbotapi.NewMessageToChannel(channelName, fmt.Sprintf("MTProto Proxy\nServer: %s\nPort: %s\nSecret: %s\nLink: %s\n\nTotal socks5 proxies: %s\nProxies:\n%s\n\nShare with friends, contacts, social networks", mtserver, mtport, mtsecret, mtlink, strconv.Itoa(len(p.Proxies)), pp))
		kb := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Enable MTProto", mtlink), tgbotapi.NewInlineKeyboardButtonURL("Open site", site), tgbotapi.NewInlineKeyboardButtonURL("Open Bot", "https://t.me/tproxies_bot")))
		msg.ReplyMarkup = &kb
		bot.Send(msg)
		to, _ := strconv.Atoi(t)
		d := time.Duration(to)
		time.Sleep(d * time.Hour)
	}
}
