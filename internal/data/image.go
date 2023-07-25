package data

import (
	"bytes"
	"context"
	"errors"
	stdimage "image"

	"github.com/fxkt-tech/picsal/internal/biz"
	"github.com/fxkt-tech/picsal/pkg/image"
	"github.com/fxkt-tech/picsal/pkg/image/filter"
	"github.com/fxkt-tech/picsal/pkg/image/io"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type imageRepo struct {
	data *Data
	log  *log.Helper
}

func NewImageRepo(data *Data, logger log.Logger) biz.ImageRepo {
	return &imageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *imageRepo) CreateJob(ctx context.Context, p *biz.CreateJobParams) (*biz.CreateJobResult, error) {
	jobid := uuid.NewString()

	img, err := io.ReadFile(p.FilePath)
	if err != nil {
		return nil, err
	}

	r.data.imgStore.Put(jobid, img)

	return &biz.CreateJobResult{Jobid: jobid}, nil
}

func (r *imageRepo) Scale(ctx context.Context, p *biz.ScaleParams) (*biz.ImageResult, error) {
	img, ok := r.data.imgStore.Get(p.Jobid)
	if !ok {
		return nil, errors.New("任务不存在")
	}

	nimg, err := filter.Scale(img, int(p.Width), int(p.Height))
	if err != nil {
		return nil, err
	}

	r.data.imgStore.Put(p.Jobid, nimg)

	buf := &bytes.Buffer{}
	err = image.Write(buf, nimg)
	if err != nil {
		return nil, err
	}

	return &biz.ImageResult{ImageBytes: buf.Bytes()}, nil
}

func (r *imageRepo) Clip(ctx context.Context, p *biz.ClipParams) (*biz.ImageResult, error) {
	img, ok := r.data.imgStore.Get(p.Jobid)
	if !ok {
		return nil, errors.New("任务不存在")
	}

	nimg, err := filter.Clip(img, stdimage.Rectangle{
		Min: stdimage.Pt(int(p.XOffset), int(p.YOffset)),
		Max: stdimage.Point{int(p.XOffset + p.Width), int(p.YOffset + p.Height)},
	})
	if err != nil {
		return nil, err
	}

	r.data.imgStore.Put(p.Jobid, nimg)

	buf := &bytes.Buffer{}
	err = image.Write(buf, nimg)
	if err != nil {
		return nil, err
	}

	return &biz.ImageResult{ImageBytes: buf.Bytes()}, nil
}

// func (r *imageRepo) dealHandle(ctx context.Context, )
