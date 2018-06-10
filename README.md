[![Build Status](https://travis-ci.com/trigun117/sender-bot.svg?branch=master)](https://travis-ci.com/trigun117/sender-bot) [![codecov](https://codecov.io/gh/trigun117/sender-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/trigun117/sender-bot) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/sender-bot)](https://goreportcard.com/report/github.com/trigun117/sender-bot)

# sender-bot

This is a telegram bot which fetch json with proxies and send to channel

![example work of bot](https://github.com/trigun117/sender-bot/blob/master/image.JPG)

# Getting Started

For start using bot, build docker image from Dockerfile and run with this command
```
sudo docker run -d \
-e URL=site_url \
-e PASS=password \
-e FI=filed_name \
-e TO=sleep_time \
-e TN=bot_token \
-e CN=channel_name \
-e MTSV=mtproto_server \ 
-e MTP=mtproto_port \
-e MTSE=mtproto_secret \
-e MTL=mtproto_link \
-e ST=site \
--restart always \
image_name
```
# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details