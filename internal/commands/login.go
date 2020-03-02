package commands

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/ravbaker/goharm/internal/config"
	"github.com/ravbaker/goharm/internal/jsonapi/client"
	"github.com/ravbaker/goharm/internal/jsonapi/resources"
)

func Login(cfg *config.Config, email, password string) {
	if len(password) == 0 {
		fmt.Printf("Now, please type in the password (mandatory): ")
		passwordBytes, _ := terminal.ReadPassword(0)
		password = string(passwordBytes)
	}
	authentication := getAuthentication(email, password)
	cfg.General.Authentication = config.Authentication{
		AccessToken: authentication.AccessToken,
		UserRole:    authentication.UserRole,
		UserId:      authentication.UserId,
	}
	err := config.UpdateConfig(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getAuthentication(email, password string) *resources.Authentication {
	accessTokenRequest := resources.AccessToken{Email: email, Password: password}
	authentication := new(resources.Authentication)
	err := client.Request("/api/v1/access-tokens", http.MethodPost, &accessTokenRequest, authentication)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return authentication
}
