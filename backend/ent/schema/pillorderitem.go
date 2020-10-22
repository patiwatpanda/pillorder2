package schema

import (
    
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )
// PillorderItem holds the schema definition for the PillorderItem entity.
type PillorderItem struct {
	ent.Schema
}

// Fields of the PillorderItem.
func (PillorderItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("PillorderItem_name").NotEmpty(),
		
	}
}

// Edges of the PillorderItem.
func (PillorderItem) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("pillorders", Pillorder.Type).StorageKey(edge.Column("pillorderitem_id")),
  
    }
}
