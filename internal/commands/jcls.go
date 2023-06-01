package commands

import (
  "log"

  "github.com/bwmarrin/discordgo"
)

type JclsCommand struct {}

func (command *JclsCommand) HandleCommand(session *discordgo.Session, messageCreate *discordgo.MessageCreate) error {
  log.Print("Limpando...")

  channel, err := session.Channel(messageCreate.ChannelID)
  if err != nil {
    log.Println("Erro ao obter configurações do canal:", err)
  }

  _, err = session.ChannelDelete(messageCreate.ChannelID)
  if err != nil {
    log.Println("Erro ao excluir canal:", err)
    return err
  }

  newChannel, err := session.GuildChannelCreate(messageCreate.GuildID, channel.Name, discordgo.ChannelTypeGuildText)
  if err != nil {
    log.Println("Erro ao criar novo canal:", err)
    return err
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
    return err
  }

  log.Println("Canal excluído e recriado com sucesso!")
  return nil
}

func (command *JclsCommand) Name() string {
  return "!jcls"
}

func (command *JclsCommand) Description() string {
  return "Comando para deletar o histórico de mensagens do canal"
}

