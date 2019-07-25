package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: When should we use structs or interfaces in packages?
// TODO: When should we use pointers?
// TODO: How do we implement error handling?

func main() {
	files := os.Args[1:]

	fmt.Printf("Exporting %d images...\n", len(files))

	settings, err := newSettings()
	check(err)

	encoder := newEncoder(settings.Quality)

	for i, f := range files {
		outDir, outFile := getOutPaths(f, settings.SubDirectory)

		fmt.Printf("%d/%d\t-> %s\n", i+1, len(files), outFile)

		err := os.MkdirAll(outDir, os.ModePerm)
		check(err)

		err = encoder.encode(f, outFile)
		check(err)
	}

	fmt.Println("Done!")

	if settings.PreventTermination {
		preventTermination()
	}
}

func check(err error) {
    if err != nil {
        panic(err)
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
	check(err)
}
