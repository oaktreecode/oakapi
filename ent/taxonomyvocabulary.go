// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/oaktreecode/oakapi/ent/taxonomyvocabulary"
)

// TaxonomyVocabulary is the model entity for the TaxonomyVocabulary schema.
type TaxonomyVocabulary struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 机器名
	Vid string `json:"vid,omitempty"`
	// 名称
	Label string `json:"label,omitempty"`
	// 排序
	Weight int `json:"weight,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaxonomyVocabulary) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case taxonomyvocabulary.FieldWeight:
			values[i] = new(sql.NullInt64)
		case taxonomyvocabulary.FieldVid, taxonomyvocabulary.FieldLabel, taxonomyvocabulary.FieldDescription:
			values[i] = new(sql.NullString)
		case taxonomyvocabulary.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaxonomyVocabulary", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaxonomyVocabulary fields.
func (tv *TaxonomyVocabulary) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taxonomyvocabulary.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tv.ID = *value
			}
		case taxonomyvocabulary.FieldVid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vid", values[i])
			} else if value.Valid {
				tv.Vid = value.String
			}
		case taxonomyvocabulary.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				tv.Label = value.String
			}
		case taxonomyvocabulary.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				tv.Weight = int(value.Int64)
			}
		case taxonomyvocabulary.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				tv.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TaxonomyVocabulary.
// Note that you need to call TaxonomyVocabulary.Unwrap() before calling this method if this TaxonomyVocabulary
// was returned from a transaction, and the transaction was committed or rolled back.
func (tv *TaxonomyVocabulary) Update() *TaxonomyVocabularyUpdateOne {
	return (&TaxonomyVocabularyClient{config: tv.config}).UpdateOne(tv)
}

// Unwrap unwraps the TaxonomyVocabulary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tv *TaxonomyVocabulary) Unwrap() *TaxonomyVocabulary {
	_tx, ok := tv.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaxonomyVocabulary is not a transactional entity")
	}
	tv.config.driver = _tx.drv
	return tv
}

// String implements the fmt.Stringer.
func (tv *TaxonomyVocabulary) String() string {
	var builder strings.Builder
	builder.WriteString("TaxonomyVocabulary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tv.ID))
	builder.WriteString("vid=")
	builder.WriteString(tv.Vid)
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(tv.Label)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", tv.Weight))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(tv.Description)
	builder.WriteByte(')')
	return builder.String()
}

// TaxonomyVocabularies is a parsable slice of TaxonomyVocabulary.
type TaxonomyVocabularies []*TaxonomyVocabulary

func (tv TaxonomyVocabularies) config(cfg config) {
	for _i := range tv {
		tv[_i].config = cfg
	}
}