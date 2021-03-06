// Code generated by entc, DO NOT EDIT.

package dentist

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/patiwatpanda/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DentistName applies equality check predicate on the "Dentist_name" field. It's identical to DentistNameEQ.
func DentistName(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDentistName), v))
	})
}

// DentistNameEQ applies the EQ predicate on the "Dentist_name" field.
func DentistNameEQ(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDentistName), v))
	})
}

// DentistNameNEQ applies the NEQ predicate on the "Dentist_name" field.
func DentistNameNEQ(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDentistName), v))
	})
}

// DentistNameIn applies the In predicate on the "Dentist_name" field.
func DentistNameIn(vs ...string) predicate.Dentist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDentistName), v...))
	})
}

// DentistNameNotIn applies the NotIn predicate on the "Dentist_name" field.
func DentistNameNotIn(vs ...string) predicate.Dentist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDentistName), v...))
	})
}

// DentistNameGT applies the GT predicate on the "Dentist_name" field.
func DentistNameGT(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDentistName), v))
	})
}

// DentistNameGTE applies the GTE predicate on the "Dentist_name" field.
func DentistNameGTE(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDentistName), v))
	})
}

// DentistNameLT applies the LT predicate on the "Dentist_name" field.
func DentistNameLT(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDentistName), v))
	})
}

// DentistNameLTE applies the LTE predicate on the "Dentist_name" field.
func DentistNameLTE(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDentistName), v))
	})
}

// DentistNameContains applies the Contains predicate on the "Dentist_name" field.
func DentistNameContains(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDentistName), v))
	})
}

// DentistNameHasPrefix applies the HasPrefix predicate on the "Dentist_name" field.
func DentistNameHasPrefix(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDentistName), v))
	})
}

// DentistNameHasSuffix applies the HasSuffix predicate on the "Dentist_name" field.
func DentistNameHasSuffix(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDentistName), v))
	})
}

// DentistNameEqualFold applies the EqualFold predicate on the "Dentist_name" field.
func DentistNameEqualFold(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDentistName), v))
	})
}

// DentistNameContainsFold applies the ContainsFold predicate on the "Dentist_name" field.
func DentistNameContainsFold(v string) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDentistName), v))
	})
}

// HasPillorders applies the HasEdge predicate on the "pillorders" edge.
func HasPillorders() predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PillordersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PillordersTable, PillordersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPillordersWith applies the HasEdge predicate on the "pillorders" edge with a given conditions (other predicates).
func HasPillordersWith(preds ...predicate.Pillorder) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PillordersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PillordersTable, PillordersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Dentist) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Dentist) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Dentist) predicate.Dentist {
	return predicate.Dentist(func(s *sql.Selector) {
		p(s.Not())
	})
}
