// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mofu/ent/history"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HistoryCreate is the builder for creating a History entity.
type HistoryCreate struct {
	config
	mutation *HistoryMutation
	hooks    []Hook
}

// SetCreatorID sets the "creator_id" field.
func (hc *HistoryCreate) SetCreatorID(s string) *HistoryCreate {
	hc.mutation.SetCreatorID(s)
	return hc
}

// SetCreateAt sets the "create_at" field.
func (hc *HistoryCreate) SetCreateAt(t time.Time) *HistoryCreate {
	hc.mutation.SetCreateAt(t)
	return hc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableCreateAt(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetCreateAt(*t)
	}
	return hc
}

// SetLastUpdate sets the "last_update" field.
func (hc *HistoryCreate) SetLastUpdate(t time.Time) *HistoryCreate {
	hc.mutation.SetLastUpdate(t)
	return hc
}

// SetNillableLastUpdate sets the "last_update" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableLastUpdate(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetLastUpdate(*t)
	}
	return hc
}

// SetContentFlag sets the "content_flag" field.
func (hc *HistoryCreate) SetContentFlag(i int) *HistoryCreate {
	hc.mutation.SetContentFlag(i)
	return hc
}

// SetNillableContentFlag sets the "content_flag" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableContentFlag(i *int) *HistoryCreate {
	if i != nil {
		hc.SetContentFlag(*i)
	}
	return hc
}

// SetSentFlag sets the "sent_flag" field.
func (hc *HistoryCreate) SetSentFlag(i int) *HistoryCreate {
	hc.mutation.SetSentFlag(i)
	return hc
}

// SetNillableSentFlag sets the "sent_flag" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableSentFlag(i *int) *HistoryCreate {
	if i != nil {
		hc.SetSentFlag(*i)
	}
	return hc
}

// SetMentionedCount sets the "mentioned_count" field.
func (hc *HistoryCreate) SetMentionedCount(i int) *HistoryCreate {
	hc.mutation.SetMentionedCount(i)
	return hc
}

// SetNillableMentionedCount sets the "mentioned_count" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableMentionedCount(i *int) *HistoryCreate {
	if i != nil {
		hc.SetMentionedCount(*i)
	}
	return hc
}

// SetTakeEffectTime sets the "take_effect_time" field.
func (hc *HistoryCreate) SetTakeEffectTime(t time.Time) *HistoryCreate {
	hc.mutation.SetTakeEffectTime(t)
	return hc
}

// SetNillableTakeEffectTime sets the "take_effect_time" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableTakeEffectTime(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetTakeEffectTime(*t)
	}
	return hc
}

// SetSendingContent sets the "sending_content" field.
func (hc *HistoryCreate) SetSendingContent(b []byte) *HistoryCreate {
	hc.mutation.SetSendingContent(b)
	return hc
}

// SetID sets the "id" field.
func (hc *HistoryCreate) SetID(s string) *HistoryCreate {
	hc.mutation.SetID(s)
	return hc
}

// Mutation returns the HistoryMutation object of the builder.
func (hc *HistoryCreate) Mutation() *HistoryMutation {
	return hc.mutation
}

// Save creates the History in the database.
func (hc *HistoryCreate) Save(ctx context.Context) (*History, error) {
	var (
		err  error
		node *History
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*History)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HistoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistoryCreate) SaveX(ctx context.Context) *History {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HistoryCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HistoryCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HistoryCreate) defaults() {
	if _, ok := hc.mutation.CreateAt(); !ok {
		v := history.DefaultCreateAt()
		hc.mutation.SetCreateAt(v)
	}
	if _, ok := hc.mutation.LastUpdate(); !ok {
		v := history.DefaultLastUpdate()
		hc.mutation.SetLastUpdate(v)
	}
	if _, ok := hc.mutation.MentionedCount(); !ok {
		v := history.DefaultMentionedCount
		hc.mutation.SetMentionedCount(v)
	}
	if _, ok := hc.mutation.TakeEffectTime(); !ok {
		v := history.DefaultTakeEffectTime
		hc.mutation.SetTakeEffectTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HistoryCreate) check() error {
	if _, ok := hc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator_id", err: errors.New(`ent: missing required field "History.creator_id"`)}
	}
	if _, ok := hc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "History.create_at"`)}
	}
	if _, ok := hc.mutation.LastUpdate(); !ok {
		return &ValidationError{Name: "last_update", err: errors.New(`ent: missing required field "History.last_update"`)}
	}
	if _, ok := hc.mutation.MentionedCount(); !ok {
		return &ValidationError{Name: "mentioned_count", err: errors.New(`ent: missing required field "History.mentioned_count"`)}
	}
	if _, ok := hc.mutation.TakeEffectTime(); !ok {
		return &ValidationError{Name: "take_effect_time", err: errors.New(`ent: missing required field "History.take_effect_time"`)}
	}
	return nil
}

func (hc *HistoryCreate) sqlSave(ctx context.Context) (*History, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected History.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (hc *HistoryCreate) createSpec() (*History, *sqlgraph.CreateSpec) {
	var (
		_node = &History{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: history.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: history.FieldID,
			},
		}
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.CreatorID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: history.FieldCreatorID,
		})
		_node.CreatorID = value
	}
	if value, ok := hc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: history.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := hc.mutation.LastUpdate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: history.FieldLastUpdate,
		})
		_node.LastUpdate = value
	}
	if value, ok := hc.mutation.ContentFlag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: history.FieldContentFlag,
		})
		_node.ContentFlag = value
	}
	if value, ok := hc.mutation.SentFlag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: history.FieldSentFlag,
		})
		_node.SentFlag = value
	}
	if value, ok := hc.mutation.MentionedCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: history.FieldMentionedCount,
		})
		_node.MentionedCount = value
	}
	if value, ok := hc.mutation.TakeEffectTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: history.FieldTakeEffectTime,
		})
		_node.TakeEffectTime = value
	}
	if value, ok := hc.mutation.SendingContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: history.FieldSendingContent,
		})
		_node.SendingContent = value
	}
	return _node, _spec
}

// HistoryCreateBulk is the builder for creating many History entities in bulk.
type HistoryCreateBulk struct {
	config
	builders []*HistoryCreate
}

// Save creates the History entities in the database.
func (hcb *HistoryCreateBulk) Save(ctx context.Context) ([]*History, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*History, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HistoryCreateBulk) SaveX(ctx context.Context) []*History {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HistoryCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
