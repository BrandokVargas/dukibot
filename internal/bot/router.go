package bot

import (
	"github.com/BrandokVargas/dukibot/internal/domain"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Router struct {
	prefix          string
	commandsExecute map[string]domain.Command
}

func NewRouter(prefix string) *Router {
	return &Router{
		prefix:          prefix,
		commandsExecute: make(map[string]domain.Command),
	}
}

func (r *Router) Register(cmd domain.Command) {
	r.commandsExecute[cmd.Name()] = cmd
}

func (r *Router) HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, r.prefix) {
		return
	}

	content := strings.TrimPrefix(m.Content, r.prefix)
	args := strings.Fields(content)
	if len(args) == 0 {
		return
	}

	cmdName := args[0]
	args = args[1:]

	if cmd, exists := r.commandsExecute[cmdName]; exists {
		cmd.ExecuteChannel(s, m, args)
	} else {
		similar, score := getMostSimilarCommand(cmdName, r.commandsExecute)
		if score > 0.5 {
			s.ChannelMessageSend(m.ChannelID, "Comando desconocido. ¿Quisiste decir `!"+similar+"`?")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Comando desconocido")
		}
	}
}
