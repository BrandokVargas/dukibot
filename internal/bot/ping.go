package bot

import (
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct{}

func (r *PingCommand) Name() string {
	return "ping"
}

func (r *PingCommand) ExecuteChannel(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {

	s.ChannelMessageSend(m.ChannelID, "DukiPing")

}
