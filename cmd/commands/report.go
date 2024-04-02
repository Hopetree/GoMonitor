package commands

import (
	"fmt"
	"github.com/Hopetree/GoMonitor/host"
	"github.com/Hopetree/GoMonitor/tool"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

var Report = &cli.Command{
	Name:  "report",
	Usage: "采集主机信息，并解析密钥，将信息上报到服务端，非调试模式不输出任何信息",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "解密Key",
		},
		&cli.StringFlag{
			Name:    "secret",
			Aliases: []string{"s"},
			Usage:   "待解密的密文",
		},
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
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "调试模式，开启则会输出采集信息，并上报一次信息到服务端",
		},
	},
	Action: reportCmd,
}

func reportCmd(ctx *cli.Context) error {
	key := ctx.String("key")
	secret := ctx.String("secret")
	interval := ctx.Int("interval")
	cmdStr := ctx.String("cmd")
	debug := ctx.Bool("debug")

	if interval < 3 {
		interval = 3
	}

	cipher := tool.NewAESCipher(key)
	decryptedText, err := cipher.Decrypt(secret)
	if err != nil {
		return err
	}

	httpInfo := strings.Split(decryptedText, "::")
	url := httpInfo[2]
	headers := map[string]string{
		"Push-Username": httpInfo[0],
		"Push-Password": httpInfo[1],
		"Push-Key":      key,
		"Push-Value":    secret,
	}

	for {
		jsonData := host.GetHostInfo(interval, cmdStr)

		result, err := tool.PushHttp(url, headers, jsonData)
		if err != nil {
			return err
		}

		if debug {
			fmt.Println("version cmd:", cmdStr)
			fmt.Println(jsonData)
			fmt.Println(decryptedText)
			fmt.Println(result)
			break
		}

		time.Sleep(time.Duration(interval-1) * time.Second)
	}
	return nil
}
