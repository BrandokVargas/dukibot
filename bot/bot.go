package bot

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/BrandokVargas/dukibot/commands"
	"github.com/bwmarrin/discordgo"
)

// Configuración del bot
type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

// Cargar configuración desde config.json
func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	fmt.Println("prefix: ", config.Prefix)

	return config, nil
}

// Inicializar el bot
func Start() (*discordgo.Session, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error cargando configuración: %v", err)
	}

	// Crear una nueva sesión de Discord
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, fmt.Errorf("error creando sesión de Discord: %v", err)
	}

	// Handler para cuando el bot esté listo
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot conectado como", s.State.User.Username)
	})

	dg.AddHandler(executeCommands)

	// Abrir conexión con Discord
	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión con Discord: %v", err)
	}

	return dg, nil
}

func executeCommands(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error cargando config", err)
		return
	}

	if !strings.HasPrefix(m.Content, config.Prefix) {
		return
	}
	// Quitar el prefijo y dividir en partes
	content := strings.TrimPrefix(m.Content, config.Prefix)
	args := strings.Fields(content) // Divide en palabras

	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0]) // Primer palabra = comando
	args = args[1:]                     // El resto son argumentos

	fmt.Println("Command:", command)
	fmt.Println("Args:", args)

	switch command {
	case "test":
		commands.Ping(s, m)
	case "ctext":
		if len(args) > 0 {
			commands.CreateTextChannel(s, m, args[0:])
		} else {
			s.ChannelMessageSend(m.ChannelID, "Comando desconocido. Usa: `!ctext <nombre-text>`")
		}

	case "ccategory":
		if len(args) > 0 {
			commands.CreatedChannelCategory(s, m, args[0:])
		} else {
			s.ChannelMessageSend(m.ChannelID, "Comando desconocido. Usa: `!ccategory <nombre-category>`")
		}

	case "cmcategory":
		if len(args) > 0 {
			commands.CreatedChannelMultiplyCategory(s, m, args[0:])
		} else {
			s.ChannelMessageSend(m.ChannelID, "Comando desconocido. Usa: `!cmcategory <category1,category2 >`")
		}
	default:
		_, _ = s.ChannelMessageSend(m.ChannelID, "Comando no reconocido.")
	}

}
