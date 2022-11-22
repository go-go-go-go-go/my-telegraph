// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"telegraph/storage_repo/ent/account"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetShortName sets the "short_name" field.
func (ac *AccountCreate) SetShortName(s string) *AccountCreate {
	ac.mutation.SetShortName(s)
	return ac
}

// SetAuthorName sets the "author_name" field.
func (ac *AccountCreate) SetAuthorName(s string) *AccountCreate {
	ac.mutation.SetAuthorName(s)
	return ac
}

// SetNillableAuthorName sets the "author_name" field if the given value is not nil.
func (ac *AccountCreate) SetNillableAuthorName(s *string) *AccountCreate {
	if s != nil {
		ac.SetAuthorName(*s)
	}
	return ac
}

// SetAuthorURL sets the "author_url" field.
func (ac *AccountCreate) SetAuthorURL(s string) *AccountCreate {
	ac.mutation.SetAuthorURL(s)
	return ac
}

// SetNillableAuthorURL sets the "author_url" field if the given value is not nil.
func (ac *AccountCreate) SetNillableAuthorURL(s *string) *AccountCreate {
	if s != nil {
		ac.SetAuthorURL(*s)
	}
	return ac
}

// SetAccessToken sets the "access_token" field.
func (ac *AccountCreate) SetAccessToken(s string) *AccountCreate {
	ac.mutation.SetAccessToken(s)
	return ac
}

// SetAuthURL sets the "auth_url" field.
func (ac *AccountCreate) SetAuthURL(s string) *AccountCreate {
	ac.mutation.SetAuthURL(s)
	return ac
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (ac *AccountCreate) SetNillableAuthURL(s *string) *AccountCreate {
	if s != nil {
		ac.SetAuthURL(*s)
	}
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.AuthorName(); !ok {
		v := account.DefaultAuthorName
		ac.mutation.SetAuthorName(v)
	}
	if _, ok := ac.mutation.AuthorURL(); !ok {
		v := account.DefaultAuthorURL
		ac.mutation.SetAuthorURL(v)
	}
	if _, ok := ac.mutation.AuthURL(); !ok {
		v := account.DefaultAuthURL
		ac.mutation.SetAuthURL(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.ShortName(); !ok {
		return &ValidationError{Name: "short_name", err: errors.New(`ent: missing required field "Account.short_name"`)}
	}
	if v, ok := ac.mutation.ShortName(); ok {
		if err := account.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Account.short_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.AuthorName(); !ok {
		return &ValidationError{Name: "author_name", err: errors.New(`ent: missing required field "Account.author_name"`)}
	}
	if _, ok := ac.mutation.AuthorURL(); !ok {
		return &ValidationError{Name: "author_url", err: errors.New(`ent: missing required field "Account.author_url"`)}
	}
	if _, ok := ac.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "Account.access_token"`)}
	}
	if v, ok := ac.mutation.AccessToken(); ok {
		if err := account.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Account.access_token": %w`, err)}
		}
	}
	if _, ok := ac.mutation.AuthURL(); !ok {
		return &ValidationError{Name: "auth_url", err: errors.New(`ent: missing required field "Account.auth_url"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.ShortName(); ok {
		_spec.SetField(account.FieldShortName, field.TypeString, value)
		_node.ShortName = value
	}
	if value, ok := ac.mutation.AuthorName(); ok {
		_spec.SetField(account.FieldAuthorName, field.TypeString, value)
		_node.AuthorName = value
	}
	if value, ok := ac.mutation.AuthorURL(); ok {
		_spec.SetField(account.FieldAuthorURL, field.TypeString, value)
		_node.AuthorURL = value
	}
	if value, ok := ac.mutation.AccessToken(); ok {
		_spec.SetField(account.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := ac.mutation.AuthURL(); ok {
		_spec.SetField(account.FieldAuthURL, field.TypeString, value)
		_node.AuthURL = value
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
