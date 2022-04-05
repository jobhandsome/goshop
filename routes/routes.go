package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Option func(engine *gin.Engine)

var Options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	Options = append(Options, opts...)
}

// 初始化
func Init() *gin.Engine {
	// 初始化日志
	gin.DisableConsoleColor()
	// 创建日志文件
	f, _ := os.Create("blogs.log")
	// 写入日志
	gin.DefaultWriter = io.MultiWriter(f)
	// 创建一个默认路由
	r := gin.Default()

	// 加载注册的app路由
	for _, opt := range Options {
		opt(r)
	}
	// 抛出指针
	return r
}
