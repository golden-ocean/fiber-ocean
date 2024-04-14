// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent/organization"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 创建人员
	CreatedBy string `json:"created_by,omitempty"`
	// 更新人员
	UpdatedBy string `json:"updated_by,omitempty"`
	// 状态: Enable/Disable
	Status string `json:"status,omitempty"`
	// 排序
	Sort int32 `json:"sort,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 组织名称
	Name string `json:"name,omitempty"`
	// 组织编码
	Code string `json:"code,omitempty"`
	// 父级ID
	ParentID string `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"-"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Organization `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Organization `json:"children,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role_Organization `json:"roles,omitempty"`
	// Staffs holds the value of the staffs edge.
	Staffs []*Staff `json:"staffs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) ParentOrErr() (*Organization, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) ChildrenOrErr() ([]*Organization, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) RolesOrErr() ([]*Role_Organization, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) StaffsOrErr() ([]*Staff, error) {
	if e.loadedTypes[3] {
		return e.Staffs, nil
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldCreatedAt, organization.FieldUpdatedAt, organization.FieldSort:
			values[i] = new(sql.NullInt64)
		case organization.FieldID, organization.FieldCreatedBy, organization.FieldUpdatedBy, organization.FieldStatus, organization.FieldRemark, organization.FieldName, organization.FieldCode, organization.FieldParentID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = value.String
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Int64
			}
		case organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Int64
			}
		case organization.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				o.CreatedBy = value.String
			}
		case organization.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				o.UpdatedBy = value.String
			}
		case organization.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = value.String
			}
		case organization.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				o.Sort = int32(value.Int64)
			}
		case organization.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				o.Remark = value.String
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				o.Code = value.String
			}
		case organization.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				o.ParentID = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Organization entity.
func (o *Organization) QueryParent() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryParent(o)
}

// QueryChildren queries the "children" edge of the Organization entity.
func (o *Organization) QueryChildren() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryChildren(o)
}

// QueryRoles queries the "roles" edge of the Organization entity.
func (o *Organization) QueryRoles() *RoleOrganizationQuery {
	return NewOrganizationClient(o.config).QueryRoles(o)
}

// QueryStaffs queries the "staffs" edge of the Organization entity.
func (o *Organization) QueryStaffs() *StaffQuery {
	return NewOrganizationClient(o.config).QueryStaffs(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", o.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", o.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(o.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(o.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(o.Status)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", o.Sort))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(o.Remark)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(o.Code)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(o.ParentID)
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
