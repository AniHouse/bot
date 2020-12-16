package util

import (
	"github.com/bwmarrin/discordgo"
)

func HasRole(member *discordgo.Member, id string) bool {
	for _, role := range member.Roles {
		if role == id {
			return true
		}
	}
	return false
}
