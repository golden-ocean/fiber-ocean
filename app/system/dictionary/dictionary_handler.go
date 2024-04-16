package dictionary

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/pkg/common/constants"
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

func (h *Handler) Create(c fiber.Ctx) error {
	r := &CreateInput{}
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := h.service.Create(r); err != nil {
		return err
	}
	return c.JSON(response.OK(CreatedSuccess))
}

func (h *Handler) Update(c fiber.Ctx) error {
	req := &UpdateInput{}
	if err := c.Bind().Body(req); err != nil {
		return err
	}
	if err := h.service.Update(req); err != nil {
		return err
	}
	return c.JSON(response.OK(UpdatedSuccess))
}

func (h *Handler) Delete(c fiber.Ctx) error {
	req := &DeleteInput{}
	if err := c.Bind().Body(req); err != nil {
		return err
	}
	if err := h.service.Delete(req); err != nil {
		return err
	}
	return c.JSON(response.OK(DeletedSuccess))
}

func (h *Handler) QueryPage(c fiber.Ctx) error {
	w := &WhereParams{
		Current:  constants.Current,
		PageSize: constants.PageSize,
	}
	if err := c.Bind().Query(w); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	es, total, err := h.service.QueryPage(w)
	if err != nil {
		return err
	}
	return c.JSON(response.Page(es, w.Current, w.PageSize, total))
}

func (h *Handler) QueryByCode(c fiber.Ctx) error {
	code := c.Params("code", "system_status")
	es, err := h.service.QueryByCode(code)
	if err != nil {
		return err
	}
	return c.JSON(response.OK(es))
}
