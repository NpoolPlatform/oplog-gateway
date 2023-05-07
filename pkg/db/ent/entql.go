// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/service-template/pkg/db/ent/detail"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/ignoreid"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/pubsubmessage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: detail.FieldID,
			},
		},
		Type: "Detail",
		Fields: map[string]*sqlgraph.FieldSpec{
			detail.FieldCreatedAt: {Type: field.TypeUint32, Column: detail.FieldCreatedAt},
			detail.FieldUpdatedAt: {Type: field.TypeUint32, Column: detail.FieldUpdatedAt},
			detail.FieldDeletedAt: {Type: field.TypeUint32, Column: detail.FieldDeletedAt},
			detail.FieldEntID:     {Type: field.TypeUUID, Column: detail.FieldEntID},
			detail.FieldSampleCol: {Type: field.TypeString, Column: detail.FieldSampleCol},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   ignoreid.Table,
			Columns: ignoreid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ignoreid.FieldID,
			},
		},
		Type: "IgnoreID",
		Fields: map[string]*sqlgraph.FieldSpec{
			ignoreid.FieldCreatedAt: {Type: field.TypeUint32, Column: ignoreid.FieldCreatedAt},
			ignoreid.FieldUpdatedAt: {Type: field.TypeUint32, Column: ignoreid.FieldUpdatedAt},
			ignoreid.FieldDeletedAt: {Type: field.TypeUint32, Column: ignoreid.FieldDeletedAt},
			ignoreid.FieldEntID:     {Type: field.TypeUUID, Column: ignoreid.FieldEntID},
			ignoreid.FieldSampleCol: {Type: field.TypeString, Column: ignoreid.FieldSampleCol},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   pubsubmessage.Table,
			Columns: pubsubmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pubsubmessage.FieldID,
			},
		},
		Type: "PubsubMessage",
		Fields: map[string]*sqlgraph.FieldSpec{
			pubsubmessage.FieldCreatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldCreatedAt},
			pubsubmessage.FieldUpdatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldUpdatedAt},
			pubsubmessage.FieldDeletedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldDeletedAt},
			pubsubmessage.FieldEntID:     {Type: field.TypeUUID, Column: pubsubmessage.FieldEntID},
			pubsubmessage.FieldMessageID: {Type: field.TypeString, Column: pubsubmessage.FieldMessageID},
			pubsubmessage.FieldState:     {Type: field.TypeString, Column: pubsubmessage.FieldState},
			pubsubmessage.FieldRespToID:  {Type: field.TypeUUID, Column: pubsubmessage.FieldRespToID},
			pubsubmessage.FieldUndoID:    {Type: field.TypeUUID, Column: pubsubmessage.FieldUndoID},
			pubsubmessage.FieldArguments: {Type: field.TypeString, Column: pubsubmessage.FieldArguments},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DetailQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DetailQuery builder.
func (dq *DetailQuery) Filter() *DetailFilter {
	return &DetailFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DetailMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DetailMutation builder.
func (m *DetailMutation) Filter() *DetailFilter {
	return &DetailFilter{config: m.config, predicateAdder: m}
}

// DetailFilter provides a generic filtering capability at runtime for DetailQuery.
type DetailFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DetailFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *DetailFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *DetailFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *DetailFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *DetailFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *DetailFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldEntID))
}

// WhereSampleCol applies the entql string predicate on the sample_col field.
func (f *DetailFilter) WhereSampleCol(p entql.StringP) {
	f.Where(p.Field(detail.FieldSampleCol))
}

// addPredicate implements the predicateAdder interface.
func (iiq *IgnoreIDQuery) addPredicate(pred func(s *sql.Selector)) {
	iiq.predicates = append(iiq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the IgnoreIDQuery builder.
func (iiq *IgnoreIDQuery) Filter() *IgnoreIDFilter {
	return &IgnoreIDFilter{config: iiq.config, predicateAdder: iiq}
}

// addPredicate implements the predicateAdder interface.
func (m *IgnoreIDMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the IgnoreIDMutation builder.
func (m *IgnoreIDMutation) Filter() *IgnoreIDFilter {
	return &IgnoreIDFilter{config: m.config, predicateAdder: m}
}

// IgnoreIDFilter provides a generic filtering capability at runtime for IgnoreIDQuery.
type IgnoreIDFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *IgnoreIDFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *IgnoreIDFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(ignoreid.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *IgnoreIDFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(ignoreid.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *IgnoreIDFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(ignoreid.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *IgnoreIDFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(ignoreid.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *IgnoreIDFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(ignoreid.FieldEntID))
}

// WhereSampleCol applies the entql string predicate on the sample_col field.
func (f *IgnoreIDFilter) WhereSampleCol(p entql.StringP) {
	f.Where(p.Field(ignoreid.FieldSampleCol))
}

// addPredicate implements the predicateAdder interface.
func (pmq *PubsubMessageQuery) addPredicate(pred func(s *sql.Selector)) {
	pmq.predicates = append(pmq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PubsubMessageQuery builder.
func (pmq *PubsubMessageQuery) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: pmq.config, predicateAdder: pmq}
}

// addPredicate implements the predicateAdder interface.
func (m *PubsubMessageMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PubsubMessageMutation builder.
func (m *PubsubMessageMutation) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: m.config, predicateAdder: m}
}

// PubsubMessageFilter provides a generic filtering capability at runtime for PubsubMessageQuery.
type PubsubMessageFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PubsubMessageFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PubsubMessageFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PubsubMessageFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PubsubMessageFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PubsubMessageFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *PubsubMessageFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldEntID))
}

// WhereMessageID applies the entql string predicate on the message_id field.
func (f *PubsubMessageFilter) WhereMessageID(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldMessageID))
}

// WhereState applies the entql string predicate on the state field.
func (f *PubsubMessageFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldState))
}

// WhereRespToID applies the entql [16]byte predicate on the resp_to_id field.
func (f *PubsubMessageFilter) WhereRespToID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldRespToID))
}

// WhereUndoID applies the entql [16]byte predicate on the undo_id field.
func (f *PubsubMessageFilter) WhereUndoID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldUndoID))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *PubsubMessageFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldArguments))
}
