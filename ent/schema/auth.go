package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.String("token"),
		field.String("user"),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return nil
}
