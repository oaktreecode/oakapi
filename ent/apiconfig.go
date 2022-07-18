// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/oaktreecode/oakapi/ent/apiconfig"
)

// ApiConfig is the model entity for the ApiConfig schema.
type ApiConfig struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApiConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case apiconfig.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApiConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApiConfig fields.
func (ac *ApiConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apiconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this ApiConfig.
// Note that you need to call ApiConfig.Unwrap() before calling this method if this ApiConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *ApiConfig) Update() *ApiConfigUpdateOne {
	return (&ApiConfigClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the ApiConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *ApiConfig) Unwrap() *ApiConfig {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApiConfig is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *ApiConfig) String() string {
	var builder strings.Builder
	builder.WriteString("ApiConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", ac.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ApiConfigs is a parsable slice of ApiConfig.
type ApiConfigs []*ApiConfig

func (ac ApiConfigs) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}