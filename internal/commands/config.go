package commands

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/ravbaker/goharm/internal/config"
)

func Config(cfg *config.Config) {
	toml.NewEncoder(os.Stdout).Encode(cfg)
}
