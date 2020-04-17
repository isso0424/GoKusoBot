package handler

import "github.com/bwmarrin/discordgo"

func OnReadyMessageSend(session *discordgo.Session, ready *discordgo.Ready, channelID string, roles *[]discordgo.Role) {
	session.ChannelMessageSend(channelID, "おはよう世界")
	session.UpdateStatus(1, "Goのべんつよ、たのしいね")
	guilds := ready.Guilds

	if len(guilds) != 1 {
		panic("Can't boot 2 or more bot with same token")
	}

	guild := guilds[0]
	roleAddress := guild.Roles
	RoleListUpdate(roleAddress, roles)
}
