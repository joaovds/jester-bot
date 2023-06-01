package jester

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Jester struct {
  session *discordgo.Session
}

func NewJester(token string) (*Jester, error) {
  session, err := discordgo.New("Bot " + token)
  if err != nil {
    return nil, err
  }

  return &Jester{
    session,
  }, nil
}

func (jester *Jester) JesterRun() error {
  jester.session.AddHandler(jester.handleCommand)

  err := jester.session.Open()
  if err != nil {
    return err
  }

  log.Println("Jester em execução!... (Ctrl+C para parar)")

  closeApp := make(chan os.Signal, 1)

  signal.Notify(closeApp, os.Interrupt, syscall.SIGTERM)
  <-closeApp

  // jester.session.Close()

  log.Println("Jester parado!...")

  return nil
}

func (jester *Jester) handleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
  if messageCreate.Content == "!jcls" {
    log.Print("Limpando...")

    channel, err := session.Channel(messageCreate.ChannelID)
    if err != nil {
      log.Println("Erro ao obter configurações do canal:", err)
      return
    }

    _, err = session.ChannelDelete(messageCreate.ChannelID)
    if err != nil {
      log.Println("Erro ao excluir canal:", err)
      return
    }

    newChannel, err := session.GuildChannelCreate(messageCreate.GuildID, channel.Name, discordgo.ChannelTypeGuildText)
    if err != nil {
      log.Println("Erro ao criar novo canal:", err)
      return
    }

    newChannelConfig := discordgo.ChannelEdit{
      Name: channel.Name,
      Topic: channel.Topic,
      PermissionOverwrites: channel.PermissionOverwrites,
      ParentID: channel.ParentID,
    }

    _, err = session.ChannelEdit(newChannel.ID, &newChannelConfig)
    if err != nil {
			log.Println("Erro ao atualizar canal:", err)
			return
		}

		log.Println("Canal excluído e recriado com sucesso!")
  }
}

