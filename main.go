package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var roles []discordgo.Role

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
	<-sc
	discord.Close()
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	sendMessage(session, "574884574778359844", "おはよう世界")
	//sendMessage(session, "690909527461199922", "おはよう世界")
	session.UpdateStatus(1, "Goのべんつよ、たのしいね")
	guilds := ready.Guilds
	if len(guilds) != 1 {
		panic("Can't boot 2 or more bot with same token")
	}
	guild := guilds[0]
	roleAddress := guild.Roles
	roleListUpdate(roleAddress)
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

func roleListUpdate(roleList []*discordgo.Role) {
	roles = []discordgo.Role{}
	for _, role := range roleList {
		roles = append(roles, *role)
	}
}

func createNewRole(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	guild, err := session.State.Guild(event.GuildRole.GuildID)
	if err != nil {
		return
	}
	roleAddress := guild.Roles
	roleListUpdate(roleAddress)
}

func updateRole(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {
	newRole := event.GuildRole.Role
	if alreadyRoleExists(newRole.Name) {
		return
	}
	message := "新しいクソRole***†" + newRole.Name + "†***が追加されたよ"

	sendMessage(session, "574884574778359844", message)
	//sendMessage(session, "690909527461199922", message)
}

func alreadyRoleExists(roleName string) bool {
	for _, value := range roles {
		if value.Name == roleName {
			return true
		}
	}
	return false
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
