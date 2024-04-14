package auth

import (
	"github.com/golden-ocean/fiber-ocean/app/system/menu"
	"github.com/golden-ocean/fiber-ocean/app/system/staff"
)

type LoginInput struct {
	Username string `zh:"用户名称" json:"username" validate:"required"`
	Password string `zh:"用户密码" json:"password" validate:"required"`
}

type LoginOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Access   string `json:"access"`
	Refresh  string `json:"refresh"`
	//ExpiredAt int64  `json:"expired_at"`
}

type RefreshInput struct {
	Refresh string `json:"refresh" validate:"required"`
}

type InfoOutput struct {
	Staff       *staff.StaffOutput `json:"staff"`
	Menus       []*menu.MenuOutput `json:"menus"`
	Permissions []string           `json:"permissions"`
}
