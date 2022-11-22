package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").NotEmpty(),
		field.String("title").NotEmpty(),
		field.JSON("content", map[string]any{}),
		field.String("url").Default(""),
		field.String("description").Default(""),
		field.String("author_name").Default(""),
		field.String("author_url").Default(""),
		field.String("image_url").Default(""),
		field.Int("views").Default(0),
		field.Bool("can_edit").Default(false),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
