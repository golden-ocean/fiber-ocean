// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent/organization"
	"github.com/golden-ocean/fiber-ocean/ent/role"
	"github.com/golden-ocean/fiber-ocean/ent/role_organization"
)

// Role_Organization is the model entity for the Role_Organization schema.
type Role_Organization struct {
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
	// 角色ID
	RoleID string `json:"role_id,omitempty"`
	// 组织ID
	OrganizationID string `json:"organization_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Role_OrganizationQuery when eager-loading is set.
	Edges        Role_OrganizationEdges `json:"-"`
	selectValues sql.SelectValues
}

// Role_OrganizationEdges holds the relations/edges for other nodes in the graph.
type Role_OrganizationEdges struct {
	// Roles holds the value of the roles edge.
	Roles *Role `json:"roles,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations *Organization `json:"organizations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Role_OrganizationEdges) RolesOrErr() (*Role, error) {
	if e.Roles != nil {
		return e.Roles, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: role.Label}
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Role_OrganizationEdges) OrganizationsOrErr() (*Organization, error) {
	if e.Organizations != nil {
		return e.Organizations, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role_Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role_organization.FieldCreatedAt, role_organization.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case role_organization.FieldID, role_organization.FieldCreatedBy, role_organization.FieldUpdatedBy, role_organization.FieldRoleID, role_organization.FieldOrganizationID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role_Organization fields.
func (ro *Role_Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role_organization.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ro.ID = value.String
			}
		case role_organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ro.CreatedAt = value.Int64
			}
		case role_organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ro.UpdatedAt = value.Int64
			}
		case role_organization.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ro.CreatedBy = value.String
			}
		case role_organization.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ro.UpdatedBy = value.String
			}
		case role_organization.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ro.RoleID = value.String
			}
		case role_organization.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				ro.OrganizationID = value.String
			}
		default:
			ro.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role_Organization.
// This includes values selected through modifiers, order, etc.
func (ro *Role_Organization) Value(name string) (ent.Value, error) {
	return ro.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Role_Organization entity.
func (ro *Role_Organization) QueryRoles() *RoleQuery {
	return NewRoleOrganizationClient(ro.config).QueryRoles(ro)
}

// QueryOrganizations queries the "organizations" edge of the Role_Organization entity.
func (ro *Role_Organization) QueryOrganizations() *OrganizationQuery {
	return NewRoleOrganizationClient(ro.config).QueryOrganizations(ro)
}

// Update returns a builder for updating this Role_Organization.
// Note that you need to call Role_Organization.Unwrap() before calling this method if this Role_Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (ro *Role_Organization) Update() *RoleOrganizationUpdateOne {
	return NewRoleOrganizationClient(ro.config).UpdateOne(ro)
}

// Unwrap unwraps the Role_Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ro *Role_Organization) Unwrap() *Role_Organization {
	_tx, ok := ro.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role_Organization is not a transactional entity")
	}
	ro.config.driver = _tx.drv
	return ro
}

// String implements the fmt.Stringer.
func (ro *Role_Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Role_Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ro.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ro.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ro.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ro.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ro.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(ro.RoleID)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(ro.OrganizationID)
	builder.WriteByte(')')
	return builder.String()
}

// Role_Organizations is a parsable slice of Role_Organization.
type Role_Organizations []*Role_Organization
