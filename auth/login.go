package auth

import (
	"QS/utils"
	"fmt"
	"github.com/urfave/cli"
	"log"
)

func LogInCommand() cli.Command {
	return cli.Command{
		Name: "login",
		Action: HandleLogIn,
		Flags: []cli.Flag {
			cli.StringFlag{
				Name:        "cookie",
			},
		},
	}
}

func HandleLogIn(c *cli.Context) {
	config, err := utils.GetOrCreateConfig()
	if err != nil {
		log.Fatalf("Unable to read config!\nError: %s", err.Error())
	}

	// Check if cookie is provided
	if c.String("cookie") != "" {
		config.Cookie = c.String("cookie")
	} else {
		// Get username and password
		log.Fatalf("Not implemented yet! Please use --cookie instead")
	}

	err = utils.SaveConfig(config)
	if err != nil {
		log.Fatalf("Was not able to save config.\nError: %s", err.Error())
	}
	fmt.Println("\nSuccessfully logged in!")
}