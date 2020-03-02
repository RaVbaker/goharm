package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ravbaker/goharm/internal/commands"
	"github.com/ravbaker/goharm/internal/config"
	"github.com/ravbaker/goharm/internal/jsonapi/client"
)

var cfg *config.Config
var Version = ""

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	client.Host = cfg.General.ApiHost
	client.AccessToken = cfg.General.Authentication.AccessToken
}

func main() {
	showVersion := flag.Bool("v", false, "Show version number")
	flag.Parse()
	
	if *showVersion {
		println(Version)
		return
	}
	
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		loginCmd := flag.NewFlagSet("login", flag.PanicOnError)
		email := loginCmd.String("email", "", "Provide your account email")
		password := loginCmd.String("password", "", "Provide your account password")
		loginCmd.Parse(os.Args[2:])
		if len(*email) == 0 {
			loginCmd.Usage()
			os.Exit(1)
		}

		commands.Login(cfg, *email, *password)
	case "logout":
		commands.Logout(cfg)
	case "config":
		commands.Config(cfg)
	case "time-logs":
		timeLogsCmd := flag.NewFlagSet("time-logs", flag.PanicOnError)
		rangeFilter := timeLogsCmd.String("range", "m", "Desired time range (either 'month' or 'week'), default is 'week'")
		timeLogsCmd.Parse(os.Args[2:])
		commands.TimeLogs(cfg, *rangeFilter)
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`Rebased Harmonogram CLI`)
	flag.Usage()
	fmt.Println("\n\tgoharm [SUBCOMMAND]")
	fmt.Println("\nexpected 'login', 'logout', 'config' or 'time-logs' subcommands")
}
