package commands

import (
	"log"

	"github.com/ravbaker/goharm/internal/config"
)

func Logout(cfg *config.Config) {
	cfg.General.Authentication = config.DefaultConfig.General.Authentication
	err := config.UpdateConfig(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
