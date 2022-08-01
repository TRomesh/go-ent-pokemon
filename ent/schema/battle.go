package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Battle holds the schema definition for the Battle entity.
type Battle struct {
	ent.Schema
}

// Fields of the Battle.
func (Battle) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`json:"oid,omitempty"`),
		field.Text("result"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Battle.
func (Battle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contender", Pokemon.Type).
			Ref("contender").
			Unique(),
		edge.From("oponent", Pokemon.Type).
			Ref("opponent").
			Unique(),
	}
}
