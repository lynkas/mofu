package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("last_landmark").Optional(),
		field.Time("last_update").Optional(),
		field.String("name").Optional(),
		field.String("username").Optional(),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}
