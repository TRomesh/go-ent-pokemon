package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pokemon holds the schema definition for the Pokemon entity.
type Pokemon struct {
	ent.Schema
}

// Fields of the Pokemon.
func (Pokemon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`json:"oid,omitempty"`),
		field.Text("name").
			NotEmpty(),
		field.Text("description").
			NotEmpty(),
		field.Float("weight"),
		field.Float("height"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Pokemon.
func (Pokemon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contender", Battle.Type),
		edge.To("opponent", Battle.Type),
	}
}
