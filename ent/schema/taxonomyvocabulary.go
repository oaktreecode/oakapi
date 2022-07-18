package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TaxonomyVocabulary holds the schema definition for the TaxonomyVocabulary entity.
type TaxonomyVocabulary struct {
	ent.Schema
}

// Fields of the TaxonomyVocabulary.
func (TaxonomyVocabulary) Fields() []ent.Field {
	return []ent.Field{
		field.String("vid").
			Unique().
			NotEmpty().
			Comment("机器名"),
		field.String("label").
			NotEmpty().
			Comment("名称"),
		field.Int("weight").
			Default(0).
			Comment("排序"),
		field.String("description").Comment("描述"),
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the TaxonomyVocabulary.
func (TaxonomyVocabulary) Edges() []ent.Edge {
	return nil
}
