package main

import (
	"github.com/joho/godotenv"
	"github.com/mitranjaykrishna/trade-recommendation/internal/config"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

}
