package commands

import (
	"fmt"
	"github.com/Hopetree/GoMonitor/tool"
	"github.com/urfave/cli/v2"
)

var Decrypt = &cli.Command{
	Name:  "decrypt",
	Usage: "解密密钥，并显示原文",
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
	},
	Action: decryptCmd,
}

func decryptCmd(ctx *cli.Context) error {
	key := ctx.String("key")
	secret := ctx.String("secret")
	cipher := tool.NewAESCipher(key)
	info, err := cipher.Decrypt(secret)
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}
