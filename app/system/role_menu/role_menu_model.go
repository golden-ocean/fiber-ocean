package role_menu

type CreateInput struct {
	RoleID  string   `zh:"角色ID" json:"role_id,omitempty"`
	MenuIDs []string `zh:"菜单ID列表" json:"menu_ids,omitempty"`
}

type UpdateInput struct {
	RoleID string `zh:"角色ID" json:"role_id,omitempty"`
	MenuID string `zh:"菜单ID" json:"menu_id,omitempty"`
}

type DeleteInput struct {
	RoleID  string   `zh:"角色ID" json:"role_id" validate:"required"`
	MenuIDs []string `zh:"菜单ID列表" json:"menu_ids" validate:"required"`
}

type WhereParams struct {
	RoleID  string   `zh:"角色ID" query:"role_id" json:"role_id" validate:"omitempty"`
	RoleIDs []string `zh:"角色ID列表" query:"role_ids" json:"role_ids" validate:"omitempty"`
}
