package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/FantasticFiasco/export-image-go/internal/pkg/encoder"
	"github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
	"github.com/FantasticFiasco/export-image-go/internal/pkg/settings"
)

// TODO: When should we use structs or interfaces in packages?
// TODO: How do we implement error handling?

func main() {
	files := os.Args[1:]

	fmt.Printf("Exporting %d images...\n", len(files))

	settings := settings.New()
	encoder := encoder.New(settings.Quality)

	for i, f := range files {
		outDir, outFile := getOutPaths(f, settings.SubDirectory)

		fmt.Printf("%d/%d\t-> %s\n", i+1, len(files), outFile)

		err := os.MkdirAll(outDir, os.ModePerm)
		errors.Check(err)

		encoder.Encode(f, outFile)
	}

	fmt.Println("Done!")

	if settings.PreventTermination {
		preventTermination()
	}
}

func getOutPaths(file string, subDirectory string) (outDir string, outFile string) {
	dir, fileName := filepath.Split(file)
	outDir = filepath.Join(dir, subDirectory)
	outFile = filepath.Join(outDir, fileName)
	return
}

func preventTermination() {
	fmt.Println("\nPress any key to continue...")
	r := bufio.NewReader(os.Stdin)
	_, _, err := r.ReadRune()
	errors.Check(err)
}
