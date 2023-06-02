package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"jester/internal/commands"
	"jester/internal/jester"
)

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Erro ao carregar .env : ", err)
  }
}

func RegisterCommands(jester *jester.Jester) {
  jester.RegisterCommand("cls", &commands.JclsCommand{})

  jester.RegisterCommand("lunch", &commands.JLunchCommand{})

  helpCommands := commands.NewHelpCommand(jester.Commands())
  jester.RegisterCommand("help", helpCommands)
}

func main() {
  loadEnv()

  token := os.Getenv("DISCORD_BOT_TOKEN")
  if token == "" {
    log.Fatal("Token não especificado")
  }
  
  jester, err := jester.NewJester(token)
  if err != nil {
    log.Fatal("Erro ao criar Jester...:", err)
  }

  RegisterCommands(jester)

  err = jester.JesterRun()
  if err != nil {
    log.Fatal("Erro ao iniciar Jester...:", err)
  }
}

