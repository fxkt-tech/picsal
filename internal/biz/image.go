package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ImageRepo interface {
	CreateJob(ctx context.Context, p *CreateJobParams) (*CreateJobResult, error)
	Scale(ctx context.Context, p *ScaleParams) (*ImageResult, error)
	Clip(ctx context.Context, p *ClipParams) (*ImageResult, error)
}

type ImageUsecase struct {
	repo ImageRepo
	log  *log.Helper
}

func NewImageUsecase(repo ImageRepo, logger log.Logger) *ImageUsecase {
	return &ImageUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ImageUsecase) CreateJob(ctx context.Context, p *CreateJobParams) (*CreateJobResult, error) {
	return uc.repo.CreateJob(ctx, p)
}

func (uc *ImageUsecase) Scale(ctx context.Context, p *ScaleParams) (*ImageResult, error) {
	return uc.repo.Scale(ctx, p)
}

func (uc *ImageUsecase) Clip(ctx context.Context, p *ClipParams) (*ImageResult, error) {
	return uc.repo.Clip(ctx, p)
}
