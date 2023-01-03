package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/nicklaw5/helix/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

type GoalsResponse struct {
	Data []struct {
		BroadcasterID    string `json:"broadcaster_id"`
		BroadcasterLogin string `json:"broadcaster_login"`
		BroadcasterName  string `json:"broadcaster_name"`
		CreatedAt        string `json:"created_at"`
		CurrentAmount    int64  `json:"current_amount"`
		Description      string `json:"description"`
		ID               string `json:"id"`
		TargetAmount     int64  `json:"target_amount"`
		Type             string `json:"type"`
	} `json:"data"`
}

func GetGoals(client *http.Client, channel string) (GoalsResponse, error) {
	var baseURL = "https://api.twitch.tv/helix/"
	// Set up the API request
	req, err := http.NewRequest("GET", baseURL+"goals?broadcaster_id="+getTwitchUser(channel)[0].ID, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Client-ID", creds.TwitchClientID)
	req.Header.Set("Authorization", "Bearer "+creds.TwitchAppToken)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Read the response and unmarshal it into a slice of Goal objects
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response GoalsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response, nil
}

func showgoals(channel string) string {
	var goalMessage string
	client := &http.Client{}

	goals, err := GetGoals(client, channel)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, goal := range goals.Data {
			goalMessage = channel + "'s Current Goal was created on " + goal.CreatedAt + " and is a " + goal.Type + " goal. " + goal.Description + " Its currently at " + fmt.Sprint(goal.CurrentAmount) + " out of " + fmt.Sprint(goal.TargetAmount) + "!"
		}
	}

	return goalMessage
}

func subscriberCount(channel string) int {
	var subscriberCount int
	client, err := helix.NewClient(&helix.Options{
		ClientID:        creds.TwitchClientID,
		UserAccessToken: creds.TwitchAppToken,
	})
	if err != nil {
		fmt.Println(err)
	}

	channel = strings.ToLower(channel)
	channel = strings.TrimSpace(channel)

	resp, err := client.GetSubscriptions(&helix.SubscriptionsParams{
		BroadcasterID: getTwitchUser(channel)[0].ID,
	})
	if err != nil {
		fmt.Println(err)
	}

	subscriberCount = resp.Data.Total
	return subscriberCount
}

func issubbed(user, channel string) bool {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        creds.TwitchClientID,
		UserAccessToken: creds.TwitchAppToken,
	})
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.CheckUserSubscription(&helix.UserSubscriptionsParams{
		BroadcasterID: string(getTwitchUser(channel)[0].ID),
		UserID:        string(getTwitchUser(user)[0].ID),
	})
	if err != nil {
		fmt.Println(err)
	}

	var issubbed bool
	// resp.Data.UserSubscriptions[0].IsGift
	if resp.Data.UserSubscriptions[0].Tier != "" {
		issubbed = true
	} else {
		issubbed = false
	}

	return issubbed
}

func cmdGoals(bb *BasicBot, cmd, userName, msg string) {
	var goals string = showgoals(settings.General.Twitch.Channel)

	if isAttr(msg) {
		botSay(bb, goals)
	} else {
		botSay(bb, goals)
	}
}

func cmdSubs(bb *BasicBot, cmd, userName, msg string) {
	var channel string = settings.General.Twitch.Channel
	var subs int = subscriberCount(channel)

	botSay(bb, cases.Title(language.Und).String(channel)+" has "+strconv.Itoa(subs)+" subscribers.")
}
