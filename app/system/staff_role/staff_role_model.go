package staff_role

type CreateInput struct {
	StaffID string   `zh:"用户ID" json:"staff_id,omitempty"`
	RoleIDs []string `zh:"角色ID列表" json:"role_ids,omitempty"`
}

type UpdateInput struct {
	StaffID string `zh:"用户ID" json:"staff_id,omitempty"`
	RoleID  string `zh:"角色ID" json:"role_id,omitempty"`
}

type DeleteInput struct {
	StaffID string   `zh:"用户ID" json:"staff_id" validate:"required"`
	RoleIDs []string `zh:"角色ID列表" json:"role_ids" validate:"required"`
}

type WhereParams struct {
	StaffID  string   `zh:"用户ID" query:"staff_id" json:"staff_id" validate:"omitempty"`
	StaffIDs []string `zh:"用户ID列表" query:"staff_ids" json:"staff_ids" validate:"omitempty"`
	RoleID   string   `zh:"角色ID" query:"role_id" json:"role_id" validate:"omitempty"`
}
