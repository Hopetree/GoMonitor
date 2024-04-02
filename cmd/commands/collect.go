package commands

import (
	"fmt"
	"github.com/Hopetree/GoMonitor/host"
	"github.com/urfave/cli/v2"
)

var Collect = &cli.Command{
	Name:  "collect",
	Usage: "采集主机信息",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Value:   6,
			Usage:   "采集间隔时间（默认值：6）",
		},
		&cli.StringFlag{
			Name:  "cmd",
			Value: "echo ''",
			Usage: "用来采集服务版本信息的命令，可以使用带管道符的shell命令",
		},
	},
	Action: collectCmd,
}

func collectCmd(ctx *cli.Context) error {
	interval := ctx.Int("interval")
	cmdStr := ctx.String("cmd")
	fmt.Println("version cmd:", cmdStr)
	info := host.GetHostInfo(interval, cmdStr)
	fmt.Println(info)
	return nil
}
