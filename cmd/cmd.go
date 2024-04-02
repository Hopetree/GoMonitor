package cmd

import (
	"fmt"
	"github.com/Hopetree/GoMonitor/cmd/commands"
	"github.com/urfave/cli/v2"
	"os"
)

func Execute(args []string) {
	newCmd().Execute(args)
}

type cmd struct {
	app *cli.App
}

func (c *cmd) Execute(args []string) {
	if err := c.app.Run(args); err != nil {
		fmt.Println("[Error]:", err)
		os.Exit(1)
	}
}

func newCmd() *cmd {
	version := "0.0.1"
	app := &cli.App{}
	app.Name = "GoMonitor"
	app.Version = version
	app.Usage = "一个简单的agent客户端，用于采集主机信息并上报到服务端"

	app.Commands = []*cli.Command{
		commands.Collect,
		commands.Decrypt,
		commands.Report,
	}
	return &cmd{app: app}
}
