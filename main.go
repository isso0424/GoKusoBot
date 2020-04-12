package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var roles []*discordgo.Role

var tmp int

func main() {
	discord, err := discordgo.New()
	discord.Token = loadTokenFromEnv()

	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(onReady)
	discord.AddHandler(messageCreate)
	discord.AddHandler(createNewRole)
	discord.AddHandler(updateRole)

	err = discord.Open()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("login")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	tt := time.NewTicker(50 * time.Second)
	loop := true

	for loop {
		select {
		case <-sc:
			loop = false
			break
		case <-tt.C:
			tmp = 0
		}
	}

	discord.Close()
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	//sendMessage(session, "574884574778359844", "おはよう世界")
	sendMessage(session, "690909527461199922", "おはよう世界")
	session.UpdateStatus(1, "Goのべんつよ、たのしいね")
	guilds := ready.Guilds
	if len(guilds) != 1 {
		panic("Can't boot 2 or more bot with same token")
	}
	guild := guilds[0]
	roles = guild.Roles
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		sendMessage(s, m.ChannelID, "Pong!")
		return
	}

	if m.Content == "pong" {
		sendMessage(s, m.ChannelID, "Ping!")
		return
	}

	if strings.HasPrefix(m.Content, "!update") {
		if len(m.Content) <= 8 {
			return
		}
		messageWithOutPrefix := m.Content[8:]
		s.UpdateStatus(0, messageWithOutPrefix)
		return
	}

	if strings.HasPrefix(m.Content, "!status") {
		if len(m.Content) <= 8 {
			return
		}
		messageWithOutPrefix := m.Content[8:9]
		i, err := strconv.Atoi(messageWithOutPrefix)
		if err != nil {
			sendMessage(s, m.ChannelID, "数値を入力してね")
		}
		s.UpdateStatus(i, "ステータスをアップデートしました")
	}
}

func createNewRole(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	guild, err := session.State.Guild(event.GuildRole.GuildID)
	if err != nil {
		return
	}
	roles = guild.Roles
	tmp = len(roles) - 1
}

func updateRole(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {
	fmt.Println(tmp)
	if tmp != 0 {
		tmp--
	}
	if tmp != 0 {
		return
	}
	newRoleName := event.GuildRole.Role.Name
	fmt.Println(event.GuildRole.Role)
	message := "新しいクソRole***†" + newRoleName + "†***が追加されたよ"

	//sendMessage(session, "574884574778359844", message)
	sendMessage(session, "690909527461199922", message)
}

func sendMessage(session *discordgo.Session, channelID, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTokenFromEnv() string {
	fp, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return "Bot " + token
}
