package commands

import (
	"strings"
	"math/rand"
	"time"
	"log"

	"jester/internal/jester"

	"github.com/bwmarrin/discordgo"
)

type JLunchCommand struct {
	commands map[string]jester.Command
}

func NewLunchCommand(commands map[string]jester.Command) *JLunchCommand {
  return &JLunchCommand{
    commands: commands,
  }
}

func (command *JLunchCommand) HandleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) error {

	var helpText strings.Builder

	restaurants := []string{
		"Vila 🏘️",
		"Brunão 👨",
		"Abelhinha 🐝",
		"Cantinho 🍽️",
		"Caiçara 🌊",
		"La Bombonera 🍫",
		"New Era 🆕",
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(restaurants))

	restaurant := restaurants[randomIndex]

	log.Println("Restaurante escolhido: ", restaurant)
	helpText.WriteString("O restaurante de hoje é:\n")
	helpText.WriteString(restaurant)

	_, err := session.ChannelMessageSend(messageCreate.ChannelID, helpText.String())
	if err != nil {
		log.Println("Erro ao enviar mensagem de ajuda:", err)
	}

  return nil
}

func (command *JLunchCommand) Name() string {
  return "!lunch"
}

func (command *JLunchCommand) Description() string {
  return "Comando para mostrar as opções de almoço"
}

