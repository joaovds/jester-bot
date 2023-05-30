package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Erro ao carregar .env")
  }
}

func main() {
  loadEnv()

  token := os.Getenv("DISCORD_BOT_TOKEN")
  if token == "" {
    log.Fatal("Token não especificado")
  }
  
  log.Println("Jester!")
}

