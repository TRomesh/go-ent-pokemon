// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tromesh/go-ent-pokemon/ent/battle"
	"github.com/tromesh/go-ent-pokemon/ent/pokemon"
	"github.com/tromesh/go-ent-pokemon/ent/predicate"
)

// PokemonUpdate is the builder for updating Pokemon entities.
type PokemonUpdate struct {
	config
	hooks    []Hook
	mutation *PokemonMutation
}

// Where appends a list predicates to the PokemonUpdate builder.
func (pu *PokemonUpdate) Where(ps ...predicate.Pokemon) *PokemonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PokemonUpdate) SetName(s string) *PokemonUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PokemonUpdate) SetDescription(s string) *PokemonUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetWeight sets the "weight" field.
func (pu *PokemonUpdate) SetWeight(f float64) *PokemonUpdate {
	pu.mutation.ResetWeight()
	pu.mutation.SetWeight(f)
	return pu
}

// AddWeight adds f to the "weight" field.
func (pu *PokemonUpdate) AddWeight(f float64) *PokemonUpdate {
	pu.mutation.AddWeight(f)
	return pu
}

// SetHeight sets the "height" field.
func (pu *PokemonUpdate) SetHeight(f float64) *PokemonUpdate {
	pu.mutation.ResetHeight()
	pu.mutation.SetHeight(f)
	return pu
}

// AddHeight adds f to the "height" field.
func (pu *PokemonUpdate) AddHeight(f float64) *PokemonUpdate {
	pu.mutation.AddHeight(f)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PokemonUpdate) SetUpdatedAt(t time.Time) *PokemonUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PokemonUpdate) SetNillableUpdatedAt(t *time.Time) *PokemonUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// AddContenderIDs adds the "contender" edge to the Battle entity by IDs.
func (pu *PokemonUpdate) AddContenderIDs(ids ...int) *PokemonUpdate {
	pu.mutation.AddContenderIDs(ids...)
	return pu
}

// AddContender adds the "contender" edges to the Battle entity.
func (pu *PokemonUpdate) AddContender(b ...*Battle) *PokemonUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddContenderIDs(ids...)
}

// AddOpponentIDs adds the "opponent" edge to the Battle entity by IDs.
func (pu *PokemonUpdate) AddOpponentIDs(ids ...int) *PokemonUpdate {
	pu.mutation.AddOpponentIDs(ids...)
	return pu
}

// AddOpponent adds the "opponent" edges to the Battle entity.
func (pu *PokemonUpdate) AddOpponent(b ...*Battle) *PokemonUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddOpponentIDs(ids...)
}

// Mutation returns the PokemonMutation object of the builder.
func (pu *PokemonUpdate) Mutation() *PokemonMutation {
	return pu.mutation
}

// ClearContender clears all "contender" edges to the Battle entity.
func (pu *PokemonUpdate) ClearContender() *PokemonUpdate {
	pu.mutation.ClearContender()
	return pu
}

// RemoveContenderIDs removes the "contender" edge to Battle entities by IDs.
func (pu *PokemonUpdate) RemoveContenderIDs(ids ...int) *PokemonUpdate {
	pu.mutation.RemoveContenderIDs(ids...)
	return pu
}

// RemoveContender removes "contender" edges to Battle entities.
func (pu *PokemonUpdate) RemoveContender(b ...*Battle) *PokemonUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveContenderIDs(ids...)
}

// ClearOpponent clears all "opponent" edges to the Battle entity.
func (pu *PokemonUpdate) ClearOpponent() *PokemonUpdate {
	pu.mutation.ClearOpponent()
	return pu
}

// RemoveOpponentIDs removes the "opponent" edge to Battle entities by IDs.
func (pu *PokemonUpdate) RemoveOpponentIDs(ids ...int) *PokemonUpdate {
	pu.mutation.RemoveOpponentIDs(ids...)
	return pu
}

// RemoveOpponent removes "opponent" edges to Battle entities.
func (pu *PokemonUpdate) RemoveOpponent(b ...*Battle) *PokemonUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveOpponentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PokemonUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PokemonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PokemonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PokemonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PokemonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PokemonUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := pokemon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Pokemon.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := pokemon.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Pokemon.description": %w`, err)}
		}
	}
	return nil
}

func (pu *PokemonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pokemon.Table,
			Columns: pokemon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pokemon.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pokemon.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pokemon.FieldDescription,
		})
	}
	if value, ok := pu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldWeight,
		})
	}
	if value, ok := pu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldWeight,
		})
	}
	if value, ok := pu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldHeight,
		})
	}
	if value, ok := pu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldHeight,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pokemon.FieldUpdatedAt,
		})
	}
	if pu.mutation.ContenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedContenderIDs(); len(nodes) > 0 && !pu.mutation.ContenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ContenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OpponentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOpponentIDs(); len(nodes) > 0 && !pu.mutation.OpponentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OpponentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pokemon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PokemonUpdateOne is the builder for updating a single Pokemon entity.
type PokemonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PokemonMutation
}

// SetName sets the "name" field.
func (puo *PokemonUpdateOne) SetName(s string) *PokemonUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PokemonUpdateOne) SetDescription(s string) *PokemonUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetWeight sets the "weight" field.
func (puo *PokemonUpdateOne) SetWeight(f float64) *PokemonUpdateOne {
	puo.mutation.ResetWeight()
	puo.mutation.SetWeight(f)
	return puo
}

// AddWeight adds f to the "weight" field.
func (puo *PokemonUpdateOne) AddWeight(f float64) *PokemonUpdateOne {
	puo.mutation.AddWeight(f)
	return puo
}

// SetHeight sets the "height" field.
func (puo *PokemonUpdateOne) SetHeight(f float64) *PokemonUpdateOne {
	puo.mutation.ResetHeight()
	puo.mutation.SetHeight(f)
	return puo
}

// AddHeight adds f to the "height" field.
func (puo *PokemonUpdateOne) AddHeight(f float64) *PokemonUpdateOne {
	puo.mutation.AddHeight(f)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PokemonUpdateOne) SetUpdatedAt(t time.Time) *PokemonUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PokemonUpdateOne) SetNillableUpdatedAt(t *time.Time) *PokemonUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// AddContenderIDs adds the "contender" edge to the Battle entity by IDs.
func (puo *PokemonUpdateOne) AddContenderIDs(ids ...int) *PokemonUpdateOne {
	puo.mutation.AddContenderIDs(ids...)
	return puo
}

// AddContender adds the "contender" edges to the Battle entity.
func (puo *PokemonUpdateOne) AddContender(b ...*Battle) *PokemonUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddContenderIDs(ids...)
}

// AddOpponentIDs adds the "opponent" edge to the Battle entity by IDs.
func (puo *PokemonUpdateOne) AddOpponentIDs(ids ...int) *PokemonUpdateOne {
	puo.mutation.AddOpponentIDs(ids...)
	return puo
}

// AddOpponent adds the "opponent" edges to the Battle entity.
func (puo *PokemonUpdateOne) AddOpponent(b ...*Battle) *PokemonUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddOpponentIDs(ids...)
}

// Mutation returns the PokemonMutation object of the builder.
func (puo *PokemonUpdateOne) Mutation() *PokemonMutation {
	return puo.mutation
}

// ClearContender clears all "contender" edges to the Battle entity.
func (puo *PokemonUpdateOne) ClearContender() *PokemonUpdateOne {
	puo.mutation.ClearContender()
	return puo
}

// RemoveContenderIDs removes the "contender" edge to Battle entities by IDs.
func (puo *PokemonUpdateOne) RemoveContenderIDs(ids ...int) *PokemonUpdateOne {
	puo.mutation.RemoveContenderIDs(ids...)
	return puo
}

// RemoveContender removes "contender" edges to Battle entities.
func (puo *PokemonUpdateOne) RemoveContender(b ...*Battle) *PokemonUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveContenderIDs(ids...)
}

// ClearOpponent clears all "opponent" edges to the Battle entity.
func (puo *PokemonUpdateOne) ClearOpponent() *PokemonUpdateOne {
	puo.mutation.ClearOpponent()
	return puo
}

// RemoveOpponentIDs removes the "opponent" edge to Battle entities by IDs.
func (puo *PokemonUpdateOne) RemoveOpponentIDs(ids ...int) *PokemonUpdateOne {
	puo.mutation.RemoveOpponentIDs(ids...)
	return puo
}

// RemoveOpponent removes "opponent" edges to Battle entities.
func (puo *PokemonUpdateOne) RemoveOpponent(b ...*Battle) *PokemonUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveOpponentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PokemonUpdateOne) Select(field string, fields ...string) *PokemonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pokemon entity.
func (puo *PokemonUpdateOne) Save(ctx context.Context) (*Pokemon, error) {
	var (
		err  error
		node *Pokemon
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PokemonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PokemonUpdateOne) SaveX(ctx context.Context) *Pokemon {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PokemonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PokemonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PokemonUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := pokemon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Pokemon.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := pokemon.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Pokemon.description": %w`, err)}
		}
	}
	return nil
}

func (puo *PokemonUpdateOne) sqlSave(ctx context.Context) (_node *Pokemon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pokemon.Table,
			Columns: pokemon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pokemon.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pokemon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pokemon.FieldID)
		for _, f := range fields {
			if !pokemon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pokemon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pokemon.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pokemon.FieldDescription,
		})
	}
	if value, ok := puo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldWeight,
		})
	}
	if value, ok := puo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldWeight,
		})
	}
	if value, ok := puo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldHeight,
		})
	}
	if value, ok := puo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: pokemon.FieldHeight,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pokemon.FieldUpdatedAt,
		})
	}
	if puo.mutation.ContenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedContenderIDs(); len(nodes) > 0 && !puo.mutation.ContenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ContenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.ContenderTable,
			Columns: []string{pokemon.ContenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OpponentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOpponentIDs(); len(nodes) > 0 && !puo.mutation.OpponentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OpponentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pokemon.OpponentTable,
			Columns: []string{pokemon.OpponentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pokemon{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pokemon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
