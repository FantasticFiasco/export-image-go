package main

import (
    "fmt"
    "github.com/FantasticFiasco/export-image-go/internal/pkg/encoder"
    "github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
    "github.com/FantasticFiasco/export-image-go/internal/pkg/settings"
    "os"
    "path/filepath"
)

func main() {
    args := os.Args[1:]

    fmt.Printf("Exporting %d images...\n", len(args))

    settings := settings.New()
    encoder := encoder.New(settings.Quality)

    for index, filePath := range args  {
        dir, file := filepath.Split(filePath)
        exportDir := filepath.Join(dir, "Exported")
        exportFilePath := filepath.Join(exportDir, file)

        fmt.Printf("%d/%d\t-> %s\n", index + 1, len(args), exportFilePath)

        err := os.MkdirAll(exportDir, os.ModePerm)
        errors.Check(err)

        encoder.Encode(filePath, exportFilePath)
    }

    fmt.Println("Done!")
}
