package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/textproto"
	"time"
)

type Bot interface {
	// Opens a connection to the Twitch.tv IRC chat server.
	Connect()

	// Closes a connection to the Twitch.tv IRC chat server.
	Disconnect()

	// Listens to chat messages and PING request from the IRC server.
	HandleChat() error

	// Joins a specific chat channel.
	JoinChannel()

	// Parses credentials needed for authentication.
	ReadCredentials() error

	// Sends a message to the connected channel.
	Say(msg string) error

	// Attempts to keep the bot connected and handling chat.
	Start()
}

// Connects the bot to the Twitch IRC server. The bot will continue to try to connect until it
// succeeds or is forcefully shutdown.
func (bb *BasicBot) Connect() {
	var err error
	CPrint("y", fmt.Sprintf("[%s] Connecting to %s...\n", timeStamp(), bb.Server))

	// makes connection to Twitch IRC server
	bb.conn, err = net.Dial("tcp", bb.Server+":"+bb.Port)
	if nil != err {
		CPrint("y", fmt.Sprintf("[%s] Cannot connect to %s, retrying.\n", timeStamp(), bb.Server))
		bb.Connect()
		return
	}
	CPrint("y", fmt.Sprintf("[%s] Connected to %s!\n", timeStamp(), bb.Server))
	bb.startTime = time.Now()
}

// Officially disconnects the bot from the Twitch IRC server.
func (bb *BasicBot) Disconnect() {
	bb.conn.Close()
	upTime := time.Since(time.Now()) // time.Now().Sub(bb.startTime).Seconds()
	CPrint("y", fmt.Sprintf("[%s] Closed connection from %s! | Live for: %v\n", timeStamp(), bb.Server, upTime))
}

// Listens for and logs messages from chat. Responds to commands from the channel owner. The bot
// continues until it gets disconnected, told to shutdown, or forcefully shutdown.
func (bb *BasicBot) HandleChat() error {
	CPrint("y", fmt.Sprintf("[%s] Watching #%s...\n", timeStamp(), bb.Channel))

	// reads from connection
	tp := textproto.NewReader(bufio.NewReader(bb.conn))

	// listens for chat messages
	for {
		line, err := tp.ReadLine()
		// rawLine := line

		if nil != err {
			// officially disconnects the bot from the server
			bb.Disconnect()
			return errors.New("bb.Bot.HandleChat: Failed to read line from channel. Disconnected")
		}

		// logs the response from the IRC server
		CPrint("y", fmt.Sprintf("[%s] %s\n", timeStamp(), line))

		if line == "PING :tmi.twitch.tv" {
			// respond to PING message with a PONG message, to maintain the connection
			bb.conn.Write([]byte("PONG :tmi.twitch.tv\r\n"))
			continue
		} else {
			// handle a PRIVMSG message
			matches := MsgRegex.FindStringSubmatch(line)
			if nil != matches {
				userName := matches[1]
				msgType := matches[2]
				msg := matches[3]

				cmdUnlurk(bb, userName)
				ParseCommand(bb, msgType, msg, userName)
			} else {
				// fmt.Println(rawLine)
				fmt.Print("")
			}
			time.Sleep(bb.MsgRate)
		}
	}
}

// Makes the bot join its pre-specified channel.
func (bb *BasicBot) JoinChannel() {
	CPrint("y", fmt.Sprintf("[%s] Joining #%s...\n", timeStamp(), bb.Channel))
	bb.conn.Write([]byte("PASS " + creds.TwitchPassword + "\r\n"))
	bb.conn.Write([]byte("NICK " + bb.Name + "\r\n"))
	bb.conn.Write([]byte("JOIN #" + bb.Channel + "\r\n"))

	CPrint("y", fmt.Sprintf("[%s] Joined #%s as @%s!\n", timeStamp(), bb.Channel, bb.Name))
}

// Makes the bot send a message to the chat channel.
func (bb *BasicBot) Say(msg string) error {
	if msg == "" {
		return errors.New("BasicBot.Say: msg was empty")
	}

	// check if message is too large for IRC
	if len(msg) > 512 {
		return errors.New("BasicBot.Say: msg exceeded 512 bytes")
	}

	_, err := bb.conn.Write([]byte(fmt.Sprintf("PRIVMSG #%s :%s\r\n", bb.Channel, msg)))
	if nil != err {
		return err
	}
	return nil
}

// Starts a loop where the bot will attempt to connect to the Twitch IRC server, then connect to the
// pre-specified channel, and then handle the chat. It will attempt to reconnect until it is told to
// shut down, or is forcefully shutdown.
func (bb *BasicBot) Start() {
	for {
		bb.Connect()
		bb.JoinChannel()
		err := bb.HandleChat()
		if nil != err {
			// attempts to reconnect upon unexpected chat error
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(err)
			fmt.Println("Starting bot again...")
		} else {
			return
		}
	}
}
