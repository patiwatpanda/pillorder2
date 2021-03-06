// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/patiwatpanda/app/ent/dentist"
	"github.com/patiwatpanda/app/ent/pillorder"
	"github.com/patiwatpanda/app/ent/predicate"
)

// DentistUpdate is the builder for updating Dentist entities.
type DentistUpdate struct {
	config
	hooks      []Hook
	mutation   *DentistMutation
	predicates []predicate.Dentist
}

// Where adds a new predicate for the builder.
func (du *DentistUpdate) Where(ps ...predicate.Dentist) *DentistUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDentistName sets the Dentist_name field.
func (du *DentistUpdate) SetDentistName(s string) *DentistUpdate {
	du.mutation.SetDentistName(s)
	return du
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (du *DentistUpdate) AddPillorderIDs(ids ...int) *DentistUpdate {
	du.mutation.AddPillorderIDs(ids...)
	return du
}

// AddPillorders adds the pillorders edges to Pillorder.
func (du *DentistUpdate) AddPillorders(p ...*Pillorder) *DentistUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddPillorderIDs(ids...)
}

// Mutation returns the DentistMutation object of the builder.
func (du *DentistUpdate) Mutation() *DentistMutation {
	return du.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (du *DentistUpdate) RemovePillorderIDs(ids ...int) *DentistUpdate {
	du.mutation.RemovePillorderIDs(ids...)
	return du
}

// RemovePillorders removes pillorders edges to Pillorder.
func (du *DentistUpdate) RemovePillorders(p ...*Pillorder) *DentistUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DentistUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DentistName(); ok {
		if err := dentist.DentistNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Dentist_name", err: fmt.Errorf("ent: validator failed for field \"Dentist_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DentistUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DentistUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DentistUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DentistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentist.Table,
			Columns: dentist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentist.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DentistName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistName,
		})
	}
	if nodes := du.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.PillordersTable,
			Columns: []string{dentist.PillordersColumn},
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
	if nodes := du.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.PillordersTable,
			Columns: []string{dentist.PillordersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DentistUpdateOne is the builder for updating a single Dentist entity.
type DentistUpdateOne struct {
	config
	hooks    []Hook
	mutation *DentistMutation
}

// SetDentistName sets the Dentist_name field.
func (duo *DentistUpdateOne) SetDentistName(s string) *DentistUpdateOne {
	duo.mutation.SetDentistName(s)
	return duo
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (duo *DentistUpdateOne) AddPillorderIDs(ids ...int) *DentistUpdateOne {
	duo.mutation.AddPillorderIDs(ids...)
	return duo
}

// AddPillorders adds the pillorders edges to Pillorder.
func (duo *DentistUpdateOne) AddPillorders(p ...*Pillorder) *DentistUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddPillorderIDs(ids...)
}

// Mutation returns the DentistMutation object of the builder.
func (duo *DentistUpdateOne) Mutation() *DentistMutation {
	return duo.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (duo *DentistUpdateOne) RemovePillorderIDs(ids ...int) *DentistUpdateOne {
	duo.mutation.RemovePillorderIDs(ids...)
	return duo
}

// RemovePillorders removes pillorders edges to Pillorder.
func (duo *DentistUpdateOne) RemovePillorders(p ...*Pillorder) *DentistUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DentistUpdateOne) Save(ctx context.Context) (*Dentist, error) {
	if v, ok := duo.mutation.DentistName(); ok {
		if err := dentist.DentistNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Dentist_name", err: fmt.Errorf("ent: validator failed for field \"Dentist_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Dentist
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DentistUpdateOne) SaveX(ctx context.Context) *Dentist {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DentistUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DentistUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DentistUpdateOne) sqlSave(ctx context.Context) (d *Dentist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentist.Table,
			Columns: dentist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentist.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dentist.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DentistName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistName,
		})
	}
	if nodes := duo.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.PillordersTable,
			Columns: []string{dentist.PillordersColumn},
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
	if nodes := duo.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.PillordersTable,
			Columns: []string{dentist.PillordersColumn},
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
	d = &Dentist{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
