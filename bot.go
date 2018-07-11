package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/trigun117/sender-bot/fetcher"
	"os"
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
	mtlink      = os.Getenv("MTLL")
	mtlink1     = os.Getenv("MTL")
	site        = os.Getenv("ST")
)

func bot() {
	// creating bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("Invalid Token")
	}

	st, _ := strconv.Atoi(sleepTime)
	duration := time.Duration(st)

	for t := time.Tick(duration * time.Hour); ; <-t {

		fetcher.FetchFreshProxy()
		joinedProxies := strings.Join(fetcher.Proxy.Proxies, "\n")

		// creating message with text and proxies
		message := tgbotapi.NewMessageToChannel(channelName, fmt.Sprintf("MTProto Proxy\nServer: %s\nPort: %s\nSecret: %s\nLink: %s\n\nTotal socks5 proxies: %s\nProxies:\n%s\n\nShare with friends, contacts, social networks", mtserver, mtport, mtsecret, mtlink, strconv.Itoa(len(fetcher.Proxy.Proxies)), joinedProxies))

		// creating inline keyboard
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Enable MTProto \xE2\x93\x82", mtlink1)), tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Open site \xF0\x9F\x8C\x90", site)), tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Open Bot \xF0\x9F\xA4\x96", "https://t.me/tproxies_bot")))

		message.ReplyMarkup = &keyboard
		bot.Send(message)
	}
}
