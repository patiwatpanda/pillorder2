package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
// Patient holds the schema definition for the Patient entity.

type Patient struct {
   ent.Schema
}
// Fields of the Patient.
func (Patient) Fields() []ent.Field {
   return []ent.Field{
       field.String("Patient_name").NotEmpty(),
       field.Int("Patient_age").Positive(),
   }
}
 

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("pillorders", Pillorder.Type).StorageKey(edge.Column("patient_id")),
    }
}