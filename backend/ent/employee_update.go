// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/patiwatpanda/app/ent/employee"
	"github.com/patiwatpanda/app/ent/pillorder"
	"github.com/patiwatpanda/app/ent/predicate"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks      []Hook
	mutation   *EmployeeMutation
	predicates []predicate.Employee
}

// Where adds a new predicate for the builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetEmployeeName sets the employee_name field.
func (eu *EmployeeUpdate) SetEmployeeName(s string) *EmployeeUpdate {
	eu.mutation.SetEmployeeName(s)
	return eu
}

// SetEmployeeEmail sets the employee_email field.
func (eu *EmployeeUpdate) SetEmployeeEmail(s string) *EmployeeUpdate {
	eu.mutation.SetEmployeeEmail(s)
	return eu
}

// SetEmployeePassword sets the employee_password field.
func (eu *EmployeeUpdate) SetEmployeePassword(s string) *EmployeeUpdate {
	eu.mutation.SetEmployeePassword(s)
	return eu
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (eu *EmployeeUpdate) AddPillorderIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddPillorderIDs(ids...)
	return eu
}

// AddPillorders adds the pillorders edges to Pillorder.
func (eu *EmployeeUpdate) AddPillorders(p ...*Pillorder) *EmployeeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddPillorderIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (eu *EmployeeUpdate) RemovePillorderIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemovePillorderIDs(ids...)
	return eu
}

// RemovePillorders removes pillorders edges to Pillorder.
func (eu *EmployeeUpdate) RemovePillorders(p ...*Pillorder) *EmployeeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := eu.mutation.EmployeeName(); ok {
		if err := employee.EmployeeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "employee_name", err: fmt.Errorf("ent: validator failed for field \"employee_name\": %w", err)}
		}
	}
	if v, ok := eu.mutation.EmployeeEmail(); ok {
		if err := employee.EmployeeEmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "employee_email", err: fmt.Errorf("ent: validator failed for field \"employee_email\": %w", err)}
		}
	}
	if v, ok := eu.mutation.EmployeePassword(); ok {
		if err := employee.EmployeePasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "employee_password", err: fmt.Errorf("ent: validator failed for field \"employee_password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employee.Table,
			Columns: employee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.EmployeeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeName,
		})
	}
	if value, ok := eu.mutation.EmployeeEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeEmail,
		})
	}
	if value, ok := eu.mutation.EmployeePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeePassword,
		})
	}
	if nodes := eu.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PillordersTable,
			Columns: []string{employee.PillordersColumn},
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
	if nodes := eu.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PillordersTable,
			Columns: []string{employee.PillordersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetEmployeeName sets the employee_name field.
func (euo *EmployeeUpdateOne) SetEmployeeName(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmployeeName(s)
	return euo
}

// SetEmployeeEmail sets the employee_email field.
func (euo *EmployeeUpdateOne) SetEmployeeEmail(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmployeeEmail(s)
	return euo
}

// SetEmployeePassword sets the employee_password field.
func (euo *EmployeeUpdateOne) SetEmployeePassword(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmployeePassword(s)
	return euo
}

// AddPillorderIDs adds the pillorders edge to Pillorder by ids.
func (euo *EmployeeUpdateOne) AddPillorderIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddPillorderIDs(ids...)
	return euo
}

// AddPillorders adds the pillorders edges to Pillorder.
func (euo *EmployeeUpdateOne) AddPillorders(p ...*Pillorder) *EmployeeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddPillorderIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// RemovePillorderIDs removes the pillorders edge to Pillorder by ids.
func (euo *EmployeeUpdateOne) RemovePillorderIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemovePillorderIDs(ids...)
	return euo
}

// RemovePillorders removes pillorders edges to Pillorder.
func (euo *EmployeeUpdateOne) RemovePillorders(p ...*Pillorder) *EmployeeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemovePillorderIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	if v, ok := euo.mutation.EmployeeName(); ok {
		if err := employee.EmployeeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "employee_name", err: fmt.Errorf("ent: validator failed for field \"employee_name\": %w", err)}
		}
	}
	if v, ok := euo.mutation.EmployeeEmail(); ok {
		if err := employee.EmployeeEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "employee_email", err: fmt.Errorf("ent: validator failed for field \"employee_email\": %w", err)}
		}
	}
	if v, ok := euo.mutation.EmployeePassword(); ok {
		if err := employee.EmployeePasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "employee_password", err: fmt.Errorf("ent: validator failed for field \"employee_password\": %w", err)}
		}
	}

	var (
		err  error
		node *Employee
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	e, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (e *Employee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employee.Table,
			Columns: employee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Employee.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.EmployeeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeName,
		})
	}
	if value, ok := euo.mutation.EmployeeEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeEmail,
		})
	}
	if value, ok := euo.mutation.EmployeePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeePassword,
		})
	}
	if nodes := euo.mutation.RemovedPillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PillordersTable,
			Columns: []string{employee.PillordersColumn},
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
	if nodes := euo.mutation.PillordersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PillordersTable,
			Columns: []string{employee.PillordersColumn},
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
	e = &Employee{config: euo.config}
	_spec.Assign = e.assignValues
	_spec.ScanValues = e.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return e, nil
}
