// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/oaktreecode/oakapi/ent/migrate"

	"github.com/oaktreecode/oakapi/ent/apiconfig"
	"github.com/oaktreecode/oakapi/ent/taxonomyvocabulary"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ApiConfig is the client for interacting with the ApiConfig builders.
	ApiConfig *ApiConfigClient
	// TaxonomyVocabulary is the client for interacting with the TaxonomyVocabulary builders.
	TaxonomyVocabulary *TaxonomyVocabularyClient
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
	c.ApiConfig = NewApiConfigClient(c.config)
	c.TaxonomyVocabulary = NewTaxonomyVocabularyClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		ApiConfig:          NewApiConfigClient(cfg),
		TaxonomyVocabulary: NewTaxonomyVocabularyClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:                ctx,
		config:             cfg,
		ApiConfig:          NewApiConfigClient(cfg),
		TaxonomyVocabulary: NewTaxonomyVocabularyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ApiConfig.
//		Query().
//		Count(ctx)
//
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
	c.ApiConfig.Use(hooks...)
	c.TaxonomyVocabulary.Use(hooks...)
}

// ApiConfigClient is a client for the ApiConfig schema.
type ApiConfigClient struct {
	config
}

// NewApiConfigClient returns a client for the ApiConfig from the given config.
func NewApiConfigClient(c config) *ApiConfigClient {
	return &ApiConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apiconfig.Hooks(f(g(h())))`.
func (c *ApiConfigClient) Use(hooks ...Hook) {
	c.hooks.ApiConfig = append(c.hooks.ApiConfig, hooks...)
}

// Create returns a builder for creating a ApiConfig entity.
func (c *ApiConfigClient) Create() *ApiConfigCreate {
	mutation := newApiConfigMutation(c.config, OpCreate)
	return &ApiConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ApiConfig entities.
func (c *ApiConfigClient) CreateBulk(builders ...*ApiConfigCreate) *ApiConfigCreateBulk {
	return &ApiConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ApiConfig.
func (c *ApiConfigClient) Update() *ApiConfigUpdate {
	mutation := newApiConfigMutation(c.config, OpUpdate)
	return &ApiConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApiConfigClient) UpdateOne(ac *ApiConfig) *ApiConfigUpdateOne {
	mutation := newApiConfigMutation(c.config, OpUpdateOne, withApiConfig(ac))
	return &ApiConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApiConfigClient) UpdateOneID(id int) *ApiConfigUpdateOne {
	mutation := newApiConfigMutation(c.config, OpUpdateOne, withApiConfigID(id))
	return &ApiConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ApiConfig.
func (c *ApiConfigClient) Delete() *ApiConfigDelete {
	mutation := newApiConfigMutation(c.config, OpDelete)
	return &ApiConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ApiConfigClient) DeleteOne(ac *ApiConfig) *ApiConfigDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ApiConfigClient) DeleteOneID(id int) *ApiConfigDeleteOne {
	builder := c.Delete().Where(apiconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApiConfigDeleteOne{builder}
}

// Query returns a query builder for ApiConfig.
func (c *ApiConfigClient) Query() *ApiConfigQuery {
	return &ApiConfigQuery{
		config: c.config,
	}
}

// Get returns a ApiConfig entity by its id.
func (c *ApiConfigClient) Get(ctx context.Context, id int) (*ApiConfig, error) {
	return c.Query().Where(apiconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApiConfigClient) GetX(ctx context.Context, id int) *ApiConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ApiConfigClient) Hooks() []Hook {
	return c.hooks.ApiConfig
}

// TaxonomyVocabularyClient is a client for the TaxonomyVocabulary schema.
type TaxonomyVocabularyClient struct {
	config
}

// NewTaxonomyVocabularyClient returns a client for the TaxonomyVocabulary from the given config.
func NewTaxonomyVocabularyClient(c config) *TaxonomyVocabularyClient {
	return &TaxonomyVocabularyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `taxonomyvocabulary.Hooks(f(g(h())))`.
func (c *TaxonomyVocabularyClient) Use(hooks ...Hook) {
	c.hooks.TaxonomyVocabulary = append(c.hooks.TaxonomyVocabulary, hooks...)
}

// Create returns a builder for creating a TaxonomyVocabulary entity.
func (c *TaxonomyVocabularyClient) Create() *TaxonomyVocabularyCreate {
	mutation := newTaxonomyVocabularyMutation(c.config, OpCreate)
	return &TaxonomyVocabularyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TaxonomyVocabulary entities.
func (c *TaxonomyVocabularyClient) CreateBulk(builders ...*TaxonomyVocabularyCreate) *TaxonomyVocabularyCreateBulk {
	return &TaxonomyVocabularyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TaxonomyVocabulary.
func (c *TaxonomyVocabularyClient) Update() *TaxonomyVocabularyUpdate {
	mutation := newTaxonomyVocabularyMutation(c.config, OpUpdate)
	return &TaxonomyVocabularyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaxonomyVocabularyClient) UpdateOne(tv *TaxonomyVocabulary) *TaxonomyVocabularyUpdateOne {
	mutation := newTaxonomyVocabularyMutation(c.config, OpUpdateOne, withTaxonomyVocabulary(tv))
	return &TaxonomyVocabularyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaxonomyVocabularyClient) UpdateOneID(id uuid.UUID) *TaxonomyVocabularyUpdateOne {
	mutation := newTaxonomyVocabularyMutation(c.config, OpUpdateOne, withTaxonomyVocabularyID(id))
	return &TaxonomyVocabularyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TaxonomyVocabulary.
func (c *TaxonomyVocabularyClient) Delete() *TaxonomyVocabularyDelete {
	mutation := newTaxonomyVocabularyMutation(c.config, OpDelete)
	return &TaxonomyVocabularyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TaxonomyVocabularyClient) DeleteOne(tv *TaxonomyVocabulary) *TaxonomyVocabularyDeleteOne {
	return c.DeleteOneID(tv.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TaxonomyVocabularyClient) DeleteOneID(id uuid.UUID) *TaxonomyVocabularyDeleteOne {
	builder := c.Delete().Where(taxonomyvocabulary.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaxonomyVocabularyDeleteOne{builder}
}

// Query returns a query builder for TaxonomyVocabulary.
func (c *TaxonomyVocabularyClient) Query() *TaxonomyVocabularyQuery {
	return &TaxonomyVocabularyQuery{
		config: c.config,
	}
}

// Get returns a TaxonomyVocabulary entity by its id.
func (c *TaxonomyVocabularyClient) Get(ctx context.Context, id uuid.UUID) (*TaxonomyVocabulary, error) {
	return c.Query().Where(taxonomyvocabulary.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaxonomyVocabularyClient) GetX(ctx context.Context, id uuid.UUID) *TaxonomyVocabulary {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TaxonomyVocabularyClient) Hooks() []Hook {
	return c.hooks.TaxonomyVocabulary
}