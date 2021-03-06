// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mazti/restless/base/ent/metatable"
	"github.com/mazti/restless/base/ent/predicate"
)

// MetaTableDelete is the builder for deleting a MetaTable entity.
type MetaTableDelete struct {
	config
	predicates []predicate.MetaTable
}

// Where adds a new predicate to the delete builder.
func (mtd *MetaTableDelete) Where(ps ...predicate.MetaTable) *MetaTableDelete {
	mtd.predicates = append(mtd.predicates, ps...)
	return mtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mtd *MetaTableDelete) Exec(ctx context.Context) (int, error) {
	return mtd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (mtd *MetaTableDelete) ExecX(ctx context.Context) int {
	n, err := mtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mtd *MetaTableDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: metatable.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metatable.FieldID,
			},
		},
	}
	if ps := mtd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mtd.driver, spec)
}

// MetaTableDeleteOne is the builder for deleting a single MetaTable entity.
type MetaTableDeleteOne struct {
	mtd *MetaTableDelete
}

// Exec executes the deletion query.
func (mtdo *MetaTableDeleteOne) Exec(ctx context.Context) error {
	n, err := mtdo.mtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{metatable.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mtdo *MetaTableDeleteOne) ExecX(ctx context.Context) {
	mtdo.mtd.ExecX(ctx)
}
