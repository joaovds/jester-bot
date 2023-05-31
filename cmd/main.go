package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

  "jester/internal/jester"
)

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Erro ao carregar .env : ", err)
  }
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

  err = jester.JesterRun()
  if err != nil {
    log.Fatal("Erro ao iniciar Jester...:", err)
  }
}

