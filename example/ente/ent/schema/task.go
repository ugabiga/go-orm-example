package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("note").
			Sensitive(),
		field.Enum("status").
			Values(
				"todo",
				"in_progress",
				"done",
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tasks").
			Unique(),
		edge.To("children", Task.Type).
			From("parent").
			Unique(),
		edge.From("projects", Project.Type).
			Ref("tasks"),
	}
}
