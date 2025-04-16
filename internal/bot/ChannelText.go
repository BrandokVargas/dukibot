package bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type ResponseChannel struct{}

func (r *ResponseChannel) Name() string {
	return "ctext"
}

func (r *ResponseChannel) ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	channelName := strings.Join(args, "-")
	_, err := s.GuildChannelCreate(m.GuildID, channelName, discordgo.ChannelTypeGuildText)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Ups! Error al crear el canal: "+err.Error())
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Hola `"+m.Author.Username+"` haz creado el canal. `"+channelName+"` ")
}
