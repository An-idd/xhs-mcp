package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// parseDotEnv 解析 .env 内容为键值对：跳过空行/注释，去掉可选的 export 前缀和两侧引号。
func parseDotEnv(r io.Reader) map[string]string {
	env := map[string]string{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.Trim(strings.TrimSpace(val), `"'`)
		if key != "" {
			env[key] = val
		}
	}
	return env
}

// loadDotEnv 读取当前目录 .env，把其中的键写入进程环境变量。
// 已存在的环境变量优先，不覆盖；无 .env 时静默跳过。
func loadDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return
	}
	defer f.Close()

	for k, v := range parseDotEnv(f) {
		if os.Getenv(k) == "" {
			_ = os.Setenv(k, v)
		}
	}
}
