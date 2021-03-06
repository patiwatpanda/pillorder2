// Code generated by entc, DO NOT EDIT.

package pillorderitem

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/patiwatpanda/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
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
func IDGT(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PillorderItemName applies equality check predicate on the "PillorderItem_name" field. It's identical to PillorderItemNameEQ.
func PillorderItemName(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameEQ applies the EQ predicate on the "PillorderItem_name" field.
func PillorderItemNameEQ(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameNEQ applies the NEQ predicate on the "PillorderItem_name" field.
func PillorderItemNameNEQ(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameIn applies the In predicate on the "PillorderItem_name" field.
func PillorderItemNameIn(vs ...string) predicate.PillorderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PillorderItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPillorderItemName), v...))
	})
}

// PillorderItemNameNotIn applies the NotIn predicate on the "PillorderItem_name" field.
func PillorderItemNameNotIn(vs ...string) predicate.PillorderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PillorderItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPillorderItemName), v...))
	})
}

// PillorderItemNameGT applies the GT predicate on the "PillorderItem_name" field.
func PillorderItemNameGT(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameGTE applies the GTE predicate on the "PillorderItem_name" field.
func PillorderItemNameGTE(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameLT applies the LT predicate on the "PillorderItem_name" field.
func PillorderItemNameLT(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameLTE applies the LTE predicate on the "PillorderItem_name" field.
func PillorderItemNameLTE(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameContains applies the Contains predicate on the "PillorderItem_name" field.
func PillorderItemNameContains(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameHasPrefix applies the HasPrefix predicate on the "PillorderItem_name" field.
func PillorderItemNameHasPrefix(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameHasSuffix applies the HasSuffix predicate on the "PillorderItem_name" field.
func PillorderItemNameHasSuffix(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameEqualFold applies the EqualFold predicate on the "PillorderItem_name" field.
func PillorderItemNameEqualFold(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPillorderItemName), v))
	})
}

// PillorderItemNameContainsFold applies the ContainsFold predicate on the "PillorderItem_name" field.
func PillorderItemNameContainsFold(v string) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPillorderItemName), v))
	})
}

// HasPillorders applies the HasEdge predicate on the "pillorders" edge.
func HasPillorders() predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PillordersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PillordersTable, PillordersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPillordersWith applies the HasEdge predicate on the "pillorders" edge with a given conditions (other predicates).
func HasPillordersWith(preds ...predicate.Pillorder) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
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
func And(predicates ...predicate.PillorderItem) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PillorderItem) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
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
func Not(p predicate.PillorderItem) predicate.PillorderItem {
	return predicate.PillorderItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
