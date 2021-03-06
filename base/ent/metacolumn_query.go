// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mazti/restless/base/ent/metacolumn"
	"github.com/mazti/restless/base/ent/metatable"
	"github.com/mazti/restless/base/ent/predicate"
)

// MetaColumnQuery is the builder for querying MetaColumn entities.
type MetaColumnQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.MetaColumn
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (mcq *MetaColumnQuery) Where(ps ...predicate.MetaColumn) *MetaColumnQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit adds a limit step to the query.
func (mcq *MetaColumnQuery) Limit(limit int) *MetaColumnQuery {
	mcq.limit = &limit
	return mcq
}

// Offset adds an offset step to the query.
func (mcq *MetaColumnQuery) Offset(offset int) *MetaColumnQuery {
	mcq.offset = &offset
	return mcq
}

// Order adds an order step to the query.
func (mcq *MetaColumnQuery) Order(o ...Order) *MetaColumnQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// QueryTable chains the current query on the table edge.
func (mcq *MetaColumnQuery) QueryTable() *MetaTableQuery {
	query := &MetaTableQuery{config: mcq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(metacolumn.Table, metacolumn.FieldID, mcq.sqlQuery()),
		sqlgraph.To(metatable.Table, metatable.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, metacolumn.TableTable, metacolumn.TableColumn),
	)
	query.sql = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
	return query
}

// First returns the first MetaColumn entity in the query. Returns *ErrNotFound when no metacolumn was found.
func (mcq *MetaColumnQuery) First(ctx context.Context) (*MetaColumn, error) {
	mcs, err := mcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(mcs) == 0 {
		return nil, &ErrNotFound{metacolumn.Label}
	}
	return mcs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MetaColumnQuery) FirstX(ctx context.Context) *MetaColumn {
	mc, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return mc
}

// FirstID returns the first MetaColumn id in the query. Returns *ErrNotFound when no id was found.
func (mcq *MetaColumnQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{metacolumn.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mcq *MetaColumnQuery) FirstXID(ctx context.Context) int {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MetaColumn entity in the query, returns an error if not exactly one entity was returned.
func (mcq *MetaColumnQuery) Only(ctx context.Context) (*MetaColumn, error) {
	mcs, err := mcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(mcs) {
	case 1:
		return mcs[0], nil
	case 0:
		return nil, &ErrNotFound{metacolumn.Label}
	default:
		return nil, &ErrNotSingular{metacolumn.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MetaColumnQuery) OnlyX(ctx context.Context) *MetaColumn {
	mc, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return mc
}

// OnlyID returns the only MetaColumn id in the query, returns an error if not exactly one id was returned.
func (mcq *MetaColumnQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{metacolumn.Label}
	default:
		err = &ErrNotSingular{metacolumn.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (mcq *MetaColumnQuery) OnlyXID(ctx context.Context) int {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetaColumns.
func (mcq *MetaColumnQuery) All(ctx context.Context) ([]*MetaColumn, error) {
	return mcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MetaColumnQuery) AllX(ctx context.Context) []*MetaColumn {
	mcs, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return mcs
}

// IDs executes the query and returns a list of MetaColumn ids.
func (mcq *MetaColumnQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mcq.Select(metacolumn.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MetaColumnQuery) IDsX(ctx context.Context) []int {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MetaColumnQuery) Count(ctx context.Context) (int, error) {
	return mcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MetaColumnQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MetaColumnQuery) Exist(ctx context.Context) (bool, error) {
	return mcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MetaColumnQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MetaColumnQuery) Clone() *MetaColumnQuery {
	return &MetaColumnQuery{
		config:     mcq.config,
		limit:      mcq.limit,
		offset:     mcq.offset,
		order:      append([]Order{}, mcq.order...),
		unique:     append([]string{}, mcq.unique...),
		predicates: append([]predicate.MetaColumn{}, mcq.predicates...),
		// clone intermediate query.
		sql: mcq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MetaColumn.Query().
//		GroupBy(metacolumn.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mcq *MetaColumnQuery) GroupBy(field string, fields ...string) *MetaColumnGroupBy {
	group := &MetaColumnGroupBy{config: mcq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = mcq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MetaColumn.Query().
//		Select(metacolumn.FieldName).
//		Scan(ctx, &v)
//
func (mcq *MetaColumnQuery) Select(field string, fields ...string) *MetaColumnSelect {
	selector := &MetaColumnSelect{config: mcq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = mcq.sqlQuery()
	return selector
}

func (mcq *MetaColumnQuery) sqlAll(ctx context.Context) ([]*MetaColumn, error) {
	var (
		nodes []*MetaColumn
		spec  = mcq.querySpec()
	)
	spec.ScanValues = func() []interface{} {
		node := &MetaColumn{config: mcq.config}
		nodes = append(nodes, node)
		return node.scanValues()
	}
	spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, spec); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (mcq *MetaColumnQuery) sqlCount(ctx context.Context) (int, error) {
	spec := mcq.querySpec()
	return sqlgraph.CountNodes(ctx, mcq.driver, spec)
}

func (mcq *MetaColumnQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mcq *MetaColumnQuery) querySpec() *sqlgraph.QuerySpec {
	spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metacolumn.Table,
			Columns: metacolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metacolumn.FieldID,
			},
		},
		From:   mcq.sql,
		Unique: true,
	}
	if ps := mcq.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.limit; limit != nil {
		spec.Limit = *limit
	}
	if offset := mcq.offset; offset != nil {
		spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return spec
}

func (mcq *MetaColumnQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(metacolumn.Table)
	selector := builder.Select(t1.Columns(metacolumn.Columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(metacolumn.Columns...)...)
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetaColumnGroupBy is the builder for group-by MetaColumn entities.
type MetaColumnGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MetaColumnGroupBy) Aggregate(fns ...Aggregate) *MetaColumnGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mcgb *MetaColumnGroupBy) Scan(ctx context.Context, v interface{}) error {
	return mcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcgb *MetaColumnGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MetaColumnGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MetaColumnGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcgb *MetaColumnGroupBy) StringsX(ctx context.Context) []string {
	v, err := mcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MetaColumnGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MetaColumnGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcgb *MetaColumnGroupBy) IntsX(ctx context.Context) []int {
	v, err := mcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MetaColumnGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MetaColumnGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcgb *MetaColumnGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mcgb *MetaColumnGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MetaColumnGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcgb *MetaColumnGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcgb *MetaColumnGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mcgb.sqlQuery().Query()
	if err := mcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mcgb *MetaColumnGroupBy) sqlQuery() *sql.Selector {
	selector := mcgb.sql
	columns := make([]string, 0, len(mcgb.fields)+len(mcgb.fns))
	columns = append(columns, mcgb.fields...)
	for _, fn := range mcgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mcgb.fields...)
}

// MetaColumnSelect is the builder for select fields of MetaColumn entities.
type MetaColumnSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (mcs *MetaColumnSelect) Scan(ctx context.Context, v interface{}) error {
	return mcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcs *MetaColumnSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mcs *MetaColumnSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MetaColumnSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcs *MetaColumnSelect) StringsX(ctx context.Context) []string {
	v, err := mcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mcs *MetaColumnSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MetaColumnSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcs *MetaColumnSelect) IntsX(ctx context.Context) []int {
	v, err := mcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mcs *MetaColumnSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MetaColumnSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcs *MetaColumnSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mcs *MetaColumnSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MetaColumnSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcs *MetaColumnSelect) BoolsX(ctx context.Context) []bool {
	v, err := mcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcs *MetaColumnSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mcs.sqlQuery().Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mcs *MetaColumnSelect) sqlQuery() sql.Querier {
	selector := mcs.sql
	selector.Select(selector.Columns(mcs.fields...)...)
	return selector
}
