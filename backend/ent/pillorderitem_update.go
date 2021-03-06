// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/patiwatpanda/app/ent/pillorder"
	"github.com/patiwatpanda/app/ent/pillorderitem"
	"github.com/patiwatpanda/app/ent/predicate"
)

// PillorderItemUpdate is the builder for updating PillorderItem entities.
type PillorderItemUpdate struct {
	config
	hooks      []Hook
	mutation   *PillorderItemMutation
	predicates []predicate.PillorderItem
}

// Where adds a new predicate for the builder.
func (piu *PillorderItemUpdate) Where(ps ...predicate.PillorderItem) *PillorderItemUpdate {
	piu.predicates = append(piu.predicates, ps...)
	return piu
}

// SetPillorderItemName sets the PillorderItem_name field.
func (piu *PillorderItemUpdate) SetPillorderItemName(s string) *PillorderItemUpdate {
	piu.mutation.SetPillorderItemName(s)
	return piu
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (piu *PillorderItemUpdate) AddPillorderIDs(ids ...int) *PillorderItemUpdate {
	piu.mutation.AddPillorderIDs(ids...)
	return piu
}

// AddPillorders adds the pillorders edges to Pillorder.
func (piu *PillorderItemUpdate) AddPillorders(p ...*Pillorder) *PillorderItemUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return piu.AddPillorderIDs(ids...)
}

// Mutation returns the PillorderItemMutation object of the builder.
func (piu *PillorderItemUpdate) Mutation() *PillorderItemMutation {
	return piu.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (piu *PillorderItemUpdate) RemovePillorderIDs(ids ...int) *PillorderItemUpdate {
	piu.mutation.RemovePillorderIDs(ids...)
	return piu
}

// RemovePillorders removes pillorders edges to Pillorder.
func (piu *PillorderItemUpdate) RemovePillorders(p ...*Pillorder) *PillorderItemUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return piu.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (piu *PillorderItemUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := piu.mutation.PillorderItemName(); ok {
		if err := pillorderitem.PillorderItemNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "PillorderItem_name", err: fmt.Errorf("ent: validator failed for field \"PillorderItem_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(piu.hooks) == 0 {
		affected, err = piu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PillorderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piu.mutation = mutation
			affected, err = piu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(piu.hooks) - 1; i >= 0; i-- {
			mut = piu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PillorderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PillorderItemUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PillorderItemUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (piu *PillorderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pillorderitem.Table,
			Columns: pillorderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pillorderitem.FieldID,
			},
		},
	}
	if ps := piu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.PillorderItemName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pillorderitem.FieldPillorderItemName,
		})
	}
	if nodes := piu.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pillorderitem.PillordersTable,
			Columns: []string{pillorderitem.PillordersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pillorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pillorderitem.PillordersTable,
			Columns: []string{pillorderitem.PillordersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pillorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pillorderitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PillorderItemUpdateOne is the builder for updating a single PillorderItem entity.
type PillorderItemUpdateOne struct {
	config
	hooks    []Hook
	mutation *PillorderItemMutation
}

// SetPillorderItemName sets the PillorderItem_name field.
func (piuo *PillorderItemUpdateOne) SetPillorderItemName(s string) *PillorderItemUpdateOne {
	piuo.mutation.SetPillorderItemName(s)
	return piuo
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (piuo *PillorderItemUpdateOne) AddPillorderIDs(ids ...int) *PillorderItemUpdateOne {
	piuo.mutation.AddPillorderIDs(ids...)
	return piuo
}

// AddPillorders adds the pillorders edges to Pillorder.
func (piuo *PillorderItemUpdateOne) AddPillorders(p ...*Pillorder) *PillorderItemUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return piuo.AddPillorderIDs(ids...)
}

// Mutation returns the PillorderItemMutation object of the builder.
func (piuo *PillorderItemUpdateOne) Mutation() *PillorderItemMutation {
	return piuo.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (piuo *PillorderItemUpdateOne) RemovePillorderIDs(ids ...int) *PillorderItemUpdateOne {
	piuo.mutation.RemovePillorderIDs(ids...)
	return piuo
}

// RemovePillorders removes pillorders edges to Pillorder.
func (piuo *PillorderItemUpdateOne) RemovePillorders(p ...*Pillorder) *PillorderItemUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return piuo.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (piuo *PillorderItemUpdateOne) Save(ctx context.Context) (*PillorderItem, error) {
	if v, ok := piuo.mutation.PillorderItemName(); ok {
		if err := pillorderitem.PillorderItemNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "PillorderItem_name", err: fmt.Errorf("ent: validator failed for field \"PillorderItem_name\": %w", err)}
		}
	}

	var (
		err  error
		node *PillorderItem
	)
	if len(piuo.hooks) == 0 {
		node, err = piuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PillorderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piuo.mutation = mutation
			node, err = piuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(piuo.hooks) - 1; i >= 0; i-- {
			mut = piuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PillorderItemUpdateOne) SaveX(ctx context.Context) *PillorderItem {
	pi, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pi
}

// Exec executes the query on the entity.
func (piuo *PillorderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PillorderItemUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (piuo *PillorderItemUpdateOne) sqlSave(ctx context.Context) (pi *PillorderItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pillorderitem.Table,
			Columns: pillorderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pillorderitem.FieldID,
			},
		},
	}
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PillorderItem.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := piuo.mutation.PillorderItemName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pillorderitem.FieldPillorderItemName,
		})
	}
	if nodes := piuo.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pillorderitem.PillordersTable,
			Columns: []string{pillorderitem.PillordersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pillorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pillorderitem.PillordersTable,
			Columns: []string{pillorderitem.PillordersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pillorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pi = &PillorderItem{config: piuo.config}
	_spec.Assign = pi.assignValues
	_spec.ScanValues = pi.scanValues()
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pillorderitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pi, nil
}
