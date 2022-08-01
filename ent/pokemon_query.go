// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tromesh/go-ent-pokemon/ent/battle"
	"github.com/tromesh/go-ent-pokemon/ent/pokemon"
	"github.com/tromesh/go-ent-pokemon/ent/predicate"
)

// PokemonQuery is the builder for querying Pokemon entities.
type PokemonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Pokemon
	// eager-loading edges.
	withContender *BattleQuery
	withOpponent  *BattleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PokemonQuery builder.
func (pq *PokemonQuery) Where(ps ...predicate.Pokemon) *PokemonQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PokemonQuery) Limit(limit int) *PokemonQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PokemonQuery) Offset(offset int) *PokemonQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PokemonQuery) Unique(unique bool) *PokemonQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PokemonQuery) Order(o ...OrderFunc) *PokemonQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryContender chains the current query on the "contender" edge.
func (pq *PokemonQuery) QueryContender() *BattleQuery {
	query := &BattleQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemon.Table, pokemon.FieldID, selector),
			sqlgraph.To(battle.Table, battle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pokemon.ContenderTable, pokemon.ContenderColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOpponent chains the current query on the "opponent" edge.
func (pq *PokemonQuery) QueryOpponent() *BattleQuery {
	query := &BattleQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemon.Table, pokemon.FieldID, selector),
			sqlgraph.To(battle.Table, battle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pokemon.OpponentTable, pokemon.OpponentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pokemon entity from the query.
// Returns a *NotFoundError when no Pokemon was found.
func (pq *PokemonQuery) First(ctx context.Context) (*Pokemon, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pokemon.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PokemonQuery) FirstX(ctx context.Context) *Pokemon {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pokemon ID from the query.
// Returns a *NotFoundError when no Pokemon ID was found.
func (pq *PokemonQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pokemon.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PokemonQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pokemon entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pokemon entity is found.
// Returns a *NotFoundError when no Pokemon entities are found.
func (pq *PokemonQuery) Only(ctx context.Context) (*Pokemon, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pokemon.Label}
	default:
		return nil, &NotSingularError{pokemon.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PokemonQuery) OnlyX(ctx context.Context) *Pokemon {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pokemon ID in the query.
// Returns a *NotSingularError when more than one Pokemon ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PokemonQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = &NotSingularError{pokemon.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PokemonQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pokemons.
func (pq *PokemonQuery) All(ctx context.Context) ([]*Pokemon, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PokemonQuery) AllX(ctx context.Context) []*Pokemon {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pokemon IDs.
func (pq *PokemonQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(pokemon.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PokemonQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PokemonQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PokemonQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PokemonQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PokemonQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PokemonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PokemonQuery) Clone() *PokemonQuery {
	if pq == nil {
		return nil
	}
	return &PokemonQuery{
		config:        pq.config,
		limit:         pq.limit,
		offset:        pq.offset,
		order:         append([]OrderFunc{}, pq.order...),
		predicates:    append([]predicate.Pokemon{}, pq.predicates...),
		withContender: pq.withContender.Clone(),
		withOpponent:  pq.withOpponent.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithContender tells the query-builder to eager-load the nodes that are connected to
// the "contender" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonQuery) WithContender(opts ...func(*BattleQuery)) *PokemonQuery {
	query := &BattleQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withContender = query
	return pq
}

// WithOpponent tells the query-builder to eager-load the nodes that are connected to
// the "opponent" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonQuery) WithOpponent(opts ...func(*BattleQuery)) *PokemonQuery {
	query := &BattleQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOpponent = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pokemon.Query().
//		GroupBy(pokemon.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PokemonQuery) GroupBy(field string, fields ...string) *PokemonGroupBy {
	group := &PokemonGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Pokemon.Query().
//		Select(pokemon.FieldName).
//		Scan(ctx, &v)
//
func (pq *PokemonQuery) Select(fields ...string) *PokemonSelect {
	pq.fields = append(pq.fields, fields...)
	return &PokemonSelect{PokemonQuery: pq}
}

func (pq *PokemonQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !pokemon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PokemonQuery) sqlAll(ctx context.Context) ([]*Pokemon, error) {
	var (
		nodes       = []*Pokemon{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withContender != nil,
			pq.withOpponent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Pokemon{config: pq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withContender; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Pokemon)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Contender = []*Battle{}
		}
		query.withFKs = true
		query.Where(predicate.Battle(func(s *sql.Selector) {
			s.Where(sql.InValues(pokemon.ContenderColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.pokemon_contender
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "pokemon_contender" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pokemon_contender" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Contender = append(node.Edges.Contender, n)
		}
	}

	if query := pq.withOpponent; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Pokemon)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Opponent = []*Battle{}
		}
		query.withFKs = true
		query.Where(predicate.Battle(func(s *sql.Selector) {
			s.Where(sql.InValues(pokemon.OpponentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.pokemon_opponent
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "pokemon_opponent" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pokemon_opponent" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Opponent = append(node.Edges.Opponent, n)
		}
	}

	return nodes, nil
}

func (pq *PokemonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PokemonQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PokemonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pokemon.Table,
			Columns: pokemon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pokemon.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pokemon.FieldID)
		for i := range fields {
			if fields[i] != pokemon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PokemonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pokemon.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = pokemon.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PokemonGroupBy is the group-by builder for Pokemon entities.
type PokemonGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PokemonGroupBy) Aggregate(fns ...AggregateFunc) *PokemonGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PokemonGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PokemonGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PokemonGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PokemonGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PokemonGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PokemonGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PokemonGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PokemonGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PokemonGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PokemonGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PokemonGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !pokemon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PokemonGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PokemonSelect is the builder for selecting fields of Pokemon entities.
type PokemonSelect struct {
	*PokemonQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PokemonSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PokemonQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PokemonSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PokemonSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PokemonSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PokemonSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PokemonSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PokemonSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PokemonSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PokemonSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ps *PokemonSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemon.Label}
	default:
		err = fmt.Errorf("ent: PokemonSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PokemonSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PokemonSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
