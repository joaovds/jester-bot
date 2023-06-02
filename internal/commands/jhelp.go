package commands

import (
	"fmt"
	"log"
	"strings"

	"jester/internal/jester"

	"github.com/bwmarrin/discordgo"
)

type JHelpCommand struct {
  commands map[string]jester.Command
}

func NewHelpCommand(commands map[string]jester.Command) *JHelpCommand {
  return &JHelpCommand{
    commands: commands,
  }
}

func (command *JHelpCommand) HandleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) error {
  log.Print("Help...")

  var helpText strings.Builder

  helpText.WriteString("## Lista de comandos do *Jester*:\n\n")

  for commandName, currentCommand := range command.commands {
    helpText.WriteString(fmt.Sprintf("**%s:** *%s*\n", commandName, currentCommand.Description()))
  }

  _, err := session.ChannelMessageSend(messageCreate.ChannelID, helpText.String())
	if err != nil {
		log.Println("Erro ao enviar mensagem de ajuda:", err)
	}

  log.Println("Help feito com sucesso!")
  return nil
}

func (command *JHelpCommand) Name() string {
  return "!jhelp"
}

func (command *JHelpCommand) Description() string {
  return "Comando para mostrar as opções do jester"
}

