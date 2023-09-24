package handler

import (
	"main/config"
	"main/packages/logger"
	"main/storage"
)

type Handler struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
}

func NewHandler(strg storage.StorageI, conf config.Config, loger logger.LoggerI) *Handler {
	return &Handler{strg: strg, cfg: conf, log: loger}
}
