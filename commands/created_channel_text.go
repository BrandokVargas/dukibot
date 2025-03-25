package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CreateTextChannel(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	// commandserrors.CommandError(s, m, args, "Ups! Debes proporcionar un nombre para el canal. Uso: `!create text <nombre-canal>`")

	// Unir los argumentos para formar el nombre del canal
	channelName := strings.Join(args, "-") // Discord usa guiones en los nombres de canales

	// Crear el canal en el mismo servidor
	_, err := s.GuildChannelCreate(m.GuildID, channelName, discordgo.ChannelTypeGuildText)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Ups! Error al crear el canal: "+err.Error())
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Hola `"+m.Author.Username+"` haz creado el canal. `"+channelName+"` ")
}
