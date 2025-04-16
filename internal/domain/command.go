package domain

import "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}
