// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"mofu/ent/migrate"

	"mofu/ent/auth"
	"mofu/ent/author"
	"mofu/ent/history"
	"mofu/ent/setting"
	"mofu/ent/subscription"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Auth is the client for interacting with the Auth builders.
	Auth *AuthClient
	// Author is the client for interacting with the Author builders.
	Author *AuthorClient
	// History is the client for interacting with the History builders.
	History *HistoryClient
	// Setting is the client for interacting with the Setting builders.
	Setting *SettingClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Auth = NewAuthClient(c.config)
	c.Author = NewAuthorClient(c.config)
	c.History = NewHistoryClient(c.config)
	c.Setting = NewSettingClient(c.config)
	c.Subscription = NewSubscriptionClient(c.config)
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
		Auth:         NewAuthClient(cfg),
		Author:       NewAuthorClient(cfg),
		History:      NewHistoryClient(cfg),
		Setting:      NewSettingClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
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
		Auth:         NewAuthClient(cfg),
		Author:       NewAuthorClient(cfg),
		History:      NewHistoryClient(cfg),
		Setting:      NewSettingClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Auth.
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
	c.Auth.Use(hooks...)
	c.Author.Use(hooks...)
	c.History.Use(hooks...)
	c.Setting.Use(hooks...)
	c.Subscription.Use(hooks...)
}

// AuthClient is a client for the Auth schema.
type AuthClient struct {
	config
}

// NewAuthClient returns a client for the Auth from the given config.
func NewAuthClient(c config) *AuthClient {
	return &AuthClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `auth.Hooks(f(g(h())))`.
func (c *AuthClient) Use(hooks ...Hook) {
	c.hooks.Auth = append(c.hooks.Auth, hooks...)
}

// Create returns a builder for creating a Auth entity.
func (c *AuthClient) Create() *AuthCreate {
	mutation := newAuthMutation(c.config, OpCreate)
	return &AuthCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Auth entities.
func (c *AuthClient) CreateBulk(builders ...*AuthCreate) *AuthCreateBulk {
	return &AuthCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Auth.
func (c *AuthClient) Update() *AuthUpdate {
	mutation := newAuthMutation(c.config, OpUpdate)
	return &AuthUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthClient) UpdateOne(a *Auth) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuth(a))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthClient) UpdateOneID(id int) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuthID(id))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Auth.
func (c *AuthClient) Delete() *AuthDelete {
	mutation := newAuthMutation(c.config, OpDelete)
	return &AuthDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthClient) DeleteOne(a *Auth) *AuthDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AuthClient) DeleteOneID(id int) *AuthDeleteOne {
	builder := c.Delete().Where(auth.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthDeleteOne{builder}
}

// Query returns a query builder for Auth.
func (c *AuthClient) Query() *AuthQuery {
	return &AuthQuery{
		config: c.config,
	}
}

// Get returns a Auth entity by its id.
func (c *AuthClient) Get(ctx context.Context, id int) (*Auth, error) {
	return c.Query().Where(auth.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthClient) GetX(ctx context.Context, id int) *Auth {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthClient) Hooks() []Hook {
	return c.hooks.Auth
}

// AuthorClient is a client for the Author schema.
type AuthorClient struct {
	config
}

// NewAuthorClient returns a client for the Author from the given config.
func NewAuthorClient(c config) *AuthorClient {
	return &AuthorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `author.Hooks(f(g(h())))`.
func (c *AuthorClient) Use(hooks ...Hook) {
	c.hooks.Author = append(c.hooks.Author, hooks...)
}

// Create returns a builder for creating a Author entity.
func (c *AuthorClient) Create() *AuthorCreate {
	mutation := newAuthorMutation(c.config, OpCreate)
	return &AuthorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Author entities.
func (c *AuthorClient) CreateBulk(builders ...*AuthorCreate) *AuthorCreateBulk {
	return &AuthorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Author.
func (c *AuthorClient) Update() *AuthorUpdate {
	mutation := newAuthorMutation(c.config, OpUpdate)
	return &AuthorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthorClient) UpdateOne(a *Author) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthor(a))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthorClient) UpdateOneID(id int) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthorID(id))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Author.
func (c *AuthorClient) Delete() *AuthorDelete {
	mutation := newAuthorMutation(c.config, OpDelete)
	return &AuthorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthorClient) DeleteOne(a *Author) *AuthorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AuthorClient) DeleteOneID(id int) *AuthorDeleteOne {
	builder := c.Delete().Where(author.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthorDeleteOne{builder}
}

// Query returns a query builder for Author.
func (c *AuthorClient) Query() *AuthorQuery {
	return &AuthorQuery{
		config: c.config,
	}
}

// Get returns a Author entity by its id.
func (c *AuthorClient) Get(ctx context.Context, id int) (*Author, error) {
	return c.Query().Where(author.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthorClient) GetX(ctx context.Context, id int) *Author {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthorClient) Hooks() []Hook {
	return c.hooks.Author
}

// HistoryClient is a client for the History schema.
type HistoryClient struct {
	config
}

// NewHistoryClient returns a client for the History from the given config.
func NewHistoryClient(c config) *HistoryClient {
	return &HistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `history.Hooks(f(g(h())))`.
func (c *HistoryClient) Use(hooks ...Hook) {
	c.hooks.History = append(c.hooks.History, hooks...)
}

// Create returns a builder for creating a History entity.
func (c *HistoryClient) Create() *HistoryCreate {
	mutation := newHistoryMutation(c.config, OpCreate)
	return &HistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of History entities.
func (c *HistoryClient) CreateBulk(builders ...*HistoryCreate) *HistoryCreateBulk {
	return &HistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for History.
func (c *HistoryClient) Update() *HistoryUpdate {
	mutation := newHistoryMutation(c.config, OpUpdate)
	return &HistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HistoryClient) UpdateOne(h *History) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistory(h))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HistoryClient) UpdateOneID(id string) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistoryID(id))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for History.
func (c *HistoryClient) Delete() *HistoryDelete {
	mutation := newHistoryMutation(c.config, OpDelete)
	return &HistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HistoryClient) DeleteOne(h *History) *HistoryDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HistoryClient) DeleteOneID(id string) *HistoryDeleteOne {
	builder := c.Delete().Where(history.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HistoryDeleteOne{builder}
}

// Query returns a query builder for History.
func (c *HistoryClient) Query() *HistoryQuery {
	return &HistoryQuery{
		config: c.config,
	}
}

// Get returns a History entity by its id.
func (c *HistoryClient) Get(ctx context.Context, id string) (*History, error) {
	return c.Query().Where(history.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HistoryClient) GetX(ctx context.Context, id string) *History {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *HistoryClient) Hooks() []Hook {
	return c.hooks.History
}

// SettingClient is a client for the Setting schema.
type SettingClient struct {
	config
}

// NewSettingClient returns a client for the Setting from the given config.
func NewSettingClient(c config) *SettingClient {
	return &SettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `setting.Hooks(f(g(h())))`.
func (c *SettingClient) Use(hooks ...Hook) {
	c.hooks.Setting = append(c.hooks.Setting, hooks...)
}

// Create returns a builder for creating a Setting entity.
func (c *SettingClient) Create() *SettingCreate {
	mutation := newSettingMutation(c.config, OpCreate)
	return &SettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Setting entities.
func (c *SettingClient) CreateBulk(builders ...*SettingCreate) *SettingCreateBulk {
	return &SettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Setting.
func (c *SettingClient) Update() *SettingUpdate {
	mutation := newSettingMutation(c.config, OpUpdate)
	return &SettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SettingClient) UpdateOne(s *Setting) *SettingUpdateOne {
	mutation := newSettingMutation(c.config, OpUpdateOne, withSetting(s))
	return &SettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SettingClient) UpdateOneID(id int) *SettingUpdateOne {
	mutation := newSettingMutation(c.config, OpUpdateOne, withSettingID(id))
	return &SettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Setting.
func (c *SettingClient) Delete() *SettingDelete {
	mutation := newSettingMutation(c.config, OpDelete)
	return &SettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SettingClient) DeleteOne(s *Setting) *SettingDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SettingClient) DeleteOneID(id int) *SettingDeleteOne {
	builder := c.Delete().Where(setting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SettingDeleteOne{builder}
}

// Query returns a query builder for Setting.
func (c *SettingClient) Query() *SettingQuery {
	return &SettingQuery{
		config: c.config,
	}
}

// Get returns a Setting entity by its id.
func (c *SettingClient) Get(ctx context.Context, id int) (*Setting, error) {
	return c.Query().Where(setting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SettingClient) GetX(ctx context.Context, id int) *Setting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SettingClient) Hooks() []Hook {
	return c.hooks.Setting
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
func (c *SubscriptionClient) UpdateOneID(id string) *SubscriptionUpdateOne {
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

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SubscriptionClient) DeleteOneID(id string) *SubscriptionDeleteOne {
	builder := c.Delete().Where(subscription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubscriptionDeleteOne{builder}
}

// Query returns a query builder for Subscription.
func (c *SubscriptionClient) Query() *SubscriptionQuery {
	return &SubscriptionQuery{
		config: c.config,
	}
}

// Get returns a Subscription entity by its id.
func (c *SubscriptionClient) Get(ctx context.Context, id string) (*Subscription, error) {
	return c.Query().Where(subscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriptionClient) GetX(ctx context.Context, id string) *Subscription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SubscriptionClient) Hooks() []Hook {
	return c.hooks.Subscription
}
