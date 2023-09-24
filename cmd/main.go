package main

import (
	"context"
	"main/api"
	"main/api/handler"
	"main/config"
	"main/packages/logger"
	"main/storage/db"
)

func main() {

	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := db.NewStorage(context.Background(), *cfg)
	if err != nil {
		return
	}

	h := handler.NewHandler(strg, *cfg, log)

	r := api.NewServer(h)
	r.Run(":8010")

}
