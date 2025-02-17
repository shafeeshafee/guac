// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// BillOfMaterialsUpdate is the builder for updating BillOfMaterials entities.
type BillOfMaterialsUpdate struct {
	config
	hooks    []Hook
	mutation *BillOfMaterialsMutation
}

// Where appends a list predicates to the BillOfMaterialsUpdate builder.
func (bomu *BillOfMaterialsUpdate) Where(ps ...predicate.BillOfMaterials) *BillOfMaterialsUpdate {
	bomu.mutation.Where(ps...)
	return bomu
}

// SetPackageID sets the "package_id" field.
func (bomu *BillOfMaterialsUpdate) SetPackageID(i int) *BillOfMaterialsUpdate {
	bomu.mutation.SetPackageID(i)
	return bomu
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (bomu *BillOfMaterialsUpdate) SetNillablePackageID(i *int) *BillOfMaterialsUpdate {
	if i != nil {
		bomu.SetPackageID(*i)
	}
	return bomu
}

// ClearPackageID clears the value of the "package_id" field.
func (bomu *BillOfMaterialsUpdate) ClearPackageID() *BillOfMaterialsUpdate {
	bomu.mutation.ClearPackageID()
	return bomu
}

// SetArtifactID sets the "artifact_id" field.
func (bomu *BillOfMaterialsUpdate) SetArtifactID(i int) *BillOfMaterialsUpdate {
	bomu.mutation.SetArtifactID(i)
	return bomu
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (bomu *BillOfMaterialsUpdate) SetNillableArtifactID(i *int) *BillOfMaterialsUpdate {
	if i != nil {
		bomu.SetArtifactID(*i)
	}
	return bomu
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (bomu *BillOfMaterialsUpdate) ClearArtifactID() *BillOfMaterialsUpdate {
	bomu.mutation.ClearArtifactID()
	return bomu
}

// SetURI sets the "uri" field.
func (bomu *BillOfMaterialsUpdate) SetURI(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetURI(s)
	return bomu
}

// SetAlgorithm sets the "algorithm" field.
func (bomu *BillOfMaterialsUpdate) SetAlgorithm(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetAlgorithm(s)
	return bomu
}

// SetDigest sets the "digest" field.
func (bomu *BillOfMaterialsUpdate) SetDigest(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetDigest(s)
	return bomu
}

// SetDownloadLocation sets the "download_location" field.
func (bomu *BillOfMaterialsUpdate) SetDownloadLocation(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetDownloadLocation(s)
	return bomu
}

// SetOrigin sets the "origin" field.
func (bomu *BillOfMaterialsUpdate) SetOrigin(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetOrigin(s)
	return bomu
}

// SetCollector sets the "collector" field.
func (bomu *BillOfMaterialsUpdate) SetCollector(s string) *BillOfMaterialsUpdate {
	bomu.mutation.SetCollector(s)
	return bomu
}

// SetKnownSince sets the "known_since" field.
func (bomu *BillOfMaterialsUpdate) SetKnownSince(t time.Time) *BillOfMaterialsUpdate {
	bomu.mutation.SetKnownSince(t)
	return bomu
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (bomu *BillOfMaterialsUpdate) SetPackage(p *PackageVersion) *BillOfMaterialsUpdate {
	return bomu.SetPackageID(p.ID)
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (bomu *BillOfMaterialsUpdate) SetArtifact(a *Artifact) *BillOfMaterialsUpdate {
	return bomu.SetArtifactID(a.ID)
}

// Mutation returns the BillOfMaterialsMutation object of the builder.
func (bomu *BillOfMaterialsUpdate) Mutation() *BillOfMaterialsMutation {
	return bomu.mutation
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (bomu *BillOfMaterialsUpdate) ClearPackage() *BillOfMaterialsUpdate {
	bomu.mutation.ClearPackage()
	return bomu
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (bomu *BillOfMaterialsUpdate) ClearArtifact() *BillOfMaterialsUpdate {
	bomu.mutation.ClearArtifact()
	return bomu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bomu *BillOfMaterialsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bomu.sqlSave, bomu.mutation, bomu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bomu *BillOfMaterialsUpdate) SaveX(ctx context.Context) int {
	affected, err := bomu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bomu *BillOfMaterialsUpdate) Exec(ctx context.Context) error {
	_, err := bomu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bomu *BillOfMaterialsUpdate) ExecX(ctx context.Context) {
	if err := bomu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bomu *BillOfMaterialsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(billofmaterials.Table, billofmaterials.Columns, sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeInt))
	if ps := bomu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bomu.mutation.URI(); ok {
		_spec.SetField(billofmaterials.FieldURI, field.TypeString, value)
	}
	if value, ok := bomu.mutation.Algorithm(); ok {
		_spec.SetField(billofmaterials.FieldAlgorithm, field.TypeString, value)
	}
	if value, ok := bomu.mutation.Digest(); ok {
		_spec.SetField(billofmaterials.FieldDigest, field.TypeString, value)
	}
	if value, ok := bomu.mutation.DownloadLocation(); ok {
		_spec.SetField(billofmaterials.FieldDownloadLocation, field.TypeString, value)
	}
	if value, ok := bomu.mutation.Origin(); ok {
		_spec.SetField(billofmaterials.FieldOrigin, field.TypeString, value)
	}
	if value, ok := bomu.mutation.Collector(); ok {
		_spec.SetField(billofmaterials.FieldCollector, field.TypeString, value)
	}
	if value, ok := bomu.mutation.KnownSince(); ok {
		_spec.SetField(billofmaterials.FieldKnownSince, field.TypeTime, value)
	}
	if bomu.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.PackageTable,
			Columns: []string{billofmaterials.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bomu.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.PackageTable,
			Columns: []string{billofmaterials.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bomu.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.ArtifactTable,
			Columns: []string{billofmaterials.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bomu.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.ArtifactTable,
			Columns: []string{billofmaterials.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bomu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billofmaterials.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bomu.mutation.done = true
	return n, nil
}

// BillOfMaterialsUpdateOne is the builder for updating a single BillOfMaterials entity.
type BillOfMaterialsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BillOfMaterialsMutation
}

// SetPackageID sets the "package_id" field.
func (bomuo *BillOfMaterialsUpdateOne) SetPackageID(i int) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetPackageID(i)
	return bomuo
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (bomuo *BillOfMaterialsUpdateOne) SetNillablePackageID(i *int) *BillOfMaterialsUpdateOne {
	if i != nil {
		bomuo.SetPackageID(*i)
	}
	return bomuo
}

// ClearPackageID clears the value of the "package_id" field.
func (bomuo *BillOfMaterialsUpdateOne) ClearPackageID() *BillOfMaterialsUpdateOne {
	bomuo.mutation.ClearPackageID()
	return bomuo
}

// SetArtifactID sets the "artifact_id" field.
func (bomuo *BillOfMaterialsUpdateOne) SetArtifactID(i int) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetArtifactID(i)
	return bomuo
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (bomuo *BillOfMaterialsUpdateOne) SetNillableArtifactID(i *int) *BillOfMaterialsUpdateOne {
	if i != nil {
		bomuo.SetArtifactID(*i)
	}
	return bomuo
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (bomuo *BillOfMaterialsUpdateOne) ClearArtifactID() *BillOfMaterialsUpdateOne {
	bomuo.mutation.ClearArtifactID()
	return bomuo
}

// SetURI sets the "uri" field.
func (bomuo *BillOfMaterialsUpdateOne) SetURI(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetURI(s)
	return bomuo
}

// SetAlgorithm sets the "algorithm" field.
func (bomuo *BillOfMaterialsUpdateOne) SetAlgorithm(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetAlgorithm(s)
	return bomuo
}

// SetDigest sets the "digest" field.
func (bomuo *BillOfMaterialsUpdateOne) SetDigest(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetDigest(s)
	return bomuo
}

// SetDownloadLocation sets the "download_location" field.
func (bomuo *BillOfMaterialsUpdateOne) SetDownloadLocation(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetDownloadLocation(s)
	return bomuo
}

// SetOrigin sets the "origin" field.
func (bomuo *BillOfMaterialsUpdateOne) SetOrigin(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetOrigin(s)
	return bomuo
}

// SetCollector sets the "collector" field.
func (bomuo *BillOfMaterialsUpdateOne) SetCollector(s string) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetCollector(s)
	return bomuo
}

// SetKnownSince sets the "known_since" field.
func (bomuo *BillOfMaterialsUpdateOne) SetKnownSince(t time.Time) *BillOfMaterialsUpdateOne {
	bomuo.mutation.SetKnownSince(t)
	return bomuo
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (bomuo *BillOfMaterialsUpdateOne) SetPackage(p *PackageVersion) *BillOfMaterialsUpdateOne {
	return bomuo.SetPackageID(p.ID)
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (bomuo *BillOfMaterialsUpdateOne) SetArtifact(a *Artifact) *BillOfMaterialsUpdateOne {
	return bomuo.SetArtifactID(a.ID)
}

// Mutation returns the BillOfMaterialsMutation object of the builder.
func (bomuo *BillOfMaterialsUpdateOne) Mutation() *BillOfMaterialsMutation {
	return bomuo.mutation
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (bomuo *BillOfMaterialsUpdateOne) ClearPackage() *BillOfMaterialsUpdateOne {
	bomuo.mutation.ClearPackage()
	return bomuo
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (bomuo *BillOfMaterialsUpdateOne) ClearArtifact() *BillOfMaterialsUpdateOne {
	bomuo.mutation.ClearArtifact()
	return bomuo
}

// Where appends a list predicates to the BillOfMaterialsUpdate builder.
func (bomuo *BillOfMaterialsUpdateOne) Where(ps ...predicate.BillOfMaterials) *BillOfMaterialsUpdateOne {
	bomuo.mutation.Where(ps...)
	return bomuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bomuo *BillOfMaterialsUpdateOne) Select(field string, fields ...string) *BillOfMaterialsUpdateOne {
	bomuo.fields = append([]string{field}, fields...)
	return bomuo
}

// Save executes the query and returns the updated BillOfMaterials entity.
func (bomuo *BillOfMaterialsUpdateOne) Save(ctx context.Context) (*BillOfMaterials, error) {
	return withHooks(ctx, bomuo.sqlSave, bomuo.mutation, bomuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bomuo *BillOfMaterialsUpdateOne) SaveX(ctx context.Context) *BillOfMaterials {
	node, err := bomuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bomuo *BillOfMaterialsUpdateOne) Exec(ctx context.Context) error {
	_, err := bomuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bomuo *BillOfMaterialsUpdateOne) ExecX(ctx context.Context) {
	if err := bomuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bomuo *BillOfMaterialsUpdateOne) sqlSave(ctx context.Context) (_node *BillOfMaterials, err error) {
	_spec := sqlgraph.NewUpdateSpec(billofmaterials.Table, billofmaterials.Columns, sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeInt))
	id, ok := bomuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BillOfMaterials.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bomuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billofmaterials.FieldID)
		for _, f := range fields {
			if !billofmaterials.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != billofmaterials.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bomuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bomuo.mutation.URI(); ok {
		_spec.SetField(billofmaterials.FieldURI, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.Algorithm(); ok {
		_spec.SetField(billofmaterials.FieldAlgorithm, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.Digest(); ok {
		_spec.SetField(billofmaterials.FieldDigest, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.DownloadLocation(); ok {
		_spec.SetField(billofmaterials.FieldDownloadLocation, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.Origin(); ok {
		_spec.SetField(billofmaterials.FieldOrigin, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.Collector(); ok {
		_spec.SetField(billofmaterials.FieldCollector, field.TypeString, value)
	}
	if value, ok := bomuo.mutation.KnownSince(); ok {
		_spec.SetField(billofmaterials.FieldKnownSince, field.TypeTime, value)
	}
	if bomuo.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.PackageTable,
			Columns: []string{billofmaterials.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bomuo.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.PackageTable,
			Columns: []string{billofmaterials.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bomuo.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.ArtifactTable,
			Columns: []string{billofmaterials.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bomuo.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   billofmaterials.ArtifactTable,
			Columns: []string{billofmaterials.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BillOfMaterials{config: bomuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bomuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billofmaterials.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bomuo.mutation.done = true
	return _node, nil
}
