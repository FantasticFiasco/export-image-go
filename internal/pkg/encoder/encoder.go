package encoder

import (
	"image/jpeg"
	"os"

	"github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
)

// Encoder is capable of encoding a JPEG image.
type Encoder struct {
	Quality int
}

// New allocates and returns a new Encoder.
func New(quality int) Encoder {
	return Encoder{
		Quality: quality,
	}
}

// Encode an input file into an output JPEG file.
func (e Encoder) Encode(inputFile string, outputFile string) {
	r, err := os.Open(inputFile)
	errors.Check(err)
	defer r.Close()

	input, err := jpeg.Decode(r)
	errors.Check(err)

	w, err := os.Create(outputFile)
	errors.Check(err)
	defer w.Close()

	options := jpeg.Options{
		Quality: e.Quality,
	}

	err = jpeg.Encode(w, input, &options)
	errors.Check(err)
}
