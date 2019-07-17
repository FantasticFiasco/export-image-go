package settings

import (
    "github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "path/filepath"
)

type settings struct {
	Quality int `yaml:"quality"`
}

func New() settings {
    exec, err := os.Executable()
    errors.Check(err)

    file := filepath.Join(filepath.Dir(exec), "exportimage.yml")

    data, err := ioutil.ReadFile(file)
	errors.Check(err)

    s := settings{}
	err = yaml.Unmarshal(data, &s)
	errors.Check(err)

	return s
}
