package organization

type CreateInput struct {
	Name     string `zh:"组织名称" json:"name" validate:"required,min=2,max=32"`
	Code     string `zh:"组织编码" json:"code" validate:"required,min=2,max=64"`
	ParentID string `zh:"父级ID" json:"parent_id" validate:"omitempty"`
	Status   string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort     int32  `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark   string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateInput struct {
	ID       string `zh:"唯一标识符" json:"id" validate:"required"`
	Name     string `zh:"组织名称" json:"name" validate:"omitempty,min=2,max=32"`
	Code     string `zh:"组织编码" json:"code" validate:"omitempty,min=2,max=64"`
	ParentID string `zh:"父级组织" json:"parent_id" validate:"omitempty"`
	Status   string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort     int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark   string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	ID       string `zh:"唯一标识符" query:"id" json:"id" validate:"omitempty"`
	Name     string `zh:"组织名称" query:"name" json:"name" validate:"omitempty,max=32"`
	ParentID string `zh:"父级菜单" query:"parent_id" json:"parent_id" validate:"omitempty"`
	Code     string `zh:"组织编码" query:"code" json:"code" validate:"omitempty,max=64"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize int    `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  int    `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type OrganizationOutput struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Code     string                `json:"code"`
	ParentID string                `json:"parent_id"`
	Status   string                `json:"status"`
	Sort     int32                 `json:"sort"`
	Remark   string                `json:"remark"`
	Children []*OrganizationOutput `json:"children,omitempty"`
}
