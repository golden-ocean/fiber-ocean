// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent/position"
	"github.com/golden-ocean/fiber-ocean/ent/staff"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
)

// Staff_Position is the model entity for the Staff_Position schema.
type Staff_Position struct {
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
	// 员工ID
	StaffID string `json:"staff_id,omitempty"`
	// 职位ID
	PositionID string `json:"position_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Staff_PositionQuery when eager-loading is set.
	Edges        Staff_PositionEdges `json:"-"`
	selectValues sql.SelectValues
}

// Staff_PositionEdges holds the relations/edges for other nodes in the graph.
type Staff_PositionEdges struct {
	// Staffs holds the value of the staffs edge.
	Staffs *Staff `json:"staffs,omitempty"`
	// Positions holds the value of the positions edge.
	Positions *Position `json:"positions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Staff_PositionEdges) StaffsOrErr() (*Staff, error) {
	if e.Staffs != nil {
		return e.Staffs, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: staff.Label}
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Staff_PositionEdges) PositionsOrErr() (*Position, error) {
	if e.Positions != nil {
		return e.Positions, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: position.Label}
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Staff_Position) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case staff_position.FieldCreatedAt, staff_position.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case staff_position.FieldID, staff_position.FieldCreatedBy, staff_position.FieldUpdatedBy, staff_position.FieldStaffID, staff_position.FieldPositionID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Staff_Position fields.
func (sp *Staff_Position) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case staff_position.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sp.ID = value.String
			}
		case staff_position.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Int64
			}
		case staff_position.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Int64
			}
		case staff_position.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sp.CreatedBy = value.String
			}
		case staff_position.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sp.UpdatedBy = value.String
			}
		case staff_position.FieldStaffID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field staff_id", values[i])
			} else if value.Valid {
				sp.StaffID = value.String
			}
		case staff_position.FieldPositionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position_id", values[i])
			} else if value.Valid {
				sp.PositionID = value.String
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Staff_Position.
// This includes values selected through modifiers, order, etc.
func (sp *Staff_Position) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// QueryStaffs queries the "staffs" edge of the Staff_Position entity.
func (sp *Staff_Position) QueryStaffs() *StaffQuery {
	return NewStaffPositionClient(sp.config).QueryStaffs(sp)
}

// QueryPositions queries the "positions" edge of the Staff_Position entity.
func (sp *Staff_Position) QueryPositions() *PositionQuery {
	return NewStaffPositionClient(sp.config).QueryPositions(sp)
}

// Update returns a builder for updating this Staff_Position.
// Note that you need to call Staff_Position.Unwrap() before calling this method if this Staff_Position
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *Staff_Position) Update() *StaffPositionUpdateOne {
	return NewStaffPositionClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the Staff_Position entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *Staff_Position) Unwrap() *Staff_Position {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: Staff_Position is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *Staff_Position) String() string {
	var builder strings.Builder
	builder.WriteString("Staff_Position(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", sp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", sp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(sp.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(sp.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("staff_id=")
	builder.WriteString(sp.StaffID)
	builder.WriteString(", ")
	builder.WriteString("position_id=")
	builder.WriteString(sp.PositionID)
	builder.WriteByte(')')
	return builder.String()
}

// Staff_Positions is a parsable slice of Staff_Position.
type Staff_Positions []*Staff_Position
