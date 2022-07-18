package schema

import (
	"entgo.io/ent"
)

// ApiConfig holds the schema definition for the ApiConfig entity.
type ApiConfig struct {
	ent.Schema
}

// Fields of the ApiConfig.
func (ApiConfig) Fields() []ent.Field {
	return nil
}

// Edges of the ApiConfig.
func (ApiConfig) Edges() []ent.Edge {
	return nil
}
