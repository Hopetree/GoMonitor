package tool

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func PushHttp(url string, headers map[string]string, jsonData string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return "", fmt.Errorf("http req error: %s", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http get resp error: %s", err)
	}

	// 断开连接
	defer func() {
		resp.Body.Close()
	}()

	// 检查响应状态码是否为 200
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status is %d not 200", resp.StatusCode)
	}

	// 读取响应体的内容
	bytesResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read resp error: %s", err)
	}

	return string(bytesResponse), nil

}
