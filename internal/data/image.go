package data

import (
	"github.com/fxkt-tech/picsal/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type imageRepo struct {
	data *Data
	log  *log.Helper
}

func NewAIRepo(data *Data, logger log.Logger) biz.ImageRepo {
	return &imageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
