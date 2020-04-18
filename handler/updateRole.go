package handler

import "github.com/bwmarrin/discordgo"

func UpdateRole(session *discordgo.Session, event *discordgo.GuildRoleUpdate, roles []discordgo.Role, rootChannelID string) {
	newRole := event.GuildRole.Role
	if alreadyRoleExists(newRole.Name, roles) {
		return
	}
	message := "新しいクソRole***†" + newRole.Name + "†***が追加されたよ"

	session.ChannelMessageSend(rootChannelID, message)
}

func alreadyRoleExists(roleName string, roles []discordgo.Role) bool {
	for _, value := range roles {
		if value.Name == roleName {
			return true
		}
	}
	return false
}
