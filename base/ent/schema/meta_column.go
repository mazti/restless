package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Meta holds the schema definition for the Meta entity.
type MetaColumn struct {
	ent.Schema
}

// Fields of the Meta.
func (MetaColumn) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
		field.String("default"),
		field.Text("type_option"),
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
func (MetaColumn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("table", MetaTable.Type).
			Ref("columns").
			Unique(),
	}
}
