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
func (e encoder) encode(inputFile string, outputFile string) (err error) {
	r, err := os.Open(inputFile)
	if err != nil {
	    return err
    }
    defer r.Close()

	input, err := jpeg.Decode(r)
    if err != nil {
        return err
    }

	w, err := os.Create(outputFile)
	if err != nil {
	    return err
    }
	defer w.Close()

	err = jpeg.Encode(w, input, &e.options)
    if err != nil {
        return err
    }

	return nil
}
