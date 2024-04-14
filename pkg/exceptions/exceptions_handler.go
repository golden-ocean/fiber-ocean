package exceptions

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

func Handler(ctx fiber.Ctx, err error) error {
	// Status code defaults to 500
	host := ctx.Hostname()
	url := ctx.OriginalURL()
	// Retrieve the custom status code if it's a fiber.*Error
	if e, ok := err.(*Error); ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			Error{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				TraceId: url,
				Host:    host,
			},
		)
	}
	if e, ok := err.(*fiber.Error); ok {
		code := e.Code
		return ctx.Status(code).JSON(
			Error{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				TraceId: url,
				Host:    host,
			},
		)
	}
	if e, ok := err.(validator.ValidationErrors); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			Error{
				Success: false,
				Code:    fiber.StatusBadRequest,
				Message: utils.ValidatorErrors(e)[0],
				TraceId: url,
				Host:    host,
			},
		)
	}
	// Send custom error page
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			Error{
				Success: false,
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				TraceId: url,
				Host:    host,
			},
		)
	}

	// Return from handler
	return nil
}
