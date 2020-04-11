package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()
	discord.Token = loadTokenFromEnv()

	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(onReady)
	discord.AddHandler(messageCreate)

	err = discord.Open()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("login")
	<-make(chan bool)
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	sendMessage(session, "574884574778359844", "おはよう世界")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		sendMessage(s, m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		sendMessage(s, m.ChannelID, "Ping!")
	}
}

func sendMessage(session *discordgo.Session, channelID, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		panic(err)
	}
}

func loadTokenFromEnv() string {
	fp, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return "Bot " + token
}
