package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	LoadJSONFileTOStruct("etb-settings.json", &settings)

	etb := BasicBot{
		Channel:     settings.General.Twitch.Channel,
		Name:        settings.General.Twitch.BotUserName,
		Port:        settings.General.Twitch.IRCPort,
		PrivatePath: settings.General.CredentialFile,
		Server:      settings.General.Twitch.Server,
		MsgRate:     time.Duration(settings.General.Twitch.MSGRate),
	}

	LoadJSONFileTOStruct(etb.PrivatePath, &creds)

	//We can do stuff after

	// fmt.Println(SearchUser("mrpewpewlaser"))
	// fmt.Println(UserLevel("mrpewpewlaser"))
	// fmt.Println(getLevelCoolDown(UserLevel("mrpewpewlaser")))

	// fmt.Println(getUserSocials("sireeeki"))

	// for _, soc := range getUserSocials("earentir") {
	// 	fmt.Println(soc)
	// }

	// fmt.Println(CMDCanRun("wulgaru", "fr"))

	// if userExists("earentir") {
	// 	twusr := getTwitchUser("earentir")
	// 	fmt.Println(twusr[0].ID)
	// 	fmt.Println(twusr[0].Description)
	// 	fmt.Println(twusr[0].ViewCount)
	// 	fmt.Println(twusr[0].CreatedAt)
	// }

	// fmt.Println(getUserSocials("earentir"))

	if settings.Servers.BotServers.Chat {
		etb.Start()
	}
}

func LoadJSONFileTOStruct(jsonFileName string, onTo interface{}) {
	//read json here
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if nil != err {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonFile), &onTo)
}

func LoadJSONTOStruct(jsondata []byte, onTo interface{}) {
	//read json here
	json.Unmarshal(jsondata, &onTo)
}
