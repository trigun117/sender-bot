package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	sleepTime   = os.Getenv("TO")
	token       = os.Getenv("TN")
	channelName = os.Getenv("CN")
	mtserver    = os.Getenv("MTSV")
	mtport      = os.Getenv("MTP")
	mtsecret    = os.Getenv("MTSE")
	mtlink      = os.Getenv("MTL")
	site        = os.Getenv("ST")
)

func bot() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("Invalid Token")
	}

	for {
		fetchFreshProxy()
		joinedProxies := strings.Join(proxy.Proxies, "\n")

		message := tgbotapi.NewMessageToChannel(channelName, fmt.Sprintf("MTProto Proxy\nServer: %s\nPort: %s\nSecret: %s\nLink: %s\n\nTotal socks5 proxies: %s\nProxies:\n%s\n\nShare with friends, contacts, social networks", mtserver, mtport, mtsecret, mtlink, strconv.Itoa(len(proxy.Proxies)), joinedProxies))

		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Enable MTProto", mtlink), tgbotapi.NewInlineKeyboardButtonURL("Open site", site), tgbotapi.NewInlineKeyboardButtonURL("Open Bot", "https://t.me/tproxies_bot")))

		message.ReplyMarkup = &keyboard
		bot.Send(message)

		runtime.GC()

		st, _ := strconv.Atoi(sleepTime)
		duration := time.Duration(st)
		time.Sleep(duration * time.Hour)
	}
}
