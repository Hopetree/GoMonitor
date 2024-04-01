package main

import (
	"fmt"
	"github.com/Hopetree/GoMonitor/host"
	"github.com/Hopetree/GoMonitor/tool"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	SecretKey   string
	SecretValue string
	Interval    = 3
	Debug       bool
)

func loadValues() {
	args := os.Args
	if len(args) == 1 {
		Debug = true
	} else if len(args) >= 4 {
		Debug = false
		// 获取前三个参数并分配给变量
		SecretKey = args[1]
		SecretValue = args[2]
		Interval, _ = strconv.Atoi(strings.TrimSpace(args[3]))
	} else {
		fmt.Println("Not enough command-line arguments provided")
		os.Exit(1)
	}
}

func main() {
	loadValues()
	// fmt.Println(SecretKey, SecretValue, Interval)

	for {
		// 采集信息
		jsonData := host.GetHostInfo(Interval)
		fmt.Println(jsonData)

		if Debug {
			fmt.Println("调试模式，打印结果并退出")
			os.Exit(0)
		}

		// 解析密钥，
		cipher := tool.NewAESCipher(SecretKey)
		decryptedText := cipher.Decrypt(SecretValue)
		httpInfo := strings.Split(decryptedText, "::")
		// fmt.Println(httpInfo)

		// 推送信息
		url := httpInfo[2]
		headers := map[string]string{
			"Push-Username": httpInfo[0],
			"Push-Password": httpInfo[1],
			"Push-Key":      SecretKey,
			"Push-Value":    SecretValue,
		}

		result, err := tool.PushHttp(url, headers, jsonData)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

		time.Sleep(time.Duration(Interval-1) * time.Second)
	}
}
