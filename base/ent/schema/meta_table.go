package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Meta holds the schema definition for the Meta entity.
type MetaTable struct {
	ent.Schema
}

// Fields of the Meta.
func (MetaTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
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
func (MetaTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schema", MetaSchema.Type).
			Ref("tables").
			Unique(),
		edge.To("columns", MetaColumn.Type),
	}
}
