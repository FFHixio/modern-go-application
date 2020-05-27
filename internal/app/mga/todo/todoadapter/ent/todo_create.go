// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/todo"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
}

// SetUID sets the uid field.
func (tc *TodoCreate) SetUID(s string) *TodoCreate {
	tc.mutation.SetUID(s)
	return tc
}

// SetTitle sets the title field.
func (tc *TodoCreate) SetTitle(s string) *TodoCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetCompleted sets the completed field.
func (tc *TodoCreate) SetCompleted(b bool) *TodoCreate {
	tc.mutation.SetCompleted(b)
	return tc
}

// SetCreatedAt sets the created_at field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the updated_at field.
func (tc *TodoCreate) SetUpdatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (tc *TodoCreate) SetNillableUpdatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	if _, ok := tc.mutation.UID(); !ok {
		return nil, errors.New("ent: missing required field \"uid\"")
	}
	if v, ok := tc.mutation.UID(); ok {
		if err := todo.UIDValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"uid\": %v", err)
		}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if _, ok := tc.mutation.Completed(); !ok {
		return nil, errors.New("ent: missing required field \"completed\"")
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := todo.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	var (
		err  error
		node *Todo
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	var (
		t     = &Todo{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldUID,
		})
		t.UID = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldTitle,
		})
		t.Title = value
	}
	if value, ok := tc.mutation.Completed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: todo.FieldCompleted,
		})
		t.Completed = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
		t.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldUpdatedAt,
		})
		t.UpdatedAt = value
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}
