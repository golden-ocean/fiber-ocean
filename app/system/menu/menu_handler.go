package menu

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/pkg/common/response"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Create(ctx fiber.Ctx) error {
	r := &CreateInput{}
	if err := ctx.Bind().Body(r); err != nil {
		return err
	}
	if err := h.service.Create(r); err != nil {
		return err
	}
	return ctx.JSON(response.OK(CreatedSuccess))
}

func (h *Handler) Update(ctx fiber.Ctx) error {
	r := &UpdateInput{}
	if err := ctx.Bind().Body(r); err != nil {
		return err
	}
	if err := h.service.Update(r); err != nil {
		return err
	}
	return ctx.JSON(response.OK(UpdatedSuccess))
}

func (h *Handler) Delete(ctx fiber.Ctx) error {
	r := &DeleteInput{}
	if err := ctx.Bind().Body(r); err != nil {
		return err
	}
	if err := h.service.Delete(r); err != nil {
		return err
	}
	return ctx.JSON(response.OK(DeletedSuccess))
}

func (h *Handler) QueryTree(ctx fiber.Ctx) error {
	w := &WhereParams{}
	if err := ctx.Bind().Query(w); err != nil {
		return err
	}
	es, err := h.service.QueryTree(w)
	if err != nil {
		return err
	}
	return ctx.JSON(response.OK(es))
}
