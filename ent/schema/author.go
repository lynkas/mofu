package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique(),
		field.String("user_name"),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return nil
}
