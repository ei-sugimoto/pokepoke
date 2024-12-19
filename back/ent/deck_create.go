// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ei-sugimoto/pokepoke/back/ent/deck"
)

// DeckCreate is the builder for creating a Deck entity.
type DeckCreate struct {
	config
	mutation *DeckMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (dc *DeckCreate) SetTitle(s string) *DeckCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DeckCreate) SetDescription(s string) *DeckCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// Mutation returns the DeckMutation object of the builder.
func (dc *DeckCreate) Mutation() *DeckMutation {
	return dc.mutation
}

// Save creates the Deck in the database.
func (dc *DeckCreate) Save(ctx context.Context) (*Deck, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeckCreate) SaveX(ctx context.Context) *Deck {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeckCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeckCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeckCreate) check() error {
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Deck.title"`)}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Deck.description"`)}
	}
	return nil
}

func (dc *DeckCreate) sqlSave(ctx context.Context) (*Deck, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeckCreate) createSpec() (*Deck, *sqlgraph.CreateSpec) {
	var (
		_node = &Deck{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(deck.Table, sqlgraph.NewFieldSpec(deck.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.Title(); ok {
		_spec.SetField(deck.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// DeckCreateBulk is the builder for creating many Deck entities in bulk.
type DeckCreateBulk struct {
	config
	err      error
	builders []*DeckCreate
}

// Save creates the Deck entities in the database.
func (dcb *DeckCreateBulk) Save(ctx context.Context) ([]*Deck, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deck, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeckMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeckCreateBulk) SaveX(ctx context.Context) []*Deck {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeckCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeckCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
