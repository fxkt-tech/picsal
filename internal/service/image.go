package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/fxkt-tech/picsal/api/image/v1"
	"github.com/fxkt-tech/picsal/internal/biz"
)

type ImageService struct {
	v1.UnimplementedProcessingServiceServer

	img *biz.ImageUsecase
	log *log.Helper
}

func NewImageService(img *biz.ImageUsecase, logger log.Logger) *ImageService {
	return &ImageService{
		img: img,
		log: log.NewHelper(logger),
	}
}

func (s *ImageService) CreateJob(ctx context.Context, in *v1.CreateJobRequest) (*v1.CreateJobResponse, error) {
	s.log.Infof("CreateJob.req: %s", in)

	p := &biz.CreateJobParams{
		FilePath: in.F,
	}
	result, err := s.img.CreateJob(ctx, p)
	if err != nil {
		return nil, err
	}

	return &v1.CreateJobResponse{
		Id: result.Jobid,
	}, nil
}

func (s *ImageService) Scale(ctx context.Context, in *v1.ScaleRequest) (*v1.ImageStreamResponse, error) {
	s.log.Infow("Scale.req: %s", in)

	result, err := s.img.Scale(ctx, &biz.ScaleParams{
		Jobid:  in.Id,
		Width:  in.W,
		Height: in.H,
	})
	if err != nil {
		return nil, err
	}

	return &v1.ImageStreamResponse{
		ImageStream: result.ImageBytes,
		ContentType: "image/jpeg",
	}, nil
}
