package commands

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CreatedChannelCategory(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	categoryName := strings.Join(params, " ")

	_, err := s.GuildChannelCreate(m.GuildID, categoryName, discordgo.ChannelTypeGuildCategory)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Ups! Error al crear una categoria: "+err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Hola `"+m.Author.Username+"` haz creado la cateogia. `"+categoryName+"` ")
}

func CreatedChannelMultiplyCategory(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	// Unir todo en un solo string y dividir por "-"
	categories := strings.Split(strings.Join(params, " "), ",")
	var countCategory int = 0
	for _, categoryName := range categories {
		categoryName = strings.TrimSpace(categoryName) // Eliminar espacios extras
		if categoryName == "" {
			continue
		}

		_, err := s.GuildChannelCreate(m.GuildID, categoryName, discordgo.ChannelTypeGuildCategory)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error al crear la categoría '"+categoryName+"': "+err.Error())
			continue
		}
		countCategory++
	}

	s.ChannelMessageSend(m.ChannelID, " ["+strconv.Itoa(countCategory)+"] categorias han sido creadas")
}
