package staff

type CreateInput struct {
	Username       string   `zh:"用户名称" json:"username" validate:"required,min=4,max=64,alphanum"`
	Password       string   `zh:"密码" json:"password" validate:"required,min=6,max=64"`
	Name           string   `zh:"用户姓名" json:"name" validate:"required,min=2,max=32"`
	Email          string   `zh:"电子邮件" json:"email" validate:"required,email"`
	Mobile         string   `zh:"移动电话" json:"mobile" validate:"required,len=11"`
	Avatar         string   `zh:"用户头像" json:"avatar" validate:"omitempty"`
	Gender         string   `zh:"用户性别" json:"gender" validate:"omitempty"`
	OrganizationID string   `zh:"归属部门" json:"organization_id" validate:"omitempty"`
	WorkStatus     string   `zh:"在职状态" json:"work_status" validate:"omitempty"`
	Status         string   `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort           int32    `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark         string   `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	PositionIDs    []string `zh:"职位IDs" json:"position_ids" validate:"omitempty"`
	RoleIDs        []string `zh:"角色IDs" json:"role_ids" validate:"omitempty"`
}

type UpdateInput struct {
	ID             string   `zh:"唯一标识符" json:"id" validate:"required"`
	Username       string   `zh:"用户名称" json:"username" validate:"omitempty,min=4,max=64,alphanum"`
	Password       string   `zh:"密码" json:"password" validate:"omitempty,min=6,max=64"`
	Name           string   `zh:"用户姓名" json:"name" validate:"omitempty,min=2,max=32"`
	Email          string   `zh:"电子邮件" json:"email" validate:"omitempty,email"`
	Mobile         string   `zh:"移动电话" json:"mobile" validate:"omitempty,len=11"`
	Avatar         string   `zh:"用户头像" json:"avatar" validate:"omitempty"`
	Gender         string   `zh:"用户性别" json:"gender" validate:"omitempty,oneof=Male Female Unknown"`
	OrganizationID string   `zh:"归属部门ID" json:"organization_id" validate:"omitempty"`
	WorkStatus     string   `zh:"在职状态" json:"work_status" validate:"omitempty"`
	Status         string   `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort           int32    `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark         string   `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	PositionIDs    []string `zh:"职位IDs" json:"position_ids" validate:"omitempty"`
	RoleIDs        []string `zh:"角色IDs" json:"role_ids" validate:"omitempty"`
}

type DeleteInput struct {
	ID          string   `zh:"唯一标识符" json:"id" validate:"required"`
	PositionIDs []string `zh:"职位IDs" json:"position_ids" validate:"omitempty"`
	RoleIDs     []string `zh:"角色IDs" json:"role_ids" validate:"omitempty"`
}

type WhereParams struct {
	ID             string `zh:"唯一标识符" query:"id" json:"id" validate:"omitempty"`
	Username       string `zh:"用户名称" query:"username" json:"username" validate:"omitempty,max=64"`
	Name           string `zh:"真实姓名" query:"name" json:"name" validate:"omitempty,max=16"`
	Email          string `zh:"邮件地址" query:"email" json:"email" validate:"omitempty,max=32"`
	Mobile         string `zh:"电话号码" query:"mobile" json:"mobile" validate:"omitempty,max=11"`
	Gender         string `zh:"性别" query:"gender" json:"gender" validate:"omitempty"`
	OrganizationID string `zh:"部门" query:"organization_id" json:"organization_id" validate:"omitempty"`
	WorkStatus     string `zh:"在职状态" query:"work_status" json:"work_status" validate:"omitempty"`
	Status         string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark         string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize       int    `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current        int    `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type StaffOutput struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Mobile         string   `json:"mobile"`
	Avatar         string   `json:"avatar"`
	Gender         string   `json:"gender"`
	OrganizationID string   `json:"organization_id"`
	WorkStatus     string   `json:"work_status"`
	Status         string   `json:"status"`
	Sort           int32    `json:"sort"`
	Remark         string   `json:"remark"`
	PositionIDs    []string `json:"position_ids,omitempty"`
	RoleIDs        []string `json:"role_ids,omitempty"`
}

// func (output *StaffOutput) Edges(e ent.StaffEdges) {
// 	p_ids := lo.Map(e.StaffsPositions, func(item *ent.Staff_Position, _ int) string {
// 		return item.PositionID
// 	})
// 	r_ids := lo.Map(e.StaffsRoles, func(item *ent.Staff_Role, _ int) string {
// 		return item.RoleID
// 	})
// 	output.PositionIDs = p_ids
// 	output.RoleIDs = r_ids
// }
