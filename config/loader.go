package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

// LoadTOML load toml file to obj
func LoadTOML(filename string, obj any) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(bs, obj)
	return err
}
