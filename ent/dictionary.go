// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary"
)

// Dictionary is the model entity for the Dictionary schema.
type Dictionary struct {
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
	// 字典名称
	Name string `json:"name,omitempty"`
	// 字典编码
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DictionaryQuery when eager-loading is set.
	Edges        DictionaryEdges `json:"-"`
	selectValues sql.SelectValues
}

// DictionaryEdges holds the relations/edges for other nodes in the graph.
type DictionaryEdges struct {
	// Items holds the value of the items edge.
	Items []*Dictionary_Item `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e DictionaryEdges) ItemsOrErr() ([]*Dictionary_Item, error) {
	if e.loadedTypes[0] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dictionary) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldCreatedAt, dictionary.FieldUpdatedAt, dictionary.FieldSort:
			values[i] = new(sql.NullInt64)
		case dictionary.FieldID, dictionary.FieldCreatedBy, dictionary.FieldUpdatedBy, dictionary.FieldStatus, dictionary.FieldRemark, dictionary.FieldName, dictionary.FieldCode:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dictionary fields.
func (d *Dictionary) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case dictionary.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Int64
			}
		case dictionary.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Int64
			}
		case dictionary.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				d.CreatedBy = value.String
			}
		case dictionary.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				d.UpdatedBy = value.String
			}
		case dictionary.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = value.String
			}
		case dictionary.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				d.Sort = int32(value.Int64)
			}
		case dictionary.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				d.Remark = value.String
			}
		case dictionary.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case dictionary.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				d.Code = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dictionary.
// This includes values selected through modifiers, order, etc.
func (d *Dictionary) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryItems queries the "items" edge of the Dictionary entity.
func (d *Dictionary) QueryItems() *DictionaryItemQuery {
	return NewDictionaryClient(d.config).QueryItems(d)
}

// Update returns a builder for updating this Dictionary.
// Note that you need to call Dictionary.Unwrap() before calling this method if this Dictionary
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dictionary) Update() *DictionaryUpdateOne {
	return NewDictionaryClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dictionary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dictionary) Unwrap() *Dictionary {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dictionary is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dictionary) String() string {
	var builder strings.Builder
	builder.WriteString("Dictionary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(d.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(d.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(d.Status)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", d.Sort))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(d.Remark)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(d.Code)
	builder.WriteByte(')')
	return builder.String()
}

// Dictionaries is a parsable slice of Dictionary.
type Dictionaries []*Dictionary
