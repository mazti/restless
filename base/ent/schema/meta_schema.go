package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Meta holds the schema definition for the Meta entity.
type MetaSchema struct {
	ent.Schema
}

// Fields of the Meta.
func (MetaSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("base"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Meta.
func (MetaSchema) Edges() []ent.Edge {
	return nil
}
