package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PageView holds the schema definition for the PageView entity.
type PageView struct {
	ent.Schema
}

// Fields of the PageView.
func (PageView) Fields() []ent.Field {
	return []ent.Field{
		field.Int("page_id"),
		field.String("path").NotEmpty(),
		field.Int("year").Min(2000).Max(2100),
		field.Int("month").Min(1).Max(12),
		field.Int("day").Min(1).Max(31),
		field.Int("hour").Min(0).Max(24),
		field.Int("views").Default(0).Min(0),
	}
}

func (PageView) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("page_id"),
		index.Fields("path"),
	}
}

// Edges of the PageView.
func (PageView) Edges() []ent.Edge {
	return nil
}
