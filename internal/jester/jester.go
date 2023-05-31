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

  jester.session.Close()

  log.Println("Jester parado!...")

  return nil
}

func (jester *Jester) handleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
  log.Printf("Mensagem recebida: %s", messageCreate.Content)
  if messageCreate.Content == "!jcls" {
    log.Println("Limpando...")
    log.Println("ServerID: ", messageCreate.GuildID, " CanalID: ", messageCreate.ChannelID);
  }
}

