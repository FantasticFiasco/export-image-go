package main

import (
	"image/jpeg"
	"os"
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
func (e encoder) encode(inputFile string, outputFile string) {
	r, err := os.Open(inputFile)
	check(err)
	defer r.Close()

	input, err := jpeg.Decode(r)
	check(err)

	w, err := os.Create(outputFile)
	check(err)
	defer w.Close()

	err = jpeg.Encode(w, input, &e.options)
	check(err)
}
