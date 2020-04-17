package handler

import "github.com/bwmarrin/discordgo"

func RoleListUpdate(roleList []*discordgo.Role, roles *[]discordgo.Role) {
	*roles = []discordgo.Role{}
	for _, role := range roleList {
		*roles = append(*roles, *role)
	}
}
