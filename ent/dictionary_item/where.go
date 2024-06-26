// Code generated by ent, DO NOT EDIT.

package dictionary_item

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldUpdatedBy, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldSort, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldRemark, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldValue, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldColor, v))
}

// DictionaryID applies equality check predicate on the "dictionary_id" field. It's identical to DictionaryIDEQ.
func DictionaryID(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldDictionaryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldStatus, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldSort, v))
}

// SortIsNil applies the IsNil predicate on the "sort" field.
func SortIsNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIsNull(FieldSort))
}

// SortNotNil applies the NotNil predicate on the "sort" field.
func SortNotNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotNull(FieldSort))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldRemark, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldLabel, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldValue, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldColor, v))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldColor, v))
}

// DictionaryIDEQ applies the EQ predicate on the "dictionary_id" field.
func DictionaryIDEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEQ(FieldDictionaryID, v))
}

// DictionaryIDNEQ applies the NEQ predicate on the "dictionary_id" field.
func DictionaryIDNEQ(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNEQ(FieldDictionaryID, v))
}

// DictionaryIDIn applies the In predicate on the "dictionary_id" field.
func DictionaryIDIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldIn(FieldDictionaryID, vs...))
}

// DictionaryIDNotIn applies the NotIn predicate on the "dictionary_id" field.
func DictionaryIDNotIn(vs ...string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldNotIn(FieldDictionaryID, vs...))
}

// DictionaryIDGT applies the GT predicate on the "dictionary_id" field.
func DictionaryIDGT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGT(FieldDictionaryID, v))
}

// DictionaryIDGTE applies the GTE predicate on the "dictionary_id" field.
func DictionaryIDGTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldGTE(FieldDictionaryID, v))
}

// DictionaryIDLT applies the LT predicate on the "dictionary_id" field.
func DictionaryIDLT(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLT(FieldDictionaryID, v))
}

// DictionaryIDLTE applies the LTE predicate on the "dictionary_id" field.
func DictionaryIDLTE(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldLTE(FieldDictionaryID, v))
}

// DictionaryIDContains applies the Contains predicate on the "dictionary_id" field.
func DictionaryIDContains(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContains(FieldDictionaryID, v))
}

// DictionaryIDHasPrefix applies the HasPrefix predicate on the "dictionary_id" field.
func DictionaryIDHasPrefix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasPrefix(FieldDictionaryID, v))
}

// DictionaryIDHasSuffix applies the HasSuffix predicate on the "dictionary_id" field.
func DictionaryIDHasSuffix(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldHasSuffix(FieldDictionaryID, v))
}

// DictionaryIDEqualFold applies the EqualFold predicate on the "dictionary_id" field.
func DictionaryIDEqualFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldEqualFold(FieldDictionaryID, v))
}

// DictionaryIDContainsFold applies the ContainsFold predicate on the "dictionary_id" field.
func DictionaryIDContainsFold(v string) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.FieldContainsFold(FieldDictionaryID, v))
}

// HasDictionary applies the HasEdge predicate on the "dictionary" edge.
func HasDictionary() predicate.Dictionary_Item {
	return predicate.Dictionary_Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DictionaryTable, DictionaryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDictionaryWith applies the HasEdge predicate on the "dictionary" edge with a given conditions (other predicates).
func HasDictionaryWith(preds ...predicate.Dictionary) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(func(s *sql.Selector) {
		step := newDictionaryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Dictionary_Item) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Dictionary_Item) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Dictionary_Item) predicate.Dictionary_Item {
	return predicate.Dictionary_Item(sql.NotPredicates(p))
}
