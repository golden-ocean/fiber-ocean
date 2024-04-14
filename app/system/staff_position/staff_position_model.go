package staff_position

type CreateInput struct {
	StaffID     string   `zh:"用户ID" json:"staff_id,omitempty"`
	PositionIDs []string `zh:"职位ID列表" json:"position_ids,omitempty"`
}

type UpdateInput struct {
	StaffID    string `zh:"用户ID" json:"staff_id,omitempty"`
	PositionID string `zh:"职位ID" json:"position_id,omitempty"`
}

type DeleteInput struct {
	StaffID     string   `zh:"用户ID" json:"staff_id" validate:"required"`
	PositionIDs []string `zh:"职位ID列表" json:"position_ids" validate:"required"`
}

type WhereParams struct {
	StaffID    string   `zh:"用户ID" query:"staff_id" json:"staff_id" validate:"omitempty"`
	StaffIDs   []string `zh:"用户ID列表" query:"staff_ids" json:"staff_ids" validate:"omitempty"`
	PositionID string   `zh:"职位ID" query:"position_id" json:"position_id" validate:"omitempty"`
}

// type WhereParams struct {
// 	StaffID    string `zh:"员工ID" query:"staff_id" json:"staff_id" validate:"omitempty"`
// 	PositionID string `zh:"职位ID" query:"position_id" json:"position_id" validate:"omitempty"`
// }
