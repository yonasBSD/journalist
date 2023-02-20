// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/migrate"

	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/read"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/token"
	"github.com/mrusme/journalist/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Feed is the client for interacting with the Feed builders.
	Feed *FeedClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// Read is the client for interacting with the Read builders.
	Read *ReadClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Feed = NewFeedClient(c.config)
	c.Item = NewItemClient(c.config)
	c.Read = NewReadClient(c.config)
	c.Subscription = NewSubscriptionClient(c.config)
	c.Token = NewTokenClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Feed:         NewFeedClient(cfg),
		Item:         NewItemClient(cfg),
		Read:         NewReadClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		Token:        NewTokenClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Feed:         NewFeedClient(cfg),
		Item:         NewItemClient(cfg),
		Read:         NewReadClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		Token:        NewTokenClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Feed.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Feed.Use(hooks...)
	c.Item.Use(hooks...)
	c.Read.Use(hooks...)
	c.Subscription.Use(hooks...)
	c.Token.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Feed.Intercept(interceptors...)
	c.Item.Intercept(interceptors...)
	c.Read.Intercept(interceptors...)
	c.Subscription.Intercept(interceptors...)
	c.Token.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *FeedMutation:
		return c.Feed.mutate(ctx, m)
	case *ItemMutation:
		return c.Item.mutate(ctx, m)
	case *ReadMutation:
		return c.Read.mutate(ctx, m)
	case *SubscriptionMutation:
		return c.Subscription.mutate(ctx, m)
	case *TokenMutation:
		return c.Token.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// FeedClient is a client for the Feed schema.
type FeedClient struct {
	config
}

// NewFeedClient returns a client for the Feed from the given config.
func NewFeedClient(c config) *FeedClient {
	return &FeedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `feed.Hooks(f(g(h())))`.
func (c *FeedClient) Use(hooks ...Hook) {
	c.hooks.Feed = append(c.hooks.Feed, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `feed.Intercept(f(g(h())))`.
func (c *FeedClient) Intercept(interceptors ...Interceptor) {
	c.inters.Feed = append(c.inters.Feed, interceptors...)
}

// Create returns a builder for creating a Feed entity.
func (c *FeedClient) Create() *FeedCreate {
	mutation := newFeedMutation(c.config, OpCreate)
	return &FeedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Feed entities.
func (c *FeedClient) CreateBulk(builders ...*FeedCreate) *FeedCreateBulk {
	return &FeedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Feed.
func (c *FeedClient) Update() *FeedUpdate {
	mutation := newFeedMutation(c.config, OpUpdate)
	return &FeedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeedClient) UpdateOne(f *Feed) *FeedUpdateOne {
	mutation := newFeedMutation(c.config, OpUpdateOne, withFeed(f))
	return &FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeedClient) UpdateOneID(id uuid.UUID) *FeedUpdateOne {
	mutation := newFeedMutation(c.config, OpUpdateOne, withFeedID(id))
	return &FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Feed.
func (c *FeedClient) Delete() *FeedDelete {
	mutation := newFeedMutation(c.config, OpDelete)
	return &FeedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeedClient) DeleteOne(f *Feed) *FeedDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeedClient) DeleteOneID(id uuid.UUID) *FeedDeleteOne {
	builder := c.Delete().Where(feed.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeedDeleteOne{builder}
}

// Query returns a query builder for Feed.
func (c *FeedClient) Query() *FeedQuery {
	return &FeedQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFeed},
		inters: c.Interceptors(),
	}
}

// Get returns a Feed entity by its id.
func (c *FeedClient) Get(ctx context.Context, id uuid.UUID) (*Feed, error) {
	return c.Query().Where(feed.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeedClient) GetX(ctx context.Context, id uuid.UUID) *Feed {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItems queries the items edge of a Feed.
func (c *FeedClient) QueryItems(f *Feed) *ItemQuery {
	query := (&ItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, feed.ItemsTable, feed.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubscribedUsers queries the subscribed_users edge of a Feed.
func (c *FeedClient) QuerySubscribedUsers(f *Feed) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, feed.SubscribedUsersTable, feed.SubscribedUsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubscriptions queries the subscriptions edge of a Feed.
func (c *FeedClient) QuerySubscriptions(f *Feed) *SubscriptionQuery {
	query := (&SubscriptionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(feed.Table, feed.FieldID, id),
			sqlgraph.To(subscription.Table, subscription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, feed.SubscriptionsTable, feed.SubscriptionsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FeedClient) Hooks() []Hook {
	return c.hooks.Feed
}

// Interceptors returns the client interceptors.
func (c *FeedClient) Interceptors() []Interceptor {
	return c.inters.Feed
}

func (c *FeedClient) mutate(ctx context.Context, m *FeedMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeedCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeedUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeedDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Feed mutation op: %q", m.Op())
	}
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `item.Intercept(f(g(h())))`.
func (c *ItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.Item = append(c.inters.Item, interceptors...)
}

// Create returns a builder for creating a Item entity.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id uuid.UUID) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ItemClient) DeleteOneID(id uuid.UUID) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeItem},
		inters: c.Interceptors(),
	}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id uuid.UUID) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id uuid.UUID) *Item {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFeed queries the feed edge of a Item.
func (c *ItemClient) QueryFeed(i *Item) *FeedQuery {
	query := (&FeedClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(feed.Table, feed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.FeedTable, item.FeedColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReadByUsers queries the read_by_users edge of a Item.
func (c *ItemClient) QueryReadByUsers(i *Item) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.ReadByUsersTable, item.ReadByUsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReads queries the reads edge of a Item.
func (c *ItemClient) QueryReads(i *Item) *ReadQuery {
	query := (&ReadClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(read.Table, read.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, item.ReadsTable, item.ReadsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// Interceptors returns the client interceptors.
func (c *ItemClient) Interceptors() []Interceptor {
	return c.inters.Item
}

func (c *ItemClient) mutate(ctx context.Context, m *ItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Item mutation op: %q", m.Op())
	}
}

// ReadClient is a client for the Read schema.
type ReadClient struct {
	config
}

// NewReadClient returns a client for the Read from the given config.
func NewReadClient(c config) *ReadClient {
	return &ReadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `read.Hooks(f(g(h())))`.
func (c *ReadClient) Use(hooks ...Hook) {
	c.hooks.Read = append(c.hooks.Read, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `read.Intercept(f(g(h())))`.
func (c *ReadClient) Intercept(interceptors ...Interceptor) {
	c.inters.Read = append(c.inters.Read, interceptors...)
}

// Create returns a builder for creating a Read entity.
func (c *ReadClient) Create() *ReadCreate {
	mutation := newReadMutation(c.config, OpCreate)
	return &ReadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Read entities.
func (c *ReadClient) CreateBulk(builders ...*ReadCreate) *ReadCreateBulk {
	return &ReadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Read.
func (c *ReadClient) Update() *ReadUpdate {
	mutation := newReadMutation(c.config, OpUpdate)
	return &ReadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReadClient) UpdateOne(r *Read) *ReadUpdateOne {
	mutation := newReadMutation(c.config, OpUpdateOne, withRead(r))
	return &ReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReadClient) UpdateOneID(id uuid.UUID) *ReadUpdateOne {
	mutation := newReadMutation(c.config, OpUpdateOne, withReadID(id))
	return &ReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Read.
func (c *ReadClient) Delete() *ReadDelete {
	mutation := newReadMutation(c.config, OpDelete)
	return &ReadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReadClient) DeleteOne(r *Read) *ReadDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReadClient) DeleteOneID(id uuid.UUID) *ReadDeleteOne {
	builder := c.Delete().Where(read.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReadDeleteOne{builder}
}

// Query returns a query builder for Read.
func (c *ReadClient) Query() *ReadQuery {
	return &ReadQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRead},
		inters: c.Interceptors(),
	}
}

// Get returns a Read entity by its id.
func (c *ReadClient) Get(ctx context.Context, id uuid.UUID) (*Read, error) {
	return c.Query().Where(read.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReadClient) GetX(ctx context.Context, id uuid.UUID) *Read {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Read.
func (c *ReadClient) QueryUser(r *Read) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(read.Table, read.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, read.UserTable, read.UserColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItem queries the item edge of a Read.
func (c *ReadClient) QueryItem(r *Read) *ItemQuery {
	query := (&ItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(read.Table, read.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, read.ItemTable, read.ItemColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReadClient) Hooks() []Hook {
	return c.hooks.Read
}

// Interceptors returns the client interceptors.
func (c *ReadClient) Interceptors() []Interceptor {
	return c.inters.Read
}

func (c *ReadClient) mutate(ctx context.Context, m *ReadMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ReadCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ReadUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ReadDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Read mutation op: %q", m.Op())
	}
}

// SubscriptionClient is a client for the Subscription schema.
type SubscriptionClient struct {
	config
}

// NewSubscriptionClient returns a client for the Subscription from the given config.
func NewSubscriptionClient(c config) *SubscriptionClient {
	return &SubscriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subscription.Hooks(f(g(h())))`.
func (c *SubscriptionClient) Use(hooks ...Hook) {
	c.hooks.Subscription = append(c.hooks.Subscription, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `subscription.Intercept(f(g(h())))`.
func (c *SubscriptionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Subscription = append(c.inters.Subscription, interceptors...)
}

// Create returns a builder for creating a Subscription entity.
func (c *SubscriptionClient) Create() *SubscriptionCreate {
	mutation := newSubscriptionMutation(c.config, OpCreate)
	return &SubscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Subscription entities.
func (c *SubscriptionClient) CreateBulk(builders ...*SubscriptionCreate) *SubscriptionCreateBulk {
	return &SubscriptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Subscription.
func (c *SubscriptionClient) Update() *SubscriptionUpdate {
	mutation := newSubscriptionMutation(c.config, OpUpdate)
	return &SubscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubscriptionClient) UpdateOne(s *Subscription) *SubscriptionUpdateOne {
	mutation := newSubscriptionMutation(c.config, OpUpdateOne, withSubscription(s))
	return &SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubscriptionClient) UpdateOneID(id uuid.UUID) *SubscriptionUpdateOne {
	mutation := newSubscriptionMutation(c.config, OpUpdateOne, withSubscriptionID(id))
	return &SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Subscription.
func (c *SubscriptionClient) Delete() *SubscriptionDelete {
	mutation := newSubscriptionMutation(c.config, OpDelete)
	return &SubscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubscriptionClient) DeleteOne(s *Subscription) *SubscriptionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubscriptionClient) DeleteOneID(id uuid.UUID) *SubscriptionDeleteOne {
	builder := c.Delete().Where(subscription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubscriptionDeleteOne{builder}
}

// Query returns a query builder for Subscription.
func (c *SubscriptionClient) Query() *SubscriptionQuery {
	return &SubscriptionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubscription},
		inters: c.Interceptors(),
	}
}

// Get returns a Subscription entity by its id.
func (c *SubscriptionClient) Get(ctx context.Context, id uuid.UUID) (*Subscription, error) {
	return c.Query().Where(subscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriptionClient) GetX(ctx context.Context, id uuid.UUID) *Subscription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Subscription.
func (c *SubscriptionClient) QueryUser(s *Subscription) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subscription.Table, subscription.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, subscription.UserTable, subscription.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFeed queries the feed edge of a Subscription.
func (c *SubscriptionClient) QueryFeed(s *Subscription) *FeedQuery {
	query := (&FeedClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subscription.Table, subscription.FieldID, id),
			sqlgraph.To(feed.Table, feed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, subscription.FeedTable, subscription.FeedColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubscriptionClient) Hooks() []Hook {
	return c.hooks.Subscription
}

// Interceptors returns the client interceptors.
func (c *SubscriptionClient) Interceptors() []Interceptor {
	return c.inters.Subscription
}

func (c *SubscriptionClient) mutate(ctx context.Context, m *SubscriptionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Subscription mutation op: %q", m.Op())
	}
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `token.Intercept(f(g(h())))`.
func (c *TokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.Token = append(c.inters.Token, interceptors...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id uuid.UUID) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id uuid.UUID) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeToken},
		inters: c.Interceptors(),
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id uuid.UUID) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id uuid.UUID) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Token.
func (c *TokenClient) QueryOwner(t *Token) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, token.OwnerTable, token.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// Interceptors returns the client interceptors.
func (c *TokenClient) Interceptors() []Interceptor {
	return c.inters.Token
}

func (c *TokenClient) mutate(ctx context.Context, m *TokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Token mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTokens queries the tokens edge of a User.
func (c *UserClient) QueryTokens(u *User) *TokenQuery {
	query := (&TokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TokensTable, user.TokensColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubscribedFeeds queries the subscribed_feeds edge of a User.
func (c *UserClient) QuerySubscribedFeeds(u *User) *FeedQuery {
	query := (&FeedClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(feed.Table, feed.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.SubscribedFeedsTable, user.SubscribedFeedsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReadItems queries the read_items edge of a User.
func (c *UserClient) QueryReadItems(u *User) *ItemQuery {
	query := (&ItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ReadItemsTable, user.ReadItemsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubscriptions queries the subscriptions edge of a User.
func (c *UserClient) QuerySubscriptions(u *User) *SubscriptionQuery {
	query := (&SubscriptionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(subscription.Table, subscription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.SubscriptionsTable, user.SubscriptionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReads queries the reads edge of a User.
func (c *UserClient) QueryReads(u *User) *ReadQuery {
	query := (&ReadClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(read.Table, read.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.ReadsTable, user.ReadsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}
