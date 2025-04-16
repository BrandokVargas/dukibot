package bot

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

type ResponseChannelCategory struct{}

func (r *ResponseChannelCategory) Name() string {
	return "ccategory"
}

func (r *ResponseChannelCategory) ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	categoryName := strings.Join(params, " ")

	_, err := s.GuildChannelCreate(m.GuildID, categoryName, discordgo.ChannelTypeGuildCategory)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Ups! Error al crear una categoria: "+err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Hola `"+m.Author.Username+"` haz creado la cateogia. `"+categoryName+"` ")
}

type ResponseChannelMultiplyCategory struct{}

func (r *ResponseChannelMultiplyCategory) Name() string {
	return "cmcategory"
}

func (r *ResponseChannelMultiplyCategory) ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {

	categories := strings.Split(strings.Join(params, " "), ",")
	var countCategory int = 0
	for _, categoryName := range categories {
		categoryName = strings.TrimSpace(categoryName)
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
