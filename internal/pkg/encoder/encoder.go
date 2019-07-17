package encoder

import (
    "github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
    "image/jpeg"
    "os"
)

// Encoder is capable of encoding a JPEG image.
type Encoder struct {
    Quality int
}

// New allocates and returns a new Encoder.
func New(quality int) Encoder {
    e := Encoder{
        Quality: quality,
    }
    return e
}

// Encode an input file into an output JPEG file.
func (e Encoder) Encode(inputFile string, outputFile string) {
    r, err := os.Open(inputFile)
    errors.Check(err)

    input, err := jpeg.Decode(r)
    errors.Check(err)

    w, err := os.Create(outputFile)
    errors.Check(err)

    options := jpeg.Options{
        Quality: e.Quality,
    }

    err = jpeg.Encode(w, input, &options)
    errors.Check(err)

    // TODO: Can we dispose these any other way?
    r.Close()
    w.Close()
}
