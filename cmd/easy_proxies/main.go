package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"easy_proxies/internal/app"
	"easy_proxies/internal/config"
	"easy_proxies/internal/logger"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "path to config file")
	flag.Parse()

	// 预初始化日志系统（使用默认 info 级别）
	logger.Init("info")

	cfg, err := config.Load(configPath)
	if err != nil {
		logger.Errorf("load config: %v", err)
		os.Exit(1)
	}

	// 根据配置设置日志级别
	logger.SetLevel(cfg.LogLevel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := app.Run(ctx, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "proxy pool exited with error: %v\n", err)
		os.Exit(1)
	}
}
