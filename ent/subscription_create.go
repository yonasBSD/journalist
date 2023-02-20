// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/user"
)

// SubscriptionCreate is the builder for creating a Subscription entity.
type SubscriptionCreate struct {
	config
	mutation *SubscriptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (sc *SubscriptionCreate) SetUserID(u uuid.UUID) *SubscriptionCreate {
	sc.mutation.SetUserID(u)
	return sc
}

// SetFeedID sets the "feed_id" field.
func (sc *SubscriptionCreate) SetFeedID(u uuid.UUID) *SubscriptionCreate {
	sc.mutation.SetFeedID(u)
	return sc
}

// SetName sets the "name" field.
func (sc *SubscriptionCreate) SetName(s string) *SubscriptionCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetGroup sets the "group" field.
func (sc *SubscriptionCreate) SetGroup(s string) *SubscriptionCreate {
	sc.mutation.SetGroup(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubscriptionCreate) SetCreatedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableCreatedAt(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubscriptionCreate) SetID(u uuid.UUID) *SubscriptionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableID(u *uuid.UUID) *SubscriptionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SubscriptionCreate) SetUser(u *User) *SubscriptionCreate {
	return sc.SetUserID(u.ID)
}

// SetFeed sets the "feed" edge to the Feed entity.
func (sc *SubscriptionCreate) SetFeed(f *Feed) *SubscriptionCreate {
	return sc.SetFeedID(f.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (sc *SubscriptionCreate) Mutation() *SubscriptionMutation {
	return sc.mutation
}

// Save creates the Subscription in the database.
func (sc *SubscriptionCreate) Save(ctx context.Context) (*Subscription, error) {
	sc.defaults()
	return withHooks[*Subscription, SubscriptionMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriptionCreate) SaveX(ctx context.Context) *Subscription {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubscriptionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubscriptionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubscriptionCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subscription.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := subscription.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubscriptionCreate) check() error {
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Subscription.user_id"`)}
	}
	if _, ok := sc.mutation.FeedID(); !ok {
		return &ValidationError{Name: "feed_id", err: errors.New(`ent: missing required field "Subscription.feed_id"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subscription.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := subscription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subscription.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Group(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required field "Subscription.group"`)}
	}
	if v, ok := sc.mutation.Group(); ok {
		if err := subscription.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "Subscription.group": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Subscription.created_at"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Subscription.user"`)}
	}
	if _, ok := sc.mutation.FeedID(); !ok {
		return &ValidationError{Name: "feed", err: errors.New(`ent: missing required edge "Subscription.feed"`)}
	}
	return nil
}

func (sc *SubscriptionCreate) sqlSave(ctx context.Context) (*Subscription, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubscriptionCreate) createSpec() (*Subscription, *sqlgraph.CreateSpec) {
	var (
		_node = &Subscription{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subscription.Table, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subscription.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Group(); ok {
		_spec.SetField(subscription.FieldGroup, field.TypeString, value)
		_node.Group = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subscription.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   subscription.FeedTable,
			Columns: []string{subscription.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FeedID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscription.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriptionUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (sc *SubscriptionCreate) OnConflict(opts ...sql.ConflictOption) *SubscriptionUpsertOne {
	sc.conflict = opts
	return &SubscriptionUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SubscriptionCreate) OnConflictColumns(columns ...string) *SubscriptionUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SubscriptionUpsertOne{
		create: sc,
	}
}

type (
	// SubscriptionUpsertOne is the builder for "upsert"-ing
	//  one Subscription node.
	SubscriptionUpsertOne struct {
		create *SubscriptionCreate
	}

	// SubscriptionUpsert is the "OnConflict" setter.
	SubscriptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *SubscriptionUpsert) SetUserID(v uuid.UUID) *SubscriptionUpsert {
	u.Set(subscription.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateUserID() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldUserID)
	return u
}

// SetFeedID sets the "feed_id" field.
func (u *SubscriptionUpsert) SetFeedID(v uuid.UUID) *SubscriptionUpsert {
	u.Set(subscription.FieldFeedID, v)
	return u
}

// UpdateFeedID sets the "feed_id" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateFeedID() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldFeedID)
	return u
}

// SetName sets the "name" field.
func (u *SubscriptionUpsert) SetName(v string) *SubscriptionUpsert {
	u.Set(subscription.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateName() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldName)
	return u
}

// SetGroup sets the "group" field.
func (u *SubscriptionUpsert) SetGroup(v string) *SubscriptionUpsert {
	u.Set(subscription.FieldGroup, v)
	return u
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateGroup() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldGroup)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SubscriptionUpsert) SetCreatedAt(v time.Time) *SubscriptionUpsert {
	u.Set(subscription.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateCreatedAt() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscription.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubscriptionUpsertOne) UpdateNewValues() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subscription.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subscription.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubscriptionUpsertOne) Ignore() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriptionUpsertOne) DoNothing() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriptionCreate.OnConflict
// documentation for more info.
func (u *SubscriptionUpsertOne) Update(set func(*SubscriptionUpsert)) *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *SubscriptionUpsertOne) SetUserID(v uuid.UUID) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateUserID() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUserID()
	})
}

// SetFeedID sets the "feed_id" field.
func (u *SubscriptionUpsertOne) SetFeedID(v uuid.UUID) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetFeedID(v)
	})
}

// UpdateFeedID sets the "feed_id" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateFeedID() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateFeedID()
	})
}

// SetName sets the "name" field.
func (u *SubscriptionUpsertOne) SetName(v string) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateName() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateName()
	})
}

// SetGroup sets the "group" field.
func (u *SubscriptionUpsertOne) SetGroup(v string) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateGroup() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateGroup()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SubscriptionUpsertOne) SetCreatedAt(v time.Time) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateCreatedAt() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *SubscriptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubscriptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubscriptionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SubscriptionUpsertOne.ID is not supported by MySQL driver. Use SubscriptionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubscriptionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubscriptionCreateBulk is the builder for creating many Subscription entities in bulk.
type SubscriptionCreateBulk struct {
	config
	builders []*SubscriptionCreate
	conflict []sql.ConflictOption
}

// Save creates the Subscription entities in the database.
func (scb *SubscriptionCreateBulk) Save(ctx context.Context) ([]*Subscription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subscription, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) SaveX(ctx context.Context) []*Subscription {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscription.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriptionUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (scb *SubscriptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubscriptionUpsertBulk {
	scb.conflict = opts
	return &SubscriptionUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SubscriptionCreateBulk) OnConflictColumns(columns ...string) *SubscriptionUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SubscriptionUpsertBulk{
		create: scb,
	}
}

// SubscriptionUpsertBulk is the builder for "upsert"-ing
// a bulk of Subscription nodes.
type SubscriptionUpsertBulk struct {
	create *SubscriptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscription.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubscriptionUpsertBulk) UpdateNewValues() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subscription.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubscriptionUpsertBulk) Ignore() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriptionUpsertBulk) DoNothing() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriptionCreateBulk.OnConflict
// documentation for more info.
func (u *SubscriptionUpsertBulk) Update(set func(*SubscriptionUpsert)) *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *SubscriptionUpsertBulk) SetUserID(v uuid.UUID) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateUserID() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUserID()
	})
}

// SetFeedID sets the "feed_id" field.
func (u *SubscriptionUpsertBulk) SetFeedID(v uuid.UUID) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetFeedID(v)
	})
}

// UpdateFeedID sets the "feed_id" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateFeedID() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateFeedID()
	})
}

// SetName sets the "name" field.
func (u *SubscriptionUpsertBulk) SetName(v string) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateName() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateName()
	})
}

// SetGroup sets the "group" field.
func (u *SubscriptionUpsertBulk) SetGroup(v string) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetGroup(v)
	})
}

// UpdateGroup sets the "group" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateGroup() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateGroup()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SubscriptionUpsertBulk) SetCreatedAt(v time.Time) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateCreatedAt() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *SubscriptionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SubscriptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubscriptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
