// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"telegraph/storage_repo/ent/pageview"
	"telegraph/storage_repo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PageViewDelete is the builder for deleting a PageView entity.
type PageViewDelete struct {
	config
	hooks    []Hook
	mutation *PageViewMutation
}

// Where appends a list predicates to the PageViewDelete builder.
func (pvd *PageViewDelete) Where(ps ...predicate.PageView) *PageViewDelete {
	pvd.mutation.Where(ps...)
	return pvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pvd *PageViewDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pvd.hooks) == 0 {
		affected, err = pvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageViewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvd.mutation = mutation
			affected, err = pvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pvd.hooks) - 1; i >= 0; i-- {
			if pvd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvd *PageViewDelete) ExecX(ctx context.Context) int {
	n, err := pvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pvd *PageViewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: pageview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pageview.FieldID,
			},
		},
	}
	if ps := pvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PageViewDeleteOne is the builder for deleting a single PageView entity.
type PageViewDeleteOne struct {
	pvd *PageViewDelete
}

// Exec executes the deletion query.
func (pvdo *PageViewDeleteOne) Exec(ctx context.Context) error {
	n, err := pvdo.pvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pageview.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pvdo *PageViewDeleteOne) ExecX(ctx context.Context) {
	pvdo.pvd.ExecX(ctx)
}
