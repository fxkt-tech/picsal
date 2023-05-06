package image

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"os"
)

func ReadFile(path string) (draw.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	canvas := image.NewNRGBA(img.Bounds())

	rect := img.Bounds()
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			p := img.At(x, y)
			r, g, b, a := p.RGBA()
			canvas.Set(x, y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	return canvas, nil
}

func Write(rw io.ReadWriter, cvs draw.Image) error {
	// buf := new(bytes.Buffer)
	return jpeg.Encode(rw, cvs, nil)
}

// func WriteFile(cvs draw.Image, path string) (err error) {
// 	file, err := os.Create(path)
// 	if err != nil {
// 		return
// 	}

// 	suffix := filepath.Ext(path)
// 	switch suffix {
// 	case ".jpeg", ".jpg":
// 		return jpeg.Encode(file, cvs, nil)
// 	case ".png":
// 		return png.Encode(file, cvs)
// 	}
// 	return errors.ExtNotSupported
// }

func Scale(cvs draw.Image, ow, oh int) (draw.Image, error) {
	r := cvs.Bounds()
	iw := r.Max.X - r.Min.X
	ih := r.Max.Y - r.Min.Y
	newcvs := image.NewNRGBA(image.Rect(0, 0, ow, oh))
	for oy := 0; oy <= oh; oy++ {
		for ox := 0; ox <= ow; ox++ {
			ix := ox * iw / ow
			iy := oy * ih / oh
			c := cvs.At(ix, iy)
			newcvs.Set(ox, oy, c)
		}
	}

	return newcvs, nil
}
