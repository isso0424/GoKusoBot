package handler

import "github.com/bwmarrin/discordgo"

func UpdateRole(session *discordgo.Session, event *discordgo.GuildRoleUpdate, roles []discordgo.Role) {
	newRole := event.GuildRole.Role
	if alreadyRoleExists(newRole.Name) {
		return
	}
	message := "新しいクソRole***†" + newRole.Name + "†***が追加されたよ"

	session.ChannelMessageSend("574884574778359844", message)
	//sendMessage(session, "690909527461199922", message)
}

func alreadyRoleExists(roleName string, roles []discordgo.Role) bool {
	for _, value := range roles {
		if value.Name == roleName {
			return true
		}
	}
	return false
}
