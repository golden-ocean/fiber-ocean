package configs

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/pkg/exceptions"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

func FiberConfig() fiber.Config {

	return fiber.Config{
		ServerHeader:                 "",                              // default: "" 服务报头
		StrictRouting:                false,                           // default: false 严格路由(是否区分 /foo 和 /foo/
		CaseSensitive:                false,                           // default: false 路由区分大小写
		Immutable:                    false,                           // default: false 上下文不可变
		UnescapePath:                 false,                           // default: false 转换路由编码字符
		BodyLimit:                    4 * 1024 * 1024,                 // default: 4 * 1024 * 1024 请求体大小
		Concurrency:                  256 * 1024,                      // default: 256 * 1024 并发数
		ReadTimeout:                  time.Second * time.Duration(60), // default: unlimited 读取请求允许的时间
		WriteTimeout:                 time.Second * time.Duration(60), // default: unlimited 写入响应允许的时间
		IdleTimeout:                  time.Second * time.Duration(60), // default: 60s 空闲连接超时时间
		ReadBufferSize:               4096,                            // default: 4096 读取请求体缓冲区大小
		WriteBufferSize:              4096,                            // default: 4096 写入响应体缓冲区大小
		CompressedFileSuffix:         ".ocean.gz",                     // default: ".fiber.gz" 压缩文件后缀
		GETOnly:                      false,                           // default: false 只允许GET请求
		ErrorHandler:                 exceptions.Handler,              // default: DefaultErrorHandler 错误处理
		DisableKeepalive:             false,                           // default: false 禁用keepalive
		DisableDefaultDate:           false,                           // default: false 禁用默认日期
		DisableDefaultContentType:    false,                           // default: false 禁用默认响应头
		DisableHeaderNormalizing:     false,                           // default: false 禁用响应头标准化
		AppName:                      "ocean_admin",                   // default: nil APP名称
		DisablePreParseMultipartForm: false,                           // default: true 禁用预解析多部分表单
		ReduceMemoryUsage:            false,                           // default: false 减少内存使用
		EnableIPValidation:           false,                           // default: false 启用IP验证
		StructValidator:              &utils.StructValidator{},        // 结构体验证
		//JSONEncoder:               sonic.Marshal,                   // json 编码
		//JSONDecoder:               sonic.Unmarshal,                 // json 解码
		// Prefork: false, // 同一端口多进程
		// ProxyHeader:               "X-Forwarded-Proto",             // default:  "" 代理标头
	}
}
