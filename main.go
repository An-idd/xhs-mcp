package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xpzouying/xiaohongshu-mcp/configs"
)

func main() {
	var (
		headless     bool
		headlessMode string
		binPath      string // 浏览器二进制文件路径
		port         string
	)
	flag.BoolVar(&headless, "headless", false, "无头模式(默认有头,反检测更好;无显示器的服务器需显式 -headless=true)")
	flag.StringVar(&headlessMode, "headless-mode", "new", "无头模式: new(新无头,反检测更好,默认)|old(旧无头)，仅 headless=true 时生效")
	flag.StringVar(&binPath, "bin", "", "浏览器二进制文件路径")
	flag.StringVar(&port, "port", ":18060", "端口")
	flag.Parse()

	if len(binPath) == 0 {
		binPath = os.Getenv("ROD_BROWSER_BIN")
	}
	if binPath != "" {
		logrus.Infof("using browser binary: %s", binPath)
	} else {
		logrus.Infof("browser binary is not configured; rod will auto-detect or download Chromium")
	}

	configs.InitHeadless(headless)
	configs.SetHeadlessNew(headlessMode != "old")
	configs.SetBinPath(binPath)

	// 初始化服务
	xiaohongshuService := NewXiaohongshuService()

	// 创建并启动应用服务器
	appServer := NewAppServer(xiaohongshuService)
	if err := appServer.Start(port); err != nil {
		logrus.Fatalf("failed to run server: %v", err)
	}
}
