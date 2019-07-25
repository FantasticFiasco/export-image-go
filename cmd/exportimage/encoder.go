package main

import (
    "image/jpeg"
	"os"

    "github.com/pkg/errors"
)

// Encoder is capable of encoding a JPEG image.
type encoder struct {
	options jpeg.Options
}

// New allocates and returns a new Encoder.
func newEncoder(quality int) encoder {
    return encoder{
        options: jpeg.Options{
            Quality: quality,
        },
	}
}

// Encode an input file into an output JPEG file.
func (e encoder) encode(inputFile string, outputFile string) error {
	r, err := os.Open(inputFile)
	if err != nil {
	    return errors.Wrapf(err, "open image %s failed", inputFile)
    }
    defer r.Close()

	input, err := jpeg.Decode(r)
    if err != nil {
        return errors.Wrapf(err, "decode image %s failed", inputFile)
    }

	w, err := os.Create(outputFile)
	if err != nil {
	    return errors.Wrapf(err, "create file %s failed", outputFile)
    }
	defer w.Close()

	err = jpeg.Encode(w, input, &e.options)
    if err != nil {
        return errors.Wrapf(err, "encode image %s failed", outputFile)
    }

	return  nil
}
