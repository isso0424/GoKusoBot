package handler

import "github.com/bwmarrin/discordgo"

func CreateNewRole(session *discordgo.Session, event *discordgo.GuildRoleCreate, roles *[]discordgo.Role) {
	guild, err := session.State.Guild(event.GuildRole.GuildID)
	if err != nil {
		return
	}
	roleAddress := guild.Roles
	RoleListUpdate(roleAddress, roles)
}
