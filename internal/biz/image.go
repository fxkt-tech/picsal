package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type ImageRepo interface{}

type ImageUsecase struct {
	repo ImageRepo
	log  *log.Helper
}

func NewImageUsecase(repo ImageRepo, logger log.Logger) *ImageUsecase {
	return &ImageUsecase{repo: repo, log: log.NewHelper(logger)}
}
