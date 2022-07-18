// Code generated by ent, DO NOT EDIT.

package taxonomyvocabulary

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/oaktreecode/oakapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Vid applies equality check predicate on the "vid" field. It's identical to VidEQ.
func Vid(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVid), v))
	})
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// VidEQ applies the EQ predicate on the "vid" field.
func VidEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVid), v))
	})
}

// VidNEQ applies the NEQ predicate on the "vid" field.
func VidNEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVid), v))
	})
}

// VidIn applies the In predicate on the "vid" field.
func VidIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVid), v...))
	})
}

// VidNotIn applies the NotIn predicate on the "vid" field.
func VidNotIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVid), v...))
	})
}

// VidGT applies the GT predicate on the "vid" field.
func VidGT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVid), v))
	})
}

// VidGTE applies the GTE predicate on the "vid" field.
func VidGTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVid), v))
	})
}

// VidLT applies the LT predicate on the "vid" field.
func VidLT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVid), v))
	})
}

// VidLTE applies the LTE predicate on the "vid" field.
func VidLTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVid), v))
	})
}

// VidContains applies the Contains predicate on the "vid" field.
func VidContains(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVid), v))
	})
}

// VidHasPrefix applies the HasPrefix predicate on the "vid" field.
func VidHasPrefix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVid), v))
	})
}

// VidHasSuffix applies the HasSuffix predicate on the "vid" field.
func VidHasSuffix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVid), v))
	})
}

// VidEqualFold applies the EqualFold predicate on the "vid" field.
func VidEqualFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVid), v))
	})
}

// VidContainsFold applies the ContainsFold predicate on the "vid" field.
func VidContainsFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVid), v))
	})
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLabel), v))
	})
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLabel), v))
	})
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLabel), v...))
	})
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLabel), v...))
	})
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLabel), v))
	})
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLabel), v))
	})
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLabel), v))
	})
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLabel), v))
	})
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLabel), v))
	})
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLabel), v))
	})
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLabel), v))
	})
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLabel), v))
	})
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLabel), v))
	})
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeight), v))
	})
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeight), v...))
	})
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeight), v...))
	})
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeight), v))
	})
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeight), v))
	})
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeight), v))
	})
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeight), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TaxonomyVocabulary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaxonomyVocabulary) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaxonomyVocabulary) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
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
func Not(p predicate.TaxonomyVocabulary) predicate.TaxonomyVocabulary {
	return predicate.TaxonomyVocabulary(func(s *sql.Selector) {
		p(s.Not())
	})
}