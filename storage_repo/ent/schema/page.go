package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.String("path").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
		field.String("url").Default(""),
		field.String("description").Default(""),
		field.String("author_name").Default(""),
		field.String("author_url").Default(""),
		field.String("image_url").Default(""),
		field.Int("views").Default(0),
	}
}

func (Page) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path").Unique(),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
