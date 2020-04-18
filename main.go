package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"isso/go/firstbot/handler"

	"github.com/bwmarrin/discordgo"
)

var roles []discordgo.Role

// テスト環境では574884574778359844
// 限界開発鯖では690909527461199922
const rootChannelID string = "574884574778359844"

func main() {
	discord, err := discordgo.New()
	discord.Token = loadTokenFromEnv()

	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)
	discord.AddHandler(createNewRole)
	discord.AddHandler(updateRole)

	err = discord.Open()
	defer discord.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("login")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	handler.OnMessageCreate(s, m)
}

func createNewRole(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	handler.CreateNewRole(session, event, &roles)
}

func updateRole(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	handler.UpdateRole(s, e, roles, rootChannelID)
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	handler.OnReadyMessageSend(session, ready, rootChannelID, &roles)
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
