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

func main() {
	discord, err := discordgo.New()
	discord.Token = loadTokenFromEnv()

	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(onReady)
	discord.AddHandler(handler.OnReadyMessageSend)
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

func createNewRole(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	handler.CreateNewRole(session, event, &roles)
}

func updateRole(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	handler.UpdateRole(s, e, roles)
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	handler.OnReadyMessageSend(session, ready, "574884574778359844", &roles)
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
