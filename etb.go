package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Credentials struct {
	TwitchPassword    string `json:"twitch_password,omitempty"`
	TwitchClientID    string `json:"twitch_client_id,omitempty"`
	OpenWeatherAPIKey string `json:"openweathermapapi,omitempty"`
}

var creds Credentials

var settings Settings

func main() {

	LoadJSONTOStruct("etb-settings.json", &settings)

	etb := BasicBot{
		Channel:     "earentir",
		Name:        "duotronics",
		Port:        "6667",
		PrivatePath: "etb-auth.json",
		Server:      "irc.twitch.tv",
		MsgRate:     time.Duration(4000),
	}

	LoadJSONTOStruct(etb.PrivatePath, &creds)

	fmt.Println(etb.PrivatePath)
	fmt.Println(creds.OpenWeatherAPIKey)

	etb.Start()
}

func LoadJSONTOStruct(jsonFileName string, onTo interface{}) {
	//read json here
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if nil != err {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonFile), &onTo)
}
