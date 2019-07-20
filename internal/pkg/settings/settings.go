package settings

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/FantasticFiasco/export-image-go/internal/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Settings is wrapping the YAML configuration file.
type Settings struct {
	Quality            int    `yaml:"quality"`
	SubDirectory       string `yaml:"subDirectory"`
	PreventTermination bool   `yaml:"preventTermination"`
}

// New allocates and returns a new Encoder.
func New() Settings {
	exec, err := os.Executable()
	errors.Check(err)

	f := filepath.Join(filepath.Dir(exec), "exportimage.yml")

	data, err := ioutil.ReadFile(f)
	errors.Check(err)

	var s Settings
	err = yaml.Unmarshal(data, &s)
	errors.Check(err)

	return s
}
