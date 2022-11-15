package v1

import (
	"github.com/TemurMannonov/blog/config"
	"github.com/TemurMannonov/blog/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}
