package host

import (
	"os/exec"
	"strings"
)

func getVersion(cmdStr string) string {
	var version string
	// 可以按需编辑这里的命令来获取服务的版本号
	cmd := exec.Command("sh", "-c", cmdStr)

	// 获取命令输出
	out, err := cmd.Output()
	if err == nil {
		version = strings.TrimSpace(string(out))
	}

	return version
}
