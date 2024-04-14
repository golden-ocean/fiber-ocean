package menu

const (
	Table = "system_menu"

	CreatedSuccess = "添加菜单成功！"
	CreatedFail    = "添加菜单失败！"
	UpdatedSuccess = "修改菜单成功！"
	UpdatedFail    = "修改菜单失败！"
	DeletedSuccess = "删除菜单成功！"
	DeletedFail    = "删除菜单失败！"

	ErrorNotExist                = "菜单或权限不存在！"
	ErrorPidNotExist             = "父级菜单或权限不存在！"
	ErrorNameRepeat              = "菜单名称重复，请重新输入！"
	ErrorExistChildren           = "菜单下存在子菜单，不能删除！"
	ErrorPidCantEqSelfAndChildId = "父节点不能为自己和其子节点，请重新选择父节点！"
)

var SelectFields = []string{
	"id",
	"name",
	"parent_id",
	"icon",
	"path",
	"permission",
	"component",
	"type",
	"method",
	"visible",
	"status",
	"sort",
	"remark",
}
