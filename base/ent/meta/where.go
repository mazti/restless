// Code generated by entc, DO NOT EDIT.

package meta

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/mazti/restless/base/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Meta {
	return predicate.Meta(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Base applies equality check predicate on the "base" field. It's identical to BaseEQ.
func Base(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBase), v))
	},
	)
}

// Schema applies equality check predicate on the "schema" field. It's identical to SchemaEQ.
func Schema(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchema), v))
	},
	)
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	},
	)
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	},
	)
}

// BaseEQ applies the EQ predicate on the "base" field.
func BaseEQ(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBase), v))
	},
	)
}

// BaseNEQ applies the NEQ predicate on the "base" field.
func BaseNEQ(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBase), v))
	},
	)
}

// BaseIn applies the In predicate on the "base" field.
func BaseIn(vs ...string) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBase), v...))
	},
	)
}

// BaseNotIn applies the NotIn predicate on the "base" field.
func BaseNotIn(vs ...string) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBase), v...))
	},
	)
}

// BaseGT applies the GT predicate on the "base" field.
func BaseGT(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBase), v))
	},
	)
}

// BaseGTE applies the GTE predicate on the "base" field.
func BaseGTE(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBase), v))
	},
	)
}

// BaseLT applies the LT predicate on the "base" field.
func BaseLT(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBase), v))
	},
	)
}

// BaseLTE applies the LTE predicate on the "base" field.
func BaseLTE(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBase), v))
	},
	)
}

// BaseContains applies the Contains predicate on the "base" field.
func BaseContains(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBase), v))
	},
	)
}

// BaseHasPrefix applies the HasPrefix predicate on the "base" field.
func BaseHasPrefix(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBase), v))
	},
	)
}

// BaseHasSuffix applies the HasSuffix predicate on the "base" field.
func BaseHasSuffix(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBase), v))
	},
	)
}

// BaseEqualFold applies the EqualFold predicate on the "base" field.
func BaseEqualFold(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBase), v))
	},
	)
}

// BaseContainsFold applies the ContainsFold predicate on the "base" field.
func BaseContainsFold(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBase), v))
	},
	)
}

// SchemaEQ applies the EQ predicate on the "schema" field.
func SchemaEQ(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchema), v))
	},
	)
}

// SchemaNEQ applies the NEQ predicate on the "schema" field.
func SchemaNEQ(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSchema), v))
	},
	)
}

// SchemaIn applies the In predicate on the "schema" field.
func SchemaIn(vs ...string) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSchema), v...))
	},
	)
}

// SchemaNotIn applies the NotIn predicate on the "schema" field.
func SchemaNotIn(vs ...string) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSchema), v...))
	},
	)
}

// SchemaGT applies the GT predicate on the "schema" field.
func SchemaGT(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSchema), v))
	},
	)
}

// SchemaGTE applies the GTE predicate on the "schema" field.
func SchemaGTE(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSchema), v))
	},
	)
}

// SchemaLT applies the LT predicate on the "schema" field.
func SchemaLT(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSchema), v))
	},
	)
}

// SchemaLTE applies the LTE predicate on the "schema" field.
func SchemaLTE(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSchema), v))
	},
	)
}

// SchemaContains applies the Contains predicate on the "schema" field.
func SchemaContains(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSchema), v))
	},
	)
}

// SchemaHasPrefix applies the HasPrefix predicate on the "schema" field.
func SchemaHasPrefix(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSchema), v))
	},
	)
}

// SchemaHasSuffix applies the HasSuffix predicate on the "schema" field.
func SchemaHasSuffix(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSchema), v))
	},
	)
}

// SchemaEqualFold applies the EqualFold predicate on the "schema" field.
func SchemaEqualFold(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSchema), v))
	},
	)
}

// SchemaContainsFold applies the ContainsFold predicate on the "schema" field.
func SchemaContainsFold(v string) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSchema), v))
	},
	)
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	},
	)
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	},
	)
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	},
	)
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	},
	)
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	},
	)
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	},
	)
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Meta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Meta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	},
	)
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	},
	)
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	},
	)
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	},
	)
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Meta) predicate.Meta {
	return predicate.Meta(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Meta) predicate.Meta {
	return predicate.Meta(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Meta) predicate.Meta {
	return predicate.Meta(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
