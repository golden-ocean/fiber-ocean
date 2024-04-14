package menu

type CreateInput struct {
	Name       string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID   string `zh:"父级菜单" json:"parent_id" validate:"omitempty"`
	Icon       string `zh:"菜单图标" json:"icon"`
	Path       string `zh:"路径" json:"path"`
	Permission string `zh:"权限" json:"permission"`
	Component  string `zh:"组件" json:"component"`
	Type       string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method     string `zh:"方法" json:"method"`
	Visible    string `zh:"是否隐藏" json:"visible"`
	Status     string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort       uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark     string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateInput struct {
	ID         string `zh:"唯一标识符" json:"id" validate:"required"`
	Name       string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID   string `zh:"父级菜单" json:"parent_id" validate:"omitempty"`
	Icon       string `zh:"菜单图标" json:"icon"`
	Path       string `zh:"路径" json:"path"`
	Permission string `zh:"权限" json:"permission"`
	Component  string `zh:"组件" json:"component"`
	Type       string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method     string `zh:"方法" json:"method"`
	Visible    string `zh:"是否隐藏" json:"visible"`
	Status     string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort       uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark     string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"菜单名称" query:"name" json:"name" validate:"omitempty,max=32"`
	ParentID string `zh:"父级菜单" query:"parent_id" json:"parent_id" validate:"omitempty"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type MenuOutput struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	ParentID   string        `json:"parent_id"`
	Icon       string        `json:"icon"`
	Path       string        `json:"path"`
	Permission string        `json:"permission"`
	Component  string        `json:"component"`
	Type       string        `json:"type"`
	Method     string        `json:"method"`
	Visible    string        `json:"visible"`
	Status     string        `json:"status"`
	Sort       uint32        `json:"sort"`
	Remark     string        `json:"remark"`
	Children   []*MenuOutput `json:"children,omitempty"`
}
