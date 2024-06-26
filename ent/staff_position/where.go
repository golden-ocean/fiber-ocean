// Code generated by ent, DO NOT EDIT.

package staff_position

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldUpdatedBy, v))
}

// StaffID applies equality check predicate on the "staff_id" field. It's identical to StaffIDEQ.
func StaffID(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldStaffID, v))
}

// PositionID applies equality check predicate on the "position_id" field. It's identical to PositionIDEQ.
func PositionID(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldPositionID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// StaffIDEQ applies the EQ predicate on the "staff_id" field.
func StaffIDEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldStaffID, v))
}

// StaffIDNEQ applies the NEQ predicate on the "staff_id" field.
func StaffIDNEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldStaffID, v))
}

// StaffIDIn applies the In predicate on the "staff_id" field.
func StaffIDIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldStaffID, vs...))
}

// StaffIDNotIn applies the NotIn predicate on the "staff_id" field.
func StaffIDNotIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldStaffID, vs...))
}

// StaffIDGT applies the GT predicate on the "staff_id" field.
func StaffIDGT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldStaffID, v))
}

// StaffIDGTE applies the GTE predicate on the "staff_id" field.
func StaffIDGTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldStaffID, v))
}

// StaffIDLT applies the LT predicate on the "staff_id" field.
func StaffIDLT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldStaffID, v))
}

// StaffIDLTE applies the LTE predicate on the "staff_id" field.
func StaffIDLTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldStaffID, v))
}

// StaffIDContains applies the Contains predicate on the "staff_id" field.
func StaffIDContains(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContains(FieldStaffID, v))
}

// StaffIDHasPrefix applies the HasPrefix predicate on the "staff_id" field.
func StaffIDHasPrefix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasPrefix(FieldStaffID, v))
}

// StaffIDHasSuffix applies the HasSuffix predicate on the "staff_id" field.
func StaffIDHasSuffix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasSuffix(FieldStaffID, v))
}

// StaffIDEqualFold applies the EqualFold predicate on the "staff_id" field.
func StaffIDEqualFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEqualFold(FieldStaffID, v))
}

// StaffIDContainsFold applies the ContainsFold predicate on the "staff_id" field.
func StaffIDContainsFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContainsFold(FieldStaffID, v))
}

// PositionIDEQ applies the EQ predicate on the "position_id" field.
func PositionIDEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEQ(FieldPositionID, v))
}

// PositionIDNEQ applies the NEQ predicate on the "position_id" field.
func PositionIDNEQ(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNEQ(FieldPositionID, v))
}

// PositionIDIn applies the In predicate on the "position_id" field.
func PositionIDIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldIn(FieldPositionID, vs...))
}

// PositionIDNotIn applies the NotIn predicate on the "position_id" field.
func PositionIDNotIn(vs ...string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldNotIn(FieldPositionID, vs...))
}

// PositionIDGT applies the GT predicate on the "position_id" field.
func PositionIDGT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGT(FieldPositionID, v))
}

// PositionIDGTE applies the GTE predicate on the "position_id" field.
func PositionIDGTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldGTE(FieldPositionID, v))
}

// PositionIDLT applies the LT predicate on the "position_id" field.
func PositionIDLT(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLT(FieldPositionID, v))
}

// PositionIDLTE applies the LTE predicate on the "position_id" field.
func PositionIDLTE(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldLTE(FieldPositionID, v))
}

// PositionIDContains applies the Contains predicate on the "position_id" field.
func PositionIDContains(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContains(FieldPositionID, v))
}

// PositionIDHasPrefix applies the HasPrefix predicate on the "position_id" field.
func PositionIDHasPrefix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasPrefix(FieldPositionID, v))
}

// PositionIDHasSuffix applies the HasSuffix predicate on the "position_id" field.
func PositionIDHasSuffix(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldHasSuffix(FieldPositionID, v))
}

// PositionIDEqualFold applies the EqualFold predicate on the "position_id" field.
func PositionIDEqualFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldEqualFold(FieldPositionID, v))
}

// PositionIDContainsFold applies the ContainsFold predicate on the "position_id" field.
func PositionIDContainsFold(v string) predicate.Staff_Position {
	return predicate.Staff_Position(sql.FieldContainsFold(FieldPositionID, v))
}

// HasStaffs applies the HasEdge predicate on the "staffs" edge.
func HasStaffs() predicate.Staff_Position {
	return predicate.Staff_Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StaffsTable, StaffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaffsWith applies the HasEdge predicate on the "staffs" edge with a given conditions (other predicates).
func HasStaffsWith(preds ...predicate.Staff) predicate.Staff_Position {
	return predicate.Staff_Position(func(s *sql.Selector) {
		step := newStaffsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPositions applies the HasEdge predicate on the "positions" edge.
func HasPositions() predicate.Staff_Position {
	return predicate.Staff_Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionsTable, PositionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionsWith applies the HasEdge predicate on the "positions" edge with a given conditions (other predicates).
func HasPositionsWith(preds ...predicate.Position) predicate.Staff_Position {
	return predicate.Staff_Position(func(s *sql.Selector) {
		step := newPositionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Staff_Position) predicate.Staff_Position {
	return predicate.Staff_Position(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Staff_Position) predicate.Staff_Position {
	return predicate.Staff_Position(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Staff_Position) predicate.Staff_Position {
	return predicate.Staff_Position(sql.NotPredicates(p))
}
