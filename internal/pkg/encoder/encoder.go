package encoder

import (
    "github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
    "image/jpeg"
    "os"
)

var options = jpeg.Options{
    Quality: 95,
}

type encoder struct {
}

func New() encoder {
    e := encoder{}
    return e
}

func (e encoder) Encode(inputFile string, outputFile string) {
    r, err := os.Open(inputFile)
    errors.Check(err)

    input, err := jpeg.Decode(r)
    errors.Check(err)

    w, err := os.Create(outputFile)
    errors.Check(err)

    err = jpeg.Encode(w, input, &options)
    errors.Check(err)

    // TODO: Can we dispose these any other way?
    r.Close()
    w.Close()
}
