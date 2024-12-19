// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ei-sugimoto/pokepoke/back/ent/deck"
	"github.com/ei-sugimoto/pokepoke/back/ent/predicate"
)

// DeckUpdate is the builder for updating Deck entities.
type DeckUpdate struct {
	config
	hooks    []Hook
	mutation *DeckMutation
}

// Where appends a list predicates to the DeckUpdate builder.
func (du *DeckUpdate) Where(ps ...predicate.Deck) *DeckUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetTitle sets the "title" field.
func (du *DeckUpdate) SetTitle(s string) *DeckUpdate {
	du.mutation.SetTitle(s)
	return du
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (du *DeckUpdate) SetNillableTitle(s *string) *DeckUpdate {
	if s != nil {
		du.SetTitle(*s)
	}
	return du
}

// SetDescription sets the "description" field.
func (du *DeckUpdate) SetDescription(s string) *DeckUpdate {
	du.mutation.SetDescription(s)
	return du
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (du *DeckUpdate) SetNillableDescription(s *string) *DeckUpdate {
	if s != nil {
		du.SetDescription(*s)
	}
	return du
}

// Mutation returns the DeckMutation object of the builder.
func (du *DeckUpdate) Mutation() *DeckMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeckUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeckUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeckUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeckUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(deck.Table, deck.Columns, sqlgraph.NewFieldSpec(deck.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Title(); ok {
		_spec.SetField(deck.FieldTitle, field.TypeString, value)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deck.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeckUpdateOne is the builder for updating a single Deck entity.
type DeckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeckMutation
}

// SetTitle sets the "title" field.
func (duo *DeckUpdateOne) SetTitle(s string) *DeckUpdateOne {
	duo.mutation.SetTitle(s)
	return duo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (duo *DeckUpdateOne) SetNillableTitle(s *string) *DeckUpdateOne {
	if s != nil {
		duo.SetTitle(*s)
	}
	return duo
}

// SetDescription sets the "description" field.
func (duo *DeckUpdateOne) SetDescription(s string) *DeckUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (duo *DeckUpdateOne) SetNillableDescription(s *string) *DeckUpdateOne {
	if s != nil {
		duo.SetDescription(*s)
	}
	return duo
}

// Mutation returns the DeckMutation object of the builder.
func (duo *DeckUpdateOne) Mutation() *DeckMutation {
	return duo.mutation
}

// Where appends a list predicates to the DeckUpdate builder.
func (duo *DeckUpdateOne) Where(ps ...predicate.Deck) *DeckUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeckUpdateOne) Select(field string, fields ...string) *DeckUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deck entity.
func (duo *DeckUpdateOne) Save(ctx context.Context) (*Deck, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeckUpdateOne) SaveX(ctx context.Context) *Deck {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeckUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeckUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeckUpdateOne) sqlSave(ctx context.Context) (_node *Deck, err error) {
	_spec := sqlgraph.NewUpdateSpec(deck.Table, deck.Columns, sqlgraph.NewFieldSpec(deck.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Deck.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deck.FieldID)
		for _, f := range fields {
			if !deck.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deck.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Title(); ok {
		_spec.SetField(deck.FieldTitle, field.TypeString, value)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
	}
	_node = &Deck{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deck.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
