package auth

import (
	"github.com/andorr/qs/utils"
	"fmt"
	"github.com/urfave/cli"
	"log"
)

func LogOutCommand() cli.Command {
	return cli.Command{
		Name: "logout",
		Usage: "qs logout",
		Action: HandleLogOut,
		Flags: []cli.Flag {
			cli.StringFlag{
				Name: "cookie",
			},
		},
	}
}

func HandleLogOut(c *cli.Context) error {
	config, err := utils.GetOrCreateConfig()
	if err != nil {
		log.Fatalf("Was not able to read config!\nError: %s\n", err.Error())
	}

	config.Cookie = ""
	err = utils.SaveConfig(config)
	if err != nil {
		log.Fatalf("Was not able to save config.\nError: %s\n", err.Error())
	}
	fmt.Println("Logged out!")

	return nil
}