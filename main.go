package main

import (
	"github.com/gofiber/fiber/v3"
	_ "github.com/golden-ocean/fiber-ocean/ent/runtime"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/pkg/configs"
	"github.com/golden-ocean/fiber-ocean/pkg/middlewares"
	"github.com/golden-ocean/fiber-ocean/pkg/routes"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
	"github.com/golden-ocean/fiber-ocean/platform/database"

	"github.com/spf13/viper"
)

func main() {
	// 配置初始化
	configs.Init()
	// 连接数据库
	global.Client = database.OpenDBConnection()
	// 迁移数据
	database.Migration(global.Client)
	// fiber 自身配置
	fiberConfig := configs.FiberConfig()
	// 创建实例
	app := fiber.New(fiberConfig)
	// 中间件
	middlewares.FiberMiddleware(app) // 注册 fiber内置中间件
	// 路由.
	//routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if viper.GetString("server.status") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
