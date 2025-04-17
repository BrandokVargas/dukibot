package bot

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
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

type ResponseChannelTextMultiply struct{}

func (r *ResponseChannelTextMultiply) Name() string {
	return "cmtext"
}

func (r *ResponseChannelTextMultiply) ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	texts := strings.Split(strings.Join(params, " "), ",")
	var countText int = 0
	for _, categoryName := range texts {
		categoryName = strings.TrimSpace(categoryName)
		if categoryName == "" {
			continue
		}

		_, err := s.GuildChannelCreate(m.GuildID, categoryName, discordgo.ChannelTypeGuildText)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error al crear los canales textos '"+categoryName+"': "+err.Error())
			continue
		}
		countText++
	}

	s.ChannelMessageSend(m.ChannelID, " ["+strconv.Itoa(countText)+"] canales de texto han sido creados")

}
