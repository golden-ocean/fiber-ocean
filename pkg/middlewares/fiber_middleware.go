package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func FiberMiddleware(app *fiber.App) {
	// 压缩
	// app.Use(compress.New(compress.Config{
	// 	Level: compress.LevelBestSpeed, // 1
	// }))
	// 跨域
	app.Use(cors.New())
	// etag
	app.Use(etag.New())
	// 幂等性
	app.Use(idempotency.New())
	//限流
	//app.Use(limiter.New())
	// request-id
	app.Use(requestid.New())
	// logger
	app.Use(logger.New())
	// 缓存
	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return c.Query("refresh") == "true"
	// 	},
	// 	Expiration:   30 * time.Second,
	// 	CacheControl: true,
	// 	KeyGenerator: func(ctx *fiber.Ctx) string {
	// 		return utils.CopyString(ctx.OriginalURL())
	// 	},
	// }))
	// app.Use(pprof.New())
	app.Use(recover.New())
}
