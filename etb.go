package main

import (
	"time"
)

func main() {
	etb := BasicBot{
		Channel:     "earentir",
		Name:        "duotronics",
		Port:        "6667",
		PrivatePath: "etb-auth.json",
		Server:      "irc.twitch.tv",
		MsgRate:     time.Duration(4000),
	}

	etb.Start()
}
