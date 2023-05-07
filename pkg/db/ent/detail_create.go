// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/detail"
	"github.com/google/uuid"
)

// DetailCreate is the builder for creating a Detail entity.
type DetailCreate struct {
	config
	mutation *DetailMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dc *DetailCreate) SetCreatedAt(u uint32) *DetailCreate {
	dc.mutation.SetCreatedAt(u)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DetailCreate) SetNillableCreatedAt(u *uint32) *DetailCreate {
	if u != nil {
		dc.SetCreatedAt(*u)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DetailCreate) SetUpdatedAt(u uint32) *DetailCreate {
	dc.mutation.SetUpdatedAt(u)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DetailCreate) SetNillableUpdatedAt(u *uint32) *DetailCreate {
	if u != nil {
		dc.SetUpdatedAt(*u)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DetailCreate) SetDeletedAt(u uint32) *DetailCreate {
	dc.mutation.SetDeletedAt(u)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DetailCreate) SetNillableDeletedAt(u *uint32) *DetailCreate {
	if u != nil {
		dc.SetDeletedAt(*u)
	}
	return dc
}

// SetEntID sets the "ent_id" field.
func (dc *DetailCreate) SetEntID(u uuid.UUID) *DetailCreate {
	dc.mutation.SetEntID(u)
	return dc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (dc *DetailCreate) SetNillableEntID(u *uuid.UUID) *DetailCreate {
	if u != nil {
		dc.SetEntID(*u)
	}
	return dc
}

// SetSampleCol sets the "sample_col" field.
func (dc *DetailCreate) SetSampleCol(s string) *DetailCreate {
	dc.mutation.SetSampleCol(s)
	return dc
}

// SetNillableSampleCol sets the "sample_col" field if the given value is not nil.
func (dc *DetailCreate) SetNillableSampleCol(s *string) *DetailCreate {
	if s != nil {
		dc.SetSampleCol(*s)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DetailCreate) SetID(u uint32) *DetailCreate {
	dc.mutation.SetID(u)
	return dc
}

// Mutation returns the DetailMutation object of the builder.
func (dc *DetailCreate) Mutation() *DetailMutation {
	return dc.mutation
}

// Save creates the Detail in the database.
func (dc *DetailCreate) Save(ctx context.Context) (*Detail, error) {
	var (
		err  error
		node *Detail
	)
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Detail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DetailCreate) SaveX(ctx context.Context) *Detail {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DetailCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DetailCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DetailCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if detail.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := detail.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		if detail.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		if detail.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := detail.DefaultDeletedAt()
		dc.mutation.SetDeletedAt(v)
	}
	if _, ok := dc.mutation.EntID(); !ok {
		if detail.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized detail.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := detail.DefaultEntID()
		dc.mutation.SetEntID(v)
	}
	if _, ok := dc.mutation.SampleCol(); !ok {
		v := detail.DefaultSampleCol
		dc.mutation.SetSampleCol(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DetailCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Detail.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Detail.updated_at"`)}
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Detail.deleted_at"`)}
	}
	if _, ok := dc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Detail.ent_id"`)}
	}
	return nil
}

func (dc *DetailCreate) sqlSave(ctx context.Context) (*Detail, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (dc *DetailCreate) createSpec() (*Detail, *sqlgraph.CreateSpec) {
	var (
		_node = &Detail{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: detail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: detail.FieldID,
			},
		}
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := dc.mutation.SampleCol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldSampleCol,
		})
		_node.SampleCol = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Detail.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DetailUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (dc *DetailCreate) OnConflict(opts ...sql.ConflictOption) *DetailUpsertOne {
	dc.conflict = opts
	return &DetailUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Detail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dc *DetailCreate) OnConflictColumns(columns ...string) *DetailUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DetailUpsertOne{
		create: dc,
	}
}

type (
	// DetailUpsertOne is the builder for "upsert"-ing
	//  one Detail node.
	DetailUpsertOne struct {
		create *DetailCreate
	}

	// DetailUpsert is the "OnConflict" setter.
	DetailUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DetailUpsert) SetCreatedAt(v uint32) *DetailUpsert {
	u.Set(detail.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DetailUpsert) UpdateCreatedAt() *DetailUpsert {
	u.SetExcluded(detail.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DetailUpsert) AddCreatedAt(v uint32) *DetailUpsert {
	u.Add(detail.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DetailUpsert) SetUpdatedAt(v uint32) *DetailUpsert {
	u.Set(detail.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DetailUpsert) UpdateUpdatedAt() *DetailUpsert {
	u.SetExcluded(detail.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DetailUpsert) AddUpdatedAt(v uint32) *DetailUpsert {
	u.Add(detail.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DetailUpsert) SetDeletedAt(v uint32) *DetailUpsert {
	u.Set(detail.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DetailUpsert) UpdateDeletedAt() *DetailUpsert {
	u.SetExcluded(detail.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DetailUpsert) AddDeletedAt(v uint32) *DetailUpsert {
	u.Add(detail.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *DetailUpsert) SetEntID(v uuid.UUID) *DetailUpsert {
	u.Set(detail.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DetailUpsert) UpdateEntID() *DetailUpsert {
	u.SetExcluded(detail.FieldEntID)
	return u
}

// SetSampleCol sets the "sample_col" field.
func (u *DetailUpsert) SetSampleCol(v string) *DetailUpsert {
	u.Set(detail.FieldSampleCol, v)
	return u
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *DetailUpsert) UpdateSampleCol() *DetailUpsert {
	u.SetExcluded(detail.FieldSampleCol)
	return u
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *DetailUpsert) ClearSampleCol() *DetailUpsert {
	u.SetNull(detail.FieldSampleCol)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Detail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(detail.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DetailUpsertOne) UpdateNewValues() *DetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(detail.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Detail.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DetailUpsertOne) Ignore() *DetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DetailUpsertOne) DoNothing() *DetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DetailCreate.OnConflict
// documentation for more info.
func (u *DetailUpsertOne) Update(set func(*DetailUpsert)) *DetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DetailUpsertOne) SetCreatedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DetailUpsertOne) AddCreatedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DetailUpsertOne) UpdateCreatedAt() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DetailUpsertOne) SetUpdatedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DetailUpsertOne) AddUpdatedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DetailUpsertOne) UpdateUpdatedAt() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DetailUpsertOne) SetDeletedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DetailUpsertOne) AddDeletedAt(v uint32) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DetailUpsertOne) UpdateDeletedAt() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DetailUpsertOne) SetEntID(v uuid.UUID) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DetailUpsertOne) UpdateEntID() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateEntID()
	})
}

// SetSampleCol sets the "sample_col" field.
func (u *DetailUpsertOne) SetSampleCol(v string) *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.SetSampleCol(v)
	})
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *DetailUpsertOne) UpdateSampleCol() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateSampleCol()
	})
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *DetailUpsertOne) ClearSampleCol() *DetailUpsertOne {
	return u.Update(func(s *DetailUpsert) {
		s.ClearSampleCol()
	})
}

// Exec executes the query.
func (u *DetailUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DetailCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DetailUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DetailUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DetailUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DetailCreateBulk is the builder for creating many Detail entities in bulk.
type DetailCreateBulk struct {
	config
	builders []*DetailCreate
	conflict []sql.ConflictOption
}

// Save creates the Detail entities in the database.
func (dcb *DetailCreateBulk) Save(ctx context.Context) ([]*Detail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Detail, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DetailMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DetailCreateBulk) SaveX(ctx context.Context) []*Detail {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DetailCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DetailCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Detail.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DetailUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (dcb *DetailCreateBulk) OnConflict(opts ...sql.ConflictOption) *DetailUpsertBulk {
	dcb.conflict = opts
	return &DetailUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Detail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dcb *DetailCreateBulk) OnConflictColumns(columns ...string) *DetailUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DetailUpsertBulk{
		create: dcb,
	}
}

// DetailUpsertBulk is the builder for "upsert"-ing
// a bulk of Detail nodes.
type DetailUpsertBulk struct {
	create *DetailCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Detail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(detail.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DetailUpsertBulk) UpdateNewValues() *DetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(detail.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Detail.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DetailUpsertBulk) Ignore() *DetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DetailUpsertBulk) DoNothing() *DetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DetailCreateBulk.OnConflict
// documentation for more info.
func (u *DetailUpsertBulk) Update(set func(*DetailUpsert)) *DetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DetailUpsertBulk) SetCreatedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DetailUpsertBulk) AddCreatedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DetailUpsertBulk) UpdateCreatedAt() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DetailUpsertBulk) SetUpdatedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DetailUpsertBulk) AddUpdatedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DetailUpsertBulk) UpdateUpdatedAt() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DetailUpsertBulk) SetDeletedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DetailUpsertBulk) AddDeletedAt(v uint32) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DetailUpsertBulk) UpdateDeletedAt() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DetailUpsertBulk) SetEntID(v uuid.UUID) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DetailUpsertBulk) UpdateEntID() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateEntID()
	})
}

// SetSampleCol sets the "sample_col" field.
func (u *DetailUpsertBulk) SetSampleCol(v string) *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.SetSampleCol(v)
	})
}

// UpdateSampleCol sets the "sample_col" field to the value that was provided on create.
func (u *DetailUpsertBulk) UpdateSampleCol() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.UpdateSampleCol()
	})
}

// ClearSampleCol clears the value of the "sample_col" field.
func (u *DetailUpsertBulk) ClearSampleCol() *DetailUpsertBulk {
	return u.Update(func(s *DetailUpsert) {
		s.ClearSampleCol()
	})
}

// Exec executes the query.
func (u *DetailUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DetailCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DetailCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DetailUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
