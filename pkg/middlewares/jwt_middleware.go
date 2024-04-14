package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/pkg/middlewares/jwtware"
	"github.com/spf13/viper"
)

// func JWTProtected() func(fiber.Ctx) error {
// 	config := pasetoware.Config{
// 		// SuccessHandler: jwtSuccess, // token有效后执行
// 		// ErrorHandler:   jwtError,   // token无效后执行
// 		SymmetricKey: []byte(viper.GetString("jwt.secret-key")),
// 		TokenPrefix:  "Bearer",
// 	}
// 	return pasetoware.New(config)
// }

func JWTProtected() func(fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("jwt.secret-key"))},
		ContextKey: "jwt",
		// ErrorHandler:   jwtError,   // token无效后执行
		// SuccessHandler: jwtSuccess, // token有效后执行
	}
	return jwtware.New(config)
}

// func jwtSuccess(c fiber.Ctx) error {
// 	claims, err := utils.ExtractTokenMetadata(c)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
// 	}
// 	c.Locals("id", claims.ID)
// 	return c.Next()
// }

// func jwtError(c fiber.Ctx, err error) error {
// 	if strings.Contains(err.Error(), "missing or malformed JWT") {
// 		return fiber.NewError(fiber.StatusBadRequest, "令牌缺失或错误！")
// 	}
// 	if strings.Contains(err.Error(), "token is expired") {
// 		return fiber.NewError(fiber.StatusBadRequest, "令牌已过期！")
// 	}
// 	return fiber.NewError(fiber.StatusUnauthorized, err.Error())
// }
