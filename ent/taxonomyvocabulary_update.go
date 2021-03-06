// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/oaktreecode/oakapi/ent/predicate"
	"github.com/oaktreecode/oakapi/ent/taxonomyvocabulary"
)

// TaxonomyVocabularyUpdate is the builder for updating TaxonomyVocabulary entities.
type TaxonomyVocabularyUpdate struct {
	config
	hooks    []Hook
	mutation *TaxonomyVocabularyMutation
}

// Where appends a list predicates to the TaxonomyVocabularyUpdate builder.
func (tvu *TaxonomyVocabularyUpdate) Where(ps ...predicate.TaxonomyVocabulary) *TaxonomyVocabularyUpdate {
	tvu.mutation.Where(ps...)
	return tvu
}

// SetVid sets the "vid" field.
func (tvu *TaxonomyVocabularyUpdate) SetVid(s string) *TaxonomyVocabularyUpdate {
	tvu.mutation.SetVid(s)
	return tvu
}

// SetLabel sets the "label" field.
func (tvu *TaxonomyVocabularyUpdate) SetLabel(s string) *TaxonomyVocabularyUpdate {
	tvu.mutation.SetLabel(s)
	return tvu
}

// SetWeight sets the "weight" field.
func (tvu *TaxonomyVocabularyUpdate) SetWeight(i int) *TaxonomyVocabularyUpdate {
	tvu.mutation.ResetWeight()
	tvu.mutation.SetWeight(i)
	return tvu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (tvu *TaxonomyVocabularyUpdate) SetNillableWeight(i *int) *TaxonomyVocabularyUpdate {
	if i != nil {
		tvu.SetWeight(*i)
	}
	return tvu
}

// AddWeight adds i to the "weight" field.
func (tvu *TaxonomyVocabularyUpdate) AddWeight(i int) *TaxonomyVocabularyUpdate {
	tvu.mutation.AddWeight(i)
	return tvu
}

// SetDescription sets the "description" field.
func (tvu *TaxonomyVocabularyUpdate) SetDescription(s string) *TaxonomyVocabularyUpdate {
	tvu.mutation.SetDescription(s)
	return tvu
}

// Mutation returns the TaxonomyVocabularyMutation object of the builder.
func (tvu *TaxonomyVocabularyUpdate) Mutation() *TaxonomyVocabularyMutation {
	return tvu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tvu *TaxonomyVocabularyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tvu.hooks) == 0 {
		if err = tvu.check(); err != nil {
			return 0, err
		}
		affected, err = tvu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaxonomyVocabularyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tvu.check(); err != nil {
				return 0, err
			}
			tvu.mutation = mutation
			affected, err = tvu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tvu.hooks) - 1; i >= 0; i-- {
			if tvu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tvu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tvu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tvu *TaxonomyVocabularyUpdate) SaveX(ctx context.Context) int {
	affected, err := tvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tvu *TaxonomyVocabularyUpdate) Exec(ctx context.Context) error {
	_, err := tvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvu *TaxonomyVocabularyUpdate) ExecX(ctx context.Context) {
	if err := tvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tvu *TaxonomyVocabularyUpdate) check() error {
	if v, ok := tvu.mutation.Vid(); ok {
		if err := taxonomyvocabulary.VidValidator(v); err != nil {
			return &ValidationError{Name: "vid", err: fmt.Errorf(`ent: validator failed for field "TaxonomyVocabulary.vid": %w`, err)}
		}
	}
	if v, ok := tvu.mutation.Label(); ok {
		if err := taxonomyvocabulary.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "TaxonomyVocabulary.label": %w`, err)}
		}
	}
	return nil
}

func (tvu *TaxonomyVocabularyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taxonomyvocabulary.Table,
			Columns: taxonomyvocabulary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: taxonomyvocabulary.FieldID,
			},
		},
	}
	if ps := tvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvu.mutation.Vid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldVid,
		})
	}
	if value, ok := tvu.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldLabel,
		})
	}
	if value, ok := tvu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: taxonomyvocabulary.FieldWeight,
		})
	}
	if value, ok := tvu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: taxonomyvocabulary.FieldWeight,
		})
	}
	if value, ok := tvu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldDescription,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taxonomyvocabulary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TaxonomyVocabularyUpdateOne is the builder for updating a single TaxonomyVocabulary entity.
type TaxonomyVocabularyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaxonomyVocabularyMutation
}

// SetVid sets the "vid" field.
func (tvuo *TaxonomyVocabularyUpdateOne) SetVid(s string) *TaxonomyVocabularyUpdateOne {
	tvuo.mutation.SetVid(s)
	return tvuo
}

// SetLabel sets the "label" field.
func (tvuo *TaxonomyVocabularyUpdateOne) SetLabel(s string) *TaxonomyVocabularyUpdateOne {
	tvuo.mutation.SetLabel(s)
	return tvuo
}

// SetWeight sets the "weight" field.
func (tvuo *TaxonomyVocabularyUpdateOne) SetWeight(i int) *TaxonomyVocabularyUpdateOne {
	tvuo.mutation.ResetWeight()
	tvuo.mutation.SetWeight(i)
	return tvuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (tvuo *TaxonomyVocabularyUpdateOne) SetNillableWeight(i *int) *TaxonomyVocabularyUpdateOne {
	if i != nil {
		tvuo.SetWeight(*i)
	}
	return tvuo
}

// AddWeight adds i to the "weight" field.
func (tvuo *TaxonomyVocabularyUpdateOne) AddWeight(i int) *TaxonomyVocabularyUpdateOne {
	tvuo.mutation.AddWeight(i)
	return tvuo
}

// SetDescription sets the "description" field.
func (tvuo *TaxonomyVocabularyUpdateOne) SetDescription(s string) *TaxonomyVocabularyUpdateOne {
	tvuo.mutation.SetDescription(s)
	return tvuo
}

// Mutation returns the TaxonomyVocabularyMutation object of the builder.
func (tvuo *TaxonomyVocabularyUpdateOne) Mutation() *TaxonomyVocabularyMutation {
	return tvuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tvuo *TaxonomyVocabularyUpdateOne) Select(field string, fields ...string) *TaxonomyVocabularyUpdateOne {
	tvuo.fields = append([]string{field}, fields...)
	return tvuo
}

// Save executes the query and returns the updated TaxonomyVocabulary entity.
func (tvuo *TaxonomyVocabularyUpdateOne) Save(ctx context.Context) (*TaxonomyVocabulary, error) {
	var (
		err  error
		node *TaxonomyVocabulary
	)
	if len(tvuo.hooks) == 0 {
		if err = tvuo.check(); err != nil {
			return nil, err
		}
		node, err = tvuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaxonomyVocabularyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tvuo.check(); err != nil {
				return nil, err
			}
			tvuo.mutation = mutation
			node, err = tvuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tvuo.hooks) - 1; i >= 0; i-- {
			if tvuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tvuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tvuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TaxonomyVocabulary)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TaxonomyVocabularyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tvuo *TaxonomyVocabularyUpdateOne) SaveX(ctx context.Context) *TaxonomyVocabulary {
	node, err := tvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tvuo *TaxonomyVocabularyUpdateOne) Exec(ctx context.Context) error {
	_, err := tvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvuo *TaxonomyVocabularyUpdateOne) ExecX(ctx context.Context) {
	if err := tvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tvuo *TaxonomyVocabularyUpdateOne) check() error {
	if v, ok := tvuo.mutation.Vid(); ok {
		if err := taxonomyvocabulary.VidValidator(v); err != nil {
			return &ValidationError{Name: "vid", err: fmt.Errorf(`ent: validator failed for field "TaxonomyVocabulary.vid": %w`, err)}
		}
	}
	if v, ok := tvuo.mutation.Label(); ok {
		if err := taxonomyvocabulary.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "TaxonomyVocabulary.label": %w`, err)}
		}
	}
	return nil
}

func (tvuo *TaxonomyVocabularyUpdateOne) sqlSave(ctx context.Context) (_node *TaxonomyVocabulary, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taxonomyvocabulary.Table,
			Columns: taxonomyvocabulary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: taxonomyvocabulary.FieldID,
			},
		},
	}
	id, ok := tvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaxonomyVocabulary.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taxonomyvocabulary.FieldID)
		for _, f := range fields {
			if !taxonomyvocabulary.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taxonomyvocabulary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvuo.mutation.Vid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldVid,
		})
	}
	if value, ok := tvuo.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldLabel,
		})
	}
	if value, ok := tvuo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: taxonomyvocabulary.FieldWeight,
		})
	}
	if value, ok := tvuo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: taxonomyvocabulary.FieldWeight,
		})
	}
	if value, ok := tvuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taxonomyvocabulary.FieldDescription,
		})
	}
	_node = &TaxonomyVocabulary{config: tvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taxonomyvocabulary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
