package service

import (
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/fxkt-tech/picsal/api/v1"
	"github.com/fxkt-tech/picsal/internal/biz"
)

type ImageService struct {
	v1.UnimplementedImageServiceServer

	img *biz.ImageUsecase
	log *log.Helper
}

func NewImageService(img *biz.ImageUsecase, logger log.Logger) *ImageService {
	return &ImageService{
		img: img,
		log: log.NewHelper(logger),
	}
}
