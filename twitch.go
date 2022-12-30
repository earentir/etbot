package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicklaw5/helix/v2"
)

func userExists(userName string) bool {
	if len(getTwitchUser(userName)) > 0 {
		return true
	} else {
		return false
	}
}

func getTwitchUser(userName string) TwitchUserData {
	var twitchUserData TwitchUserData

	client, err := helix.NewClient(&helix.Options{
		ClientID:        creds.TwitchClientID,
		UserAccessToken: creds.TwitchAppToken,
	})
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.GetUsers(&helix.UsersParams{
		Logins: []string{userName},
	})
	if err != nil {
		fmt.Println(err)
	}
	if resp.Error != "" {
		fmt.Println("Auth Error: ", resp.Error)
	}

	userdata := &resp.Data.Users
	jsonuserdata, err := json.Marshal(userdata)
	if err != nil {
		fmt.Println(err)
	}

	LoadJSONTOStruct(jsonuserdata, &twitchUserData)
	return twitchUserData
}

func getChannelInfo(loginID string) TwitchChannelData {
	var twitchChannelData TwitchChannelData

	client, err := helix.NewClient(&helix.Options{
		ClientID:        creds.TwitchClientID,
		UserAccessToken: creds.TwitchAppToken,
	})
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.GetChannelInformation(&helix.GetChannelInformationParams{
		BroadcasterID: loginID,
	})
	if err != nil {
		fmt.Println(err)
	}

	userdata := &resp.Data.Channels
	jsonchanneldata, err := json.Marshal(userdata)
	if err != nil {
		fmt.Println(err)
	}

	LoadJSONTOStruct(jsonchanneldata, &twitchChannelData)
	return twitchChannelData
}
