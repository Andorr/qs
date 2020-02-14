package cmd

import (
	"QS/auth"
	"QS/config"
	"QS/queue"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "QS CLI"
	app.Description = "A QS command line tool for adding yourself in queue"
	app.Version = "1.0.0"
	app.Usage = "qs <COMMAND> [arguments...]"

	app.Commands = []cli.Command{
		queue.Command(),
		auth.LogInCommand(),
		auth.LogOutCommand(),
		config.Command(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}
}