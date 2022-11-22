package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_name").NotEmpty(),
		field.String("author_name").Default(""),
		field.String("author_url").Default(""),
		field.String("access_token").NotEmpty(),
		field.String("auth_url").Default(""),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("access_token").Unique(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
