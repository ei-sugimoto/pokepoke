// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ei-sugimoto/pokepoke/back/ent/card"
	"github.com/ei-sugimoto/pokepoke/back/ent/deck"
	"github.com/ei-sugimoto/pokepoke/back/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCard = "Card"
	TypeDeck = "Deck"
)

// CardMutation represents an operation that mutates the Card nodes in the graph.
type CardMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	deck          *int
	cleareddeck   bool
	done          bool
	oldValue      func(context.Context) (*Card, error)
	predicates    []predicate.Card
}

var _ ent.Mutation = (*CardMutation)(nil)

// cardOption allows management of the mutation configuration using functional options.
type cardOption func(*CardMutation)

// newCardMutation creates new mutation for the Card entity.
func newCardMutation(c config, op Op, opts ...cardOption) *CardMutation {
	m := &CardMutation{
		config:        c,
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCardID sets the ID field of the mutation.
func withCardID(id int) cardOption {
	return func(m *CardMutation) {
		var (
			err   error
			once  sync.Once
			value *Card
		)
		m.oldValue = func(ctx context.Context) (*Card, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Card.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCard sets the old Card of the mutation.
func withCard(node *Card) cardOption {
	return func(m *CardMutation) {
		m.oldValue = func(context.Context) (*Card, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CardMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Card.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CardMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CardMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CardMutation) ResetName() {
	m.name = nil
}

// SetDeckID sets the "deck" edge to the Deck entity by id.
func (m *CardMutation) SetDeckID(id int) {
	m.deck = &id
}

// ClearDeck clears the "deck" edge to the Deck entity.
func (m *CardMutation) ClearDeck() {
	m.cleareddeck = true
}

// DeckCleared reports if the "deck" edge to the Deck entity was cleared.
func (m *CardMutation) DeckCleared() bool {
	return m.cleareddeck
}

// DeckID returns the "deck" edge ID in the mutation.
func (m *CardMutation) DeckID() (id int, exists bool) {
	if m.deck != nil {
		return *m.deck, true
	}
	return
}

// DeckIDs returns the "deck" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// DeckID instead. It exists only for internal usage by the builders.
func (m *CardMutation) DeckIDs() (ids []int) {
	if id := m.deck; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetDeck resets all changes to the "deck" edge.
func (m *CardMutation) ResetDeck() {
	m.deck = nil
	m.cleareddeck = false
}

// Where appends a list predicates to the CardMutation builder.
func (m *CardMutation) Where(ps ...predicate.Card) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CardMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CardMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Card, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CardMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CardMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, card.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case card.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CardMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case card.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Card field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case card.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CardMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CardMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CardMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.deck != nil {
		edges = append(edges, card.EdgeDeck)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CardMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case card.EdgeDeck:
		if id := m.deck; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CardMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareddeck {
		edges = append(edges, card.EdgeDeck)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	switch name {
	case card.EdgeDeck:
		return m.cleareddeck
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	switch name {
	case card.EdgeDeck:
		m.ClearDeck()
		return nil
	}
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	switch name {
	case card.EdgeDeck:
		m.ResetDeck()
		return nil
	}
	return fmt.Errorf("unknown Card edge %s", name)
}

// DeckMutation represents an operation that mutates the Deck nodes in the graph.
type DeckMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	description   *string
	clearedFields map[string]struct{}
	cards         map[int]struct{}
	removedcards  map[int]struct{}
	clearedcards  bool
	done          bool
	oldValue      func(context.Context) (*Deck, error)
	predicates    []predicate.Deck
}

var _ ent.Mutation = (*DeckMutation)(nil)

// deckOption allows management of the mutation configuration using functional options.
type deckOption func(*DeckMutation)

// newDeckMutation creates new mutation for the Deck entity.
func newDeckMutation(c config, op Op, opts ...deckOption) *DeckMutation {
	m := &DeckMutation{
		config:        c,
		op:            op,
		typ:           TypeDeck,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDeckID sets the ID field of the mutation.
func withDeckID(id int) deckOption {
	return func(m *DeckMutation) {
		var (
			err   error
			once  sync.Once
			value *Deck
		)
		m.oldValue = func(ctx context.Context) (*Deck, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Deck.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDeck sets the old Deck of the mutation.
func withDeck(node *Deck) deckOption {
	return func(m *DeckMutation) {
		m.oldValue = func(context.Context) (*Deck, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DeckMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DeckMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DeckMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DeckMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Deck.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *DeckMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *DeckMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Deck entity.
// If the Deck object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DeckMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *DeckMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *DeckMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *DeckMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Deck entity.
// If the Deck object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DeckMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *DeckMutation) ResetDescription() {
	m.description = nil
}

// AddCardIDs adds the "cards" edge to the Card entity by ids.
func (m *DeckMutation) AddCardIDs(ids ...int) {
	if m.cards == nil {
		m.cards = make(map[int]struct{})
	}
	for i := range ids {
		m.cards[ids[i]] = struct{}{}
	}
}

// ClearCards clears the "cards" edge to the Card entity.
func (m *DeckMutation) ClearCards() {
	m.clearedcards = true
}

// CardsCleared reports if the "cards" edge to the Card entity was cleared.
func (m *DeckMutation) CardsCleared() bool {
	return m.clearedcards
}

// RemoveCardIDs removes the "cards" edge to the Card entity by IDs.
func (m *DeckMutation) RemoveCardIDs(ids ...int) {
	if m.removedcards == nil {
		m.removedcards = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.cards, ids[i])
		m.removedcards[ids[i]] = struct{}{}
	}
}

// RemovedCards returns the removed IDs of the "cards" edge to the Card entity.
func (m *DeckMutation) RemovedCardsIDs() (ids []int) {
	for id := range m.removedcards {
		ids = append(ids, id)
	}
	return
}

// CardsIDs returns the "cards" edge IDs in the mutation.
func (m *DeckMutation) CardsIDs() (ids []int) {
	for id := range m.cards {
		ids = append(ids, id)
	}
	return
}

// ResetCards resets all changes to the "cards" edge.
func (m *DeckMutation) ResetCards() {
	m.cards = nil
	m.clearedcards = false
	m.removedcards = nil
}

// Where appends a list predicates to the DeckMutation builder.
func (m *DeckMutation) Where(ps ...predicate.Deck) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the DeckMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *DeckMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Deck, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *DeckMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *DeckMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Deck).
func (m *DeckMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DeckMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.title != nil {
		fields = append(fields, deck.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, deck.FieldDescription)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DeckMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case deck.FieldTitle:
		return m.Title()
	case deck.FieldDescription:
		return m.Description()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DeckMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case deck.FieldTitle:
		return m.OldTitle(ctx)
	case deck.FieldDescription:
		return m.OldDescription(ctx)
	}
	return nil, fmt.Errorf("unknown Deck field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DeckMutation) SetField(name string, value ent.Value) error {
	switch name {
	case deck.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case deck.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	}
	return fmt.Errorf("unknown Deck field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DeckMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DeckMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DeckMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Deck numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DeckMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DeckMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DeckMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Deck nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DeckMutation) ResetField(name string) error {
	switch name {
	case deck.FieldTitle:
		m.ResetTitle()
		return nil
	case deck.FieldDescription:
		m.ResetDescription()
		return nil
	}
	return fmt.Errorf("unknown Deck field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DeckMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cards != nil {
		edges = append(edges, deck.EdgeCards)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DeckMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case deck.EdgeCards:
		ids := make([]ent.Value, 0, len(m.cards))
		for id := range m.cards {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DeckMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcards != nil {
		edges = append(edges, deck.EdgeCards)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DeckMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case deck.EdgeCards:
		ids := make([]ent.Value, 0, len(m.removedcards))
		for id := range m.removedcards {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DeckMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcards {
		edges = append(edges, deck.EdgeCards)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DeckMutation) EdgeCleared(name string) bool {
	switch name {
	case deck.EdgeCards:
		return m.clearedcards
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DeckMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Deck unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DeckMutation) ResetEdge(name string) error {
	switch name {
	case deck.EdgeCards:
		m.ResetCards()
		return nil
	}
	return fmt.Errorf("unknown Deck edge %s", name)
}
