package filter

import (
	"testing"

	"github.com/fxkt-tech/picsal/pkg/image/io"
)

func TestGrey(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := io.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = Grey(canvas, GreyClassic)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = io.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
