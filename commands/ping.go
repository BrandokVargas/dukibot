package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "Davicho cabro 😎")
	if err != nil {
		fmt.Println("Error enviando mensaje")
	}

}
