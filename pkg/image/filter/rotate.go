package filter

import (
	"image/draw"

	"github.com/fxkt-tech/picsal/pkg/image/errors"
)

type RotateType uint8

const (
	Rotate180 RotateType = iota
)

func Rotate(cvs draw.Image, rt RotateType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	// TODO: ...

	return nil
}
