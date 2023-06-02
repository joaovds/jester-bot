package jester

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
  HandleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) error
  Name() string
  Description() string
}

type Jester struct {
  session *discordgo.Session
  prefix  string
  commands map[string]Command
}

func NewJester(token string) (*Jester, error) {
  session, err := discordgo.New("Bot " + token)
  if err != nil {
    return nil, err
  }

  prefix := "!"

  return &Jester{
    session: session,
    prefix: prefix,
    commands: make(map[string]Command),
  }, nil
}

func (jester *Jester) RegisterCommand(name string, command Command) {
  jester.commands[name] = command
}

func (jester *Jester) Commands() map[string]Command {
  return jester.commands
}

func (jester *Jester) JesterRun() error {
  jester.session.AddHandler(jester.handleRoutine)
  err := jester.session.Open()
  if err != nil {
    return err
  }

  log.Println("Jester em execução!... (Ctrl+C para parar)")

  jester.session.AddHandler(jester.handleCommand)

  closeApp := make(chan os.Signal, 1)

  signal.Notify(closeApp, os.Interrupt, syscall.SIGTERM)
  <-closeApp

  jester.session.Close()

  log.Println("Jester parado!...")

  return nil
}

func (jester *Jester) handleRoutine(session *discordgo.Session, _ *discordgo.Connect) {
  HandleCoffee(session)
}

func (jester *Jester) handleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
  if !strings.HasPrefix(messageCreate.Content, jester.prefix) {
    return
  }

  args := strings.Split(messageCreate.Content, " ")
  command := strings.TrimPrefix(args[0], jester.prefix)

  commandInstance, ok := jester.commands[command]
  if !ok {
    return
  }

  err := commandInstance.HandleCommand(session, messageCreate)
  if err != nil {
    log.Println("Erro ao executar comando:", err)
  }
}

func HandleCoffee(session *discordgo.Session) error {
  log.Print("Hora do café...")

  for _, server := range session.State.Guilds {
		systemChannel, err := session.State.Channel(server.SystemChannelID)
		if err != nil {
			log.Println("Erro ao obter informações do canal principal:", err)
			continue
		}

		permissions, err := session.State.UserChannelPermissions(session.State.User.ID, systemChannel.ID)
		if err != nil {
			log.Println("Erro ao obter permissões do canal principal:", err)
			continue
		}

		if permissions&discordgo.PermissionViewChannel != 0 && permissions&discordgo.PermissionSendMessages != 0 {
			_, err := session.ChannelMessageSend(systemChannel.ID, "A hora do café!")
      _, err = session.ChannelMessageSend(server.ID, "## Opaaaa! Hora do cafezinho!!!")
			if err != nil {
				log.Println("Erro ao enviar mensagem:", err)
			}
		}
	}

  return nil
}

