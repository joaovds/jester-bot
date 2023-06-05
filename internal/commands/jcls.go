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

if newChannel.Name == "geral" {
  // Define a posição do novo canal para 0, para fixá-lo no topo da lista
  _, err = session.ChannelEditComplex(newChannel.ID, &discordgo.ChannelEdit{
      Position: 0,
  })
  if err != nil {
      log.Println("Erro ao fixar canal 'geral':", err)
      return err
  }

  // Redireciona os usuários para o novo canal
  _, err = session.ChannelMessageSend(messageCreate.ChannelID, fmt.Sprintf("O canal 'geral' foi recriado. Clique aqui para acessar: <#%s>", newChannel.ID))
  if err != nil {
      log.Println("Erro ao enviar mensagem de redirecionamento:", err)
      return err
  }
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
  return "!cls"
}

func (command *JclsCommand) Description() string {
  return "Comando para deletar o histórico de mensagens do canal"
}

