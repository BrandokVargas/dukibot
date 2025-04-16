package bot

import (
	"github.com/BrandokVargas/dukibot/internal/constants"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	session *discordgo.Session
}

func New(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Bot{session: dg}, nil
}

func (b *Bot) Start(router *Router) error {

	b.session.AddHandler(router.HandleMessageCreate)

	b.session.Identify.Intents = discordgo.IntentsGuildMessages
	err := b.session.Open()

	if err != nil {
		return err
	}
	log.Println("Started", constants.NameBot)
	return nil
}

func (b *Bot) Stop() {
	b.session.Close()
}
