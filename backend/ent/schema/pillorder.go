package schema
 
import (
    
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Pillorder holds the schema definition for the Pillorder entity.
type Pillorder struct {
   ent.Schema
}
 
// Fields of the Pillorder.
func (Pillorder) Fields() []ent.Field {
   return []ent.Field{
    field.String("PillorderNameID").NotEmpty(),
       field.Time("PillorderDate"),
   }
}
 
// Edges of the Pillorder.
func (Pillorder) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("employee", Employee.Type).
            Ref("pillorders").
            Unique(),
        edge.From("patient", Patient.Type).
            Ref("pillorders").
            Unique(),
        edge.From("dentist", Dentist.Type).
            Ref("pillorders").
            Unique(),
        edge.From("pillorderitem", PillorderItem.Type).
            Ref("pillorders").
            Unique(),    
            
    }
}
