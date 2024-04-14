// Code generated by ent, DO NOT EDIT.

package role

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// EdgeStaffs holds the string denoting the staffs edge name in mutations.
	EdgeStaffs = "staffs"
	// Table holds the table name of the role in the database.
	Table = "system_roles"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "system_roles_menus"
	// MenusInverseTable is the table name for the Role_Menu entity.
	// It exists in this package in order to avoid circular dependency with the "role_menu" package.
	MenusInverseTable = "system_roles_menus"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "role_id"
	// OrganizationsTable is the table that holds the organizations relation/edge.
	OrganizationsTable = "system_roles_organizations"
	// OrganizationsInverseTable is the table name for the Role_Organization entity.
	// It exists in this package in order to avoid circular dependency with the "role_organization" package.
	OrganizationsInverseTable = "system_roles_organizations"
	// OrganizationsColumn is the table column denoting the organizations relation/edge.
	OrganizationsColumn = "role_id"
	// StaffsTable is the table that holds the staffs relation/edge.
	StaffsTable = "system_staffs_roles"
	// StaffsInverseTable is the table name for the Staff_Role entity.
	// It exists in this package in order to avoid circular dependency with the "staff_role" package.
	StaffsInverseTable = "system_staffs_roles"
	// StaffsColumn is the table column denoting the staffs relation/edge.
	StaffsColumn = "role_id"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldStatus,
	FieldSort,
	FieldRemark,
	FieldName,
	FieldCode,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/golden-ocean/fiber-ocean/ent/runtime"
var (
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() int64
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Role queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByMenusCount orders the results by menus count.
func ByMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenusStep(), opts...)
	}
}

// ByMenus orders the results by menus terms.
func ByMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrganizationsCount orders the results by organizations count.
func ByOrganizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationsStep(), opts...)
	}
}

// ByOrganizations orders the results by organizations terms.
func ByOrganizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStaffsCount orders the results by staffs count.
func ByStaffsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStaffsStep(), opts...)
	}
}

// ByStaffs orders the results by staffs terms.
func ByStaffs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStaffsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MenusTable, MenusColumn),
	)
}
func newOrganizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrganizationsTable, OrganizationsColumn),
	)
}
func newStaffsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StaffsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StaffsTable, StaffsColumn),
	)
}
