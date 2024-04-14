package role

type CreateInput struct {
	Name   string `zh:"角色名称" json:"name" validate:"required,min=2,max=32"`
	Code   string `zh:"角色编码" json:"code" validate:"required,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable Unknown"`
	Sort   int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateInput struct {
	ID     string `zh:"唯一标识符" json:"id" validate:"required"`
	Name   string `zh:"角色名称" json:"name" validate:"omitempty,min=2,max=32"`
	Code   string `zh:"角色编码" json:"code" validate:"omitempty,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable Unknown"`
	Sort   int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"角色名称" query:"name" json:"name" validate:"omitempty,max=32"`
	Code     string `zh:"角色编码" query:"code" json:"code" validate:"omitempty,max=64"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable Unknown"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize int    `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  int    `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type RoleMenuInput struct {
	RoleID  string   `zh:"角色ID" json:"role_id" validate:"required"`
	MenuIDs []string `zh:"菜单ID列表" json:"menu_ids" validate:"omitempty"`
}

type RoleOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Sort   int32  `json:"sort,omitempty"`
	Remark string `json:"remark,omitempty"`
	// Menus  []*menu.MenuOutput `json:"menus,omitempty"`
	// Menus []*ent.Menu `json:"menus,omitempty"`
}

// func (output *RoleOutput) Edges(e ent.RoleEdges) {
// 	// m_ids := lo.Map(e.Menus, func(item *ent.Menu, _ int) string {
// 	// 	return item.ID
// 	// })
// 	output.Menus = e.Menus
// }


