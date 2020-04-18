package handler

import (
	"isso/go/firstbot/handler/functions"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		return
	}

	if !strings.HasPrefix(m.Content, "!") {
		return
	}
	command := strings.Split(m.Content, " ")[0][1:]

	switch command {
	case "dice":
		functions.DiceExecute(s, m)
	}
}
