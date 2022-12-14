package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("creator_id"),
		field.Time("create_at").Default(time.Now),
		field.Time("last_update").Default(time.Now).UpdateDefault(time.Now),
		field.Int("content_flag").Optional(),
		field.Int("sent_flag").Optional(),
		field.Int("mentioned_count").Default(1),
		field.Time("take_effect_time").Optional(),
		field.Bytes("sending_content").Optional(),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return nil
}

func (History) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_at"),
		index.Fields("last_update"),
		index.Fields("sent_flag"),
		index.Fields("id"),
	}
}
