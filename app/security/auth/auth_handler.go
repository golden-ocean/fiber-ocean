package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/app/system/staff"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/pkg/common/response"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Login(c fiber.Ctx) error {
	r := &LoginInput{}
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	e, err := h.service.Login(r)
	if err != nil {
		return err
	}
	t, err := utils.GenerateNewTokens(e.ID)
	if err != nil {
		return err
	}
	return c.JSON(response.OK(t))
}

// func (h *Handler) Logout(c fiber.Ctx) error {
// 	r := &UpdateInput{}
// 	if err := c.BodyParser(r); err != nil {
// 		return err
// 	}
// 	if err := utils.ValidateStruct(r); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err...)
// 	}
// 	if err := h.service.Update(r); err != nil {
// 		return err
// 	}
// 	return c.JSON(response.OK(UpdatedSuccess))
// }

func (h *Handler) QueryInfo(c fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	e, err := h.service.QueryInfo(claims.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(response.OK(e))
}

// func (h *Handler) Refresh(c fiber.Ctx) error {
// 	now := time.Now().Unix()
// 	r := &RefreshInput{}
// 	if err := c.BodyParser(r); err != nil {
// 		return err
// 	}
// 	refresh_time, err := utils.ParseRefreshToken(r.Refresh)
// 	if err != nil {
// 		return err
// 	}
// 	// 转到service 中，在redis中查找相应的过期时间，合适则返回并redis更新时间
// 	if refresh_time < now {
// 		return fiber.NewError(fiber.StatusUnauthorized, ErrorAccessExpiredTime)
// 	}
// 	jwt := c.Locals("jwt").(*jwt.Token)
// 	fmt.Println("jwt", jwt)
// 	id, _ := jwt.Claims.GetSubject()
// 	// id, err := xid.FromString(claims.ID)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// e, err := h.service.Refresh(id)
// 	// if err != nil {
// 	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	// }
// 	tokens, err := utils.GenerateNewTokens(id)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(response.OK(tokens))
// }

func (h *Handler) Test(c fiber.Ctx) error {
	j := c.Locals("id").(string)
	fmt.Println(j)
	es, err := global.Client.Staff.Query().Select(staff.SelectFields...).WithStaffsPositions().WithStaffsRoles().All(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(response.OK(es))
}
