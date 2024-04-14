package position

type CreateInput struct {
	Name   string `zh:"职位名称" json:"name" validate:"required,min=2,max=32"`
	Code   string `zh:"职位编码" json:"code" validate:"required,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort   int32  `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateInput struct {
	ID     string `zh:"唯一标识符" json:"id" validate:"required"`
	Name   string `zh:"职位名称" json:"name" validate:"omitempty,min=2,max=32"`
	Code   string `zh:"职位编码" json:"code" validate:"omitempty,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort   int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"职位名称" query:"name" json:"name" validate:"omitempty,max=32"`
	Code     string `zh:"职位编码" query:"code" json:"code" validate:"omitempty,max=64"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize int    `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  int    `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type PositionOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
	Sort   int32  `json:"sort"`
	// Remark string `json:"remark"`
}
